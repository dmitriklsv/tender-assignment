package dataReader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dixitaniket/tender-assignment/sim"
	"github.com/dixitaniket/tender-assignment/types"
)

var ComplementaryDirections = types.ComplementaryDirections

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
		newline := strings.Split(scanner.Text(), " ")
		cityName := newline[0]
		city := types.InitCity(cityName)
		// situation where there a city is isolated in itself and there are no further connections
		if len(newline) == 1 {
			cityNameMapper[cityName] = &city

			continue
		}
		for _, dircity := range newline[1:] {
			// data[0] has direction and data[1] is the city name that it points to
			data := strings.Split(dircity, "=")
			if cityNameMapper[data[1]] != nil {
				// means the city is already been loaded into memory before
				city.ConnectCity(data[0], cityNameMapper[data[1]])
				cityNameMapper[data[1]].ConnectCity(ComplementaryDirections[data[0]], &city)
				cityNameMapper[cityName] = &city
			} else {
				// new city present in the direction is also not loaded into memory
				newcity := types.InitCity(data[1])
				city.ConnectCity(data[0], &newcity)
				newcity.ConnectCity(ComplementaryDirections[data[0]], &city)
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

// writes the final map into a txt file specificed by the output file path
func FileOutcome(s *sim.Simulation, outputFilePath string) error {
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
