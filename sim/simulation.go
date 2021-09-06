package sim

import (
	"fmt"
	"math/rand"

	"github.com/dixitaniket/tender-assignment/types"
)

// collection of all aliens
type Aliens []*types.Alien

// collection of all the cities
type Cities []*types.City

// embedding aliens and cities type in simulation type
// simulation struct is generated at the start of each simulation run
type Simulation struct {
	Iterations    int
	MaxIterations int
	Aliens
	Cities
}

func InitNewSimulation(aliens Aliens, cities Cities, maxIterations int) Simulation {
	return Simulation{
		Iterations:    0,
		MaxIterations: maxIterations,
		Aliens:        aliens,
		Cities:        cities,
	}
}

// for every iteration aliens make there move and then we check what cities are destoryed and which aliens are dead or trapped
func (s *Simulation) StartSimulation() error {
	fmt.Println("starting simulation")
	fmt.Println("---------------------")
	for ; s.Iterations < s.MaxIterations; s.Iterations++ {
		fmt.Printf("Current iteration : %d\n", s.Iterations)

		for i := range s.Aliens {
			alien := s.Aliens[i]
			err := s.Move(alien)
			if err != nil {
				fmt.Println(err)
			}

		}
		for cityIndex := range s.Cities {
			city := s.Cities[cityIndex]
			if city.IfDestroyed() {
				continue
			}
			// fmt.Println(len(city.AlienPresent))

			if len(city.AlienPresent) >= 2 {
				city.DestoryCity()
			}
		}
	}
	return nil
}

func (s *Simulation) Move(a *types.Alien) error {
	// this would happen in the first iteration when we are
	if a.CurrentCity == nil {
		a.CurrentCity = s.PickStartCity()
		// if there is no area for the new alien to start
	}

	if a.IsDead() {
		return nil
	}

	newcity := s.PickNextMove(a)
	if newcity == nil {
		if a.IsTrapped() {
			return nil

		}

	}
	// fmt.Printf("alien %s moves from city %s to city %s \n", a.Name, a.CurrentCity.Name, newcity.Name)
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
	if len(Indexes) == 0 {
		return nil
	}
	randomIndex := rand.Intn(len(Indexes))

	return s.Cities[Indexes[randomIndex]]
}

func (s *Simulation) FinalOutcome() {
	for _, city := range s.Cities {
		if city.IfDestroyed() {
			msg := "City " + city.Name + " was destroyed by: "
			for alienName := range city.AlienPresent {
				msg += alienName + " "
			}
			fmt.Println(msg)
		} else {
			fmt.Printf("City %s is safe\n", city.Name)
		}
	}
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
