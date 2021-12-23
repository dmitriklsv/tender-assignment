package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/dixitaniket/tender-assignment/dataReader"
	"github.com/dixitaniket/tender-assignment/sim"
)

var (
	CityFilePath      string
	AlienNameFilePath string
	TotalAliens       uint64
	MaxIterations     uint64
	OutputFilePath    string
	Verbose           bool
)

func init() {
	// flag module to parse cmd arguments and flags
	flag.StringVar(&CityFilePath, "city", "files/city.txt", "Path for file that contains details about the cities")
	flag.StringVar(&AlienNameFilePath, "alienname", "", "Path for file that contains alien names")
	flag.Uint64Var(&TotalAliens, "totalalien", 10, "Total number of aliens for simulation")
	flag.Uint64Var(&MaxIterations, "maxiter", 10000, "Max number of iteration for the simulation")
	flag.StringVar(&OutputFilePath, "output", "files/output.txt", "output file containing the cities that are not destroyed")
	flag.BoolVar(&Verbose, "verbose", false, "if true program prints the move made by aliens")
	flag.Parse()
}

func checkParsed() error {
	// basic checks to ensure that there are no invalid paramaters
	if CityFilePath == "" {
		return errors.New("City file path cannot be blank")
	}
	if MaxIterations == 0 {
		return errors.New("Max Iterations should be >=1")
	}
	if OutputFilePath == "" {
		return errors.New("output file path cannot be blank")
	}
	return nil
}

func Run() {

	err := checkParsed()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// reading file from the cityfilepath
	cityFile, err := dataReader.ReadCityData(CityFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// reading aliennames if AlienNameFilePath is not blank otherwise generating aliens
	newAliens, err := dataReader.ReadAlienNames(AlienNameFilePath, TotalAliens)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// creating a new simulation
	newSimulation := sim.InitNewSimulation(*newAliens, *cityFile, int(MaxIterations))

	newSimulation.StartSimulation(Verbose)
	// printing final outcome such as number of cities destoryed, aliens killed/trapped or survived
	newSimulation.FinalOutcome()
	// print the remaining cities and connected city from them in a file at OutputFilePath
	err = dataReader.FinalCityMapToFile(&newSimulation, OutputFilePath)
	if err != nil {
		fmt.Println(err)
	}

}
