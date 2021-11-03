package dataReader

import (
	"strconv"
	"testing"

	"github.com/dixitaniket/tender-assignment/sim"
)

func TestReadCityData(t *testing.T) {
	path := "../testData/testCityData.txt"
	cities, err := ReadCityData(path)
	if err != nil {
		t.Error(err)
	}
	if len(*cities) != 4 {
		t.Error("all cities are not loaded")
	}

	for _, city := range *cities {
		for direction, toCity := range city.Links {
			if city.Name != toCity.Links[OppositeDirections[direction]].Name {
				t.Error("opposite direction link does not work")
			}
		}

	}

}

var sampleNames = []string{"waybig", "alienx", "heatblast"}

func TestAlienNames(t *testing.T) {
	// this should give out alien names from alien0 to alien9
	alienNames, err := ReadAlienNames("", 10)
	if err != nil {
		t.Error(err)
	}
	for index, alien := range *alienNames {
		if alien.Name != "alien"+strconv.Itoa(index) {
			t.Error("error in random alien name generation")
		}
	}

	alienNames, err = ReadAlienNames("../testData/sampleAlienNames.txt", 3)
	if err != nil {
		t.Error(err)
	}
	for index, alien := range *alienNames {
		if alien.Name != sampleNames[index] {
			t.Error("alien names from file do not match")
		}
	}

}

func TestSimulationInitAndStart(t *testing.T) {
	path := "../testData/testCityData.txt"
	cities, err := ReadCityData(path)
	if err != nil {
		t.Error(err)
	}
	alienNames, err := ReadAlienNames("", 2)
	if err != nil {
		t.Error(err)
	}
	newSimulation := sim.InitNewSimulation(*alienNames, *cities, 1000)
	err = newSimulation.StartSimulation(false)
	if err != nil {
		t.Error(err)
	}
	deadCounter := 0
	trappedCounter := 0
	for _, alien := range newSimulation.Aliens {
		if alien.IsDead() {
			deadCounter += 1
		}
		if alien.IsTrapped() {
			trappedCounter += 1
		}
	}
	if deadCounter != newSimulation.AliensDead {
		t.Error("dead aliens from array and simulation do not match")
	}
	if trappedCounter != newSimulation.AliensTrapped {
		t.Error("trapped aliens from simulation and alien array does not match")
	}
	// the number of cities that are destoryed is always <= twice the number of aliens killed
	desCityCounter := 0
	for _, city := range newSimulation.Cities {
		if city.IfDestroyed() {
			desCityCounter += 1
		}
	}
	if desCityCounter > 2*deadCounter {
		t.Error("less cities destoryed as compared to aliens")
	}
	newSimulation.FinalOutcome()

}
