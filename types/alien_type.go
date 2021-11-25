package types

import "errors"

// Alien structure contains
// name : unique name of the alien
// CurrentCity: reference to the city the alien is in
// trapped : is the alien in a city that has no other connections
// dead : is the alien destroyed during combat

type Alien struct {
	Name        string
	CurrentCity *City
	trapped     bool
	dead        bool
}

func InitAlien(name string) Alien {
	// at first all the aliens are alive and active

	newAlien := Alien{
		Name:        name,
		CurrentCity: nil,
		dead:        false,
		trapped:     false,
	}
	return newAlien

}

// invade a new city
func (a *Alien) InvadeCity(city *City) error {
	// error if the city that the alien is going to that is going to be destroyed
	if city.IfDestroyed() {
		return errors.New("this city has already been destroyed")
	}
	// if the alien moves from one city to another then delete
	// the alien from the currentCity map that tracks the alien that are present in the city

	if a.CurrentCity != nil {
		delete(a.CurrentCity.AlienPresent, a.Name)
	}
	// change the current city and add the alien to the city map for keeping track
	a.CurrentCity = city
	city.AlienPresent[a.Name] = a
	return nil
}

// getter and setter for trap status
func (a *Alien) IsTrapped() bool {
	return a.trapped
}
func (a *Alien) Trap() {
	a.trapped = true
}

// getter and setter for checking if the alien is alive or dead

func (a *Alien) IsDead() bool {
	return a.dead
}
func (a *Alien) Kill() {
	a.dead = true

}
