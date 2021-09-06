package types

import (
	"errors"
)

// complementary compass to find the direction of links
// ex : a lies in north of b
// using complimentary directions to map > b lies in south of a
var ComplementaryDirections = map[string]string{
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
	return City{
		Name:         name,
		Links:        make(map[string]*City),
		Status:       false,
		AlienPresent: make(map[string]*Alien),
	}
}

func (c *City) ConnectCity(direction string, toCity *City) error {
	if toCity.IfDestroyed() {
		return errors.New("City is destroyed")
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
		delete(c.Links[direction].Links, ComplementaryDirections[direction])
	}

	c.Status = true

}
