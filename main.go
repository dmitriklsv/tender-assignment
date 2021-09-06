package main

import (
	"fmt"
	"os"

	"github.com/dixitaniket/tender-assignment/dataReader"
	"github.com/dixitaniket/tender-assignment/sim"
	"github.com/dixitaniket/tender-assignment/types"
)

func main() {
	fmt.Println("Welcome to alien attack simulator")
	data, err := dataReader.ReadCityData("files.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newaliens := sim.Aliens{}
	alien1 := types.InitAlien("argernaut")
	alien2 := types.InitAlien("superna")
	alien3 := types.InitAlien("somename")
	alien4 := types.InitAlien("alienx ")
	alien5 := types.InitAlien("argernautt")
	alien6 := types.InitAlien("supernat")
	alien7 := types.InitAlien("somenamet")
	alien8 := types.InitAlien("alienxt")

	newaliens = append(newaliens, &alien1)
	newaliens = append(newaliens, &alien2)
	newaliens = append(newaliens, &alien3)
	newaliens = append(newaliens, &alien4)
	newaliens = append(newaliens, &alien5)
	newaliens = append(newaliens, &alien6)
	newaliens = append(newaliens, &alien7)
	newaliens = append(newaliens, &alien8)

	newSimulation := sim.InitNewSimulation(newaliens, *data, 10)
	newSimulation.StartSimulation()
	newSimulation.FinalOutcome()
	err = dataReader.FileOutcome(&newSimulation, "finalcityMap.txt")
	if err != nil {
		fmt.Println(err)
	}
}
