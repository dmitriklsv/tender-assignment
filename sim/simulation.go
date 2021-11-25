package sim

import (
	"fmt"
	"math/rand"

	"github.com/dixitaniket/tender-assignment/types"
)

// list of all aliens for the simulation
type Aliens []*types.Alien

// list of all the cities
type Cities []*types.City

// embedding aliens and cities type in simulation type
// simulation struct is generated at the start of each simulation run

// Simulation structure
// Iterations: current iterations
// MaxIterations: max number of iterations that are allowed in simulation
// AliensDead: Number of aliens that are dead
// AliensTrapped: Number of aliens that are trapped
// Aliens: list of the aliens
// Cities: list of the cities that are present in the simulation

type Simulation struct {
	Iterations    int
	MaxIterations int
	AliensDead    int
	AliensTrapped int
	Aliens
	Cities
}

// generate the new simulation
func InitNewSimulation(aliens Aliens, cities Cities, maxIterations int) Simulation {
	return Simulation{
		Iterations:    0,
		MaxIterations: maxIterations,
		Aliens:        aliens,
		Cities:        cities,
		AliensDead:    0,
		AliensTrapped: 0,
	}
}

// for every iteration aliens make there move and then we check what cities are destoryed and which aliens are dead or trapped
func (s *Simulation) StartSimulation(verbose bool) error {
	fmt.Println("starting simulation")
	fmt.Println("---------------------")
	for ; s.Iterations < s.MaxIterations; s.Iterations++ {
		fmt.Printf("\nCurrent iteration : %d\n", s.Iterations)
		if s.AliensDead+s.AliensTrapped == len(s.Aliens) {
			break
		}
		// iterate through all the aliens and update there state
		for i := range s.Aliens {
			alien := s.Aliens[i]
			err := s.Move(alien, verbose)
			if err != nil {
				fmt.Println(err)
			}

		}
		// iterate through all the cities and check if a city has more then 2 aliens then destroy that city
		for cityIndex := range s.Cities {
			city := s.Cities[cityIndex]
			if city.IfDestroyed() {
				continue
			}
			// fmt.Println(len(city.AlienPresent))

			if len(city.AlienPresent) >= 2 {
				s.AliensDead += len(city.AlienPresent)
				city.DestoryCity()
			}
		}
	}
	fmt.Printf("Simulation ended at iteration :%d\n\n", s.Iterations)
	return nil
}

func (s *Simulation) Move(a *types.Alien, verbose bool) error {
	// this would happen in the first iteration when we are
	if a.IsDead() || a.IsTrapped() {
		return nil
	}
	if a.CurrentCity == nil {
		a.InvadeCity(s.PickStartCity())
		fmt.Printf("Alien %s spanned in city %s\n", a.Name, a.CurrentCity.Name)
		return nil
		// if there is no area for the new alien to start
	}

	newcity := s.PickNextMove(a)
	if newcity == nil {
		return nil
	}
	if verbose {
		fmt.Printf("alien %s moves from city %s to city %s\n", a.Name, a.CurrentCity.Name, newcity.Name)
	}
	err := a.InvadeCity(newcity)
	if err != nil {
		return err
	}

	return nil
}

func (s *Simulation) PickNextMove(a *types.Alien) *types.City {

	var directions []string

	for i := range a.CurrentCity.Links {
		if c := a.CurrentCity.Links[i]; !c.IfDestroyed() {
			directions = append(directions, i)
		}

	}
	if len(directions) == 0 {
		a.Trap()
		s.AliensTrapped += 1
		return nil
	}
	index := rand.Intn(len(directions))
	return a.CurrentCity.Links[directions[index]]

}
func (s *Simulation) PickStartCity() *types.City {
	var Indexes []int
	for i := range s.Cities {
		if c := s.Cities[i]; !c.IfDestroyed() {
			Indexes = append(Indexes, i)
		}
	}
	randomIndex := rand.Intn(len(Indexes))

	return s.Cities[Indexes[randomIndex]]
}

func (s *Simulation) FinalOutcome() {
	fmt.Printf("Total Aliens Dead %d | Total Aliens Trapped %d | Total Aliens Survided %d\n", s.AliensDead, s.AliensTrapped, len(s.Aliens)-s.AliensDead)
	fmt.Println("\n-----------------------------\n")
	for _, city := range s.Cities {
		if city.IfDestroyed() {
			msg := "City " + city.Name + " was destroyed by: "
			for alienName := range city.AlienPresent {
				msg += alienName + ", "
			}
			fmt.Println(msg[:len(msg)-2])
		} else {
			fmt.Printf("City %s is safe\n", city.Name)
		}
	}

	fmt.Println("\n-----------------------------\n")
	for _, Alien := range s.Aliens {
		if Alien.IsDead() {
			fmt.Printf("Alien %s is dead\n", Alien.Name)
		} else if Alien.IsTrapped() {
			fmt.Printf("Alien %s is trapped in %s\n", Alien.Name, Alien.CurrentCity.Name)
		} else {
			fmt.Printf("Alien %s Survived and is currently in %s\n", Alien.Name, Alien.CurrentCity.Name)
		}

	}
}
