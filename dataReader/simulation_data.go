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
	// load all the cities from the file
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	// a store to check if any city is present in the map before
	cityNameMapper := make(map[string]*types.City)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// reads a new line from the file ,each new line contains name of the city and other cities to which it is connected
		newline := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		// newline structure : citynama<a> direction=cityname<b> direction=cityname+<c>
		cityName := newline[0]
		var city types.City
		// if the city has been already loaded into memory then fetch it from there
		if cityNameMapper[cityName] != nil {
			city = *cityNameMapper[cityName]
		} else {
			// otherwise create a new instance of the city with the cityname
			city = types.InitCity(cityName)
			cityNameMapper[cityName] = &city
		}
		// situation where there a city is isolated in itself and there are no further connections
		for _, dircity := range newline[1:] {
			// data[0] has direction and data[1] is the city name that it points to
			data := strings.Split(dircity, "=")
			// if the  city it is connected to is already loaded into memory they use to the reference
			if cityNameMapper[data[1]] != nil {
				// means the city is already been loaded into memory before
				// connect the new city to other cities based in direction
				if err := city.ConnectCity(data[0], cityNameMapper[data[1]]); err != nil {
					return nil, err
				}
				// the already exisiting city is connected to the new city with the opposite direction (reverse direction connection)
				if err := cityNameMapper[data[1]].ConnectCity(OppositeDirections[data[0]], &city); err != nil {
					return nil, err
				}
				// store the new city in the cityNameMapper
				cityNameMapper[cityName] = &city
			} else {
				// new city present in the direction is also not loaded into memory
				// generate a new city that is present in the direction
				newcity := types.InitCity(data[1])
				// connect newcity to the other newcity that was parsed from the direction
				if err := city.ConnectCity(data[0], &newcity); err != nil {
					return nil, err
				}
				// make the reverse direction connection
				if err := newcity.ConnectCity(OppositeDirections[data[0]], &city); err != nil {
					return nil, err
				}
				// store the newcity(parsed from the directon) and new city created at first
				cityNameMapper[newcity.Name] = &newcity
				cityNameMapper[cityName] = &city
			}
		}

	}
	// return the list of all the cities by iterating though the cityNameMapper
	cityMap := sim.Cities{}
	for _, city := range cityNameMapper {
		cityMap = append(cityMap, city)
	}
	return &cityMap, nil

}

func ReadCityData(path string) (*sim.Cities, error) {
	// reading the cities from the path and the other cities that are connected with it
	// cityList is the list of all the loaded city
	cityList, err := readCityData(path)
	if err != nil {
		return nil, err
	}
	// print the citymap : cityname and the cities that are connected with
	fmt.Println("CityMap")
	fmt.Println("-----------------")
	for _, i := range *cityList {
		fmt.Println("city name " + i.Name)
		for direction, j := range i.Links {
			fmt.Println(j.Name + " present in " + direction)
		}
		fmt.Println("--------------------")

	}
	// return the
	return cityList, nil

}

func ReadAlienNames(path string, N uint64) (*sim.Aliens, error) {

	aliens := sim.Aliens{}

	// if path is none then generate random alien names
	if path == "" {
		for i := uint64(0); i < N; i++ {
			newAlien := types.InitAlien("alien" + strconv.FormatUint(i, 10))
			fmt.Printf("New Alient %s generated\n", newAlien.Name)
			aliens = append(aliens, &newAlien)
		}
	} else {
		// opens the file specified in the path and read names of aliens from there
		f, err := os.Open(path)
		defer f.Close()
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(f)

		for i := uint64(0); i < N; i++ {
			// if the number of names are less then reuqired aliens throw and error
			if !scanner.Scan() {
				return nil, errors.New("total alien names less then the names provided in file")
			}
			// read alienname and create and store alien in memory
			alienName := strings.TrimSpace(scanner.Text())
			newAlien := types.InitAlien(alienName)
			fmt.Printf("New Alien %s generated\n", newAlien.Name)
			// append alien pointer to list of aliens
			aliens = append(aliens, &newAlien)
		}

	}
	fmt.Println("---------------")
	// return the list of aliens
	return &aliens, nil

}

// writes the final map into a txt file specificed by the output file path
func FinalCityMapToFile(s *sim.Simulation, outputFilePath string) error {
	// takes in file path to write all the cities that are not destroyed
	// into file and the other cities they are connected with
	f, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, fromCity := range s.Cities {
		if fromCity.IfDestroyed() {
			// if city is destoryed then it is not written in the file
			continue
		}
		// construct a message containing the city names and the other cities they are connected with
		msg := fromCity.Name
		for direction, toCity := range fromCity.Links {
			msg += " " + direction + "=" + toCity.Name
		}
		// write to file
		_, err := f.WriteString(msg + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
