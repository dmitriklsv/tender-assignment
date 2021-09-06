package types

import "errors"

type Alien struct {
	Name        string
	CurrentCity *City
	Trapped     bool
	Dead        bool
}

func InitAlien(name string) Alien {
	// at first all the aliens are alive and active

	newAlien := Alien{
		Name:        name,
		CurrentCity: nil,
		Dead:        false,
		Trapped:     false,
	}
	return newAlien

}

func (a *Alien) InvadeCity(city *City) error {
	delete(a.CurrentCity.AlienPresent, a.Name)
	if city.IfDestroyed() {
		return errors.New("this city has already been destroyed")
	}
	a.CurrentCity = city
	city.AlienPresent[a.Name] = a
	return nil
}

func (a *Alien) IsTrapped() bool {
	return a.Trapped
}
func (a *Alien) Trap() {
	a.Trapped = true
}

func (a *Alien) IsDead() bool {
	return a.Dead
}
func (a *Alien) Kill() {
	a.Dead = true

}
