package dataReader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dixitaniket/tender-assignment/sim"
	"github.com/dixitaniket/tender-assignment/types"
)

var OppositeDirections = types.OppositeDirections

// helps in updating the city connections

func readCityData(path string) (*sim.Cities, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	// loads all the cities from the file

	// a store to check if any city is present in the map before

	cityNameMapper := make(map[string]*types.City)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// reads a new line from the file ,each new line contains name of the city and other cities to which it is connected
		newline := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		cityName := newline[0]
		var city types.City
		if cityNameMapper[cityName] != nil {
			city = *cityNameMapper[cityName]
		} else {
			city = types.InitCity(cityName)
		}
		// situation where there a city is isolated in itself and there are no further connections
		if len(newline) == 1 {
			continue
		}
		for _, dircity := range newline[1:] {
			// data[0] has direction and data[1] is the city name that it points to
			data := strings.Split(dircity, "=")
			if cityNameMapper[data[1]] != nil {
				// means the city is already been loaded into memory before
				if err := city.ConnectCity(data[0], cityNameMapper[data[1]]); err != nil {
					return nil, err
				}
				if err := cityNameMapper[data[1]].ConnectCity(OppositeDirections[data[0]], &city); err != nil {
					return nil, err
				}
				cityNameMapper[cityName] = &city
			} else {
				// new city present in the direction is also not loaded into memory
				newcity := types.InitCity(data[1])
				if err := city.ConnectCity(data[0], &newcity); err != nil {
					return nil, err
				}
				if err := newcity.ConnectCity(OppositeDirections[data[0]], &city); err != nil {
					return nil, err
				}

				cityNameMapper[newcity.Name] = &newcity
				cityNameMapper[cityName] = &city
			}
		}

	}
	// return the map by iterating though all the cities that have been loaded to memory
	cityMap := sim.Cities{}
	for _, city := range cityNameMapper {
		cityMap = append(cityMap, city)
	}
	return &cityMap, nil

}

func ReadCityData(path string) (*sim.Cities, error) {
	cityMap, err := readCityData(path)
	if err != nil {
		return nil, err
	}
	fmt.Println("CityMap")
	fmt.Println("-----------------")
	for _, i := range *cityMap {
		fmt.Println("city name " + i.Name)
		for direction, j := range i.Links {
			fmt.Println(j.Name + " present in " + direction)
		}
		fmt.Println("--------------------")

	}
	return cityMap, nil

}

func ReadAlienNames(path string, N uint64) (*sim.Aliens, error) {
	// if path is none then generate random alien names
	aliens := sim.Aliens{}
	if path == "" {
		for i := uint64(0); i < N; i++ {
			newAlien := types.InitAlien("alien" + strconv.FormatUint(i, 10))
			fmt.Printf("New Alient %s generated\n", newAlien.Name)
			aliens = append(aliens, &newAlien)
		}
	} else {
		f, err := os.Open(path)
		defer f.Close()
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(f)

		for i := uint64(0); i < N; i++ {
			if !scanner.Scan() {
				return nil, errors.New("total aliens less then the names provided in file")
			}
			alienName := scanner.Text()
			newAlien := types.InitAlien(alienName)
			fmt.Printf("New Alien %s generated\n", newAlien.Name)
			aliens = append(aliens, &newAlien)
		}

	}
	fmt.Println("---------------")
	return &aliens, nil

}

// writes the final map into a txt file specificed by the output file path
func FinalCityMapToFile(s *sim.Simulation, outputFilePath string) error {
	f, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, fromCity := range s.Cities {
		if fromCity.IfDestroyed() {
			// fmt.Println(fromCity.Name)
			continue
		}
		msg := fromCity.Name
		for direction, toCity := range fromCity.Links {
			msg += " " + direction + "=" + toCity.Name
		}
		// fmt.Println(msg)
		_, err := f.WriteString(msg + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
