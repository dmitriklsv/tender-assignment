package types

import (
	"errors"
	"fmt"
)

// complementary compass to find the direction of links
// ex : a lies in north of b
// using complimentary directions to map > b lies in south of a
var OppositeDirections = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

// City structure
// Name : name of the city
// Links : contains map for direction->reference to the city (ex north->city<panda>)
// status: treu if the city is destroyed otherwise false
// AlienPresent: keeps tracks of reference of the aliens that are present in the city

type City struct {
	Name         string
	Links        map[string]*City
	status       bool
	AlienPresent map[string]*Alien
}

func InitCity(name string) City {
	// generates a new city with a given name
	return City{
		Name:         name,
		Links:        make(map[string]*City),
		status:       false,
		AlienPresent: make(map[string]*Alien),
	}
}

func (c *City) ConnectCity(direction string, toCity *City) error {
	// connects a exisiting city to a new city in a given direction
	// helps in generating the map

	// check to ensure that the city that we are connected to is not destroyed
	if toCity.Name == c.Name {
		return errors.New("self Loop is not allowed")
	}
	if toCity.IfDestroyed() {
		return errors.New("City is destroyed")
	}
	// if a link already exisits in the direction for a city
	if c.Links[direction] != nil {
		// if the preexisiting city and the new city is not the same then this is an error as we are trying to connect 2 cities in the same dirction
		if c.Links[direction].Name != toCity.Name {
			return fmt.Errorf("Link between city %s and city %s in direction %s already exist", c.Name, c.Links[direction].Name, direction)
		}
	}
	// link the toCity
	c.Links[direction] = toCity
	return nil
}

// checks id the city is destroyed
func (c *City) IfDestroyed() bool {
	return c.status
}

func (c *City) DestoryCity() {
	// when the city is destoryed then all the present aliens in the city is killed
	for alienName := range c.AlienPresent {
		alien := c.AlienPresent[alienName]
		alien.Kill()
	}
	// remove connections from 1 city to other

	// removing link from city2->city1
	// removing the link from city1->city2
	for direction := range c.Links {

		delete(c.Links[direction].Links, OppositeDirections[direction])
		delete(c.Links, direction)
	}
	// change the status of city to true
	c.status = true

}
