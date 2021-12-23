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

type City struct {
	Name         string
	Links        map[string]*City
	Status       bool
	AlienPresent map[string]*Alien
}

func InitCity(name string) City {
	// generates a new city with a given name
	return City{
		Name:         name,
		Links:        make(map[string]*City),
		Status:       false,
		AlienPresent: make(map[string]*Alien),
	}
}

func (c *City) ConnectCity(direction string, toCity *City) error {
	// connects a exisiting city to a new city in a given direction
	if toCity.IfDestroyed() {
		return errors.New("City is destroyed")
	}
	if c.Links[direction] != nil {
		if c.Links[direction].Name != toCity.Name {
			return fmt.Errorf("Link between city %s and city %s in direction %s already exist", c.Name, c.Links[direction].Name, direction)
		}
	}
	c.Links[direction] = toCity
	return nil
}

func (c *City) IfDestroyed() bool {
	return c.Status
}

func (c *City) DestoryCity() {

	// var alien *Alien

	for alienName := range c.AlienPresent {
		alien := c.AlienPresent[alienName]
		alien.Kill()
	}
	for direction := range c.Links {
		delete(c.Links[direction].Links, OppositeDirections[direction])
	}

	c.Status = true

}
