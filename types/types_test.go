package types

import (
	"testing"
)

var alien1 Alien = InitAlien("alienx")
var firstCity City = InitCity("firstcity")
var secondCity City = InitCity("secondcity")

func TestAlienAndCityStructs(t *testing.T) {
	if err := firstCity.ConnectCity("north", &secondCity); err != nil {
		t.Error(err)
	}

	if err := firstCity.ConnectCity("something", &firstCity); err == nil {
		t.Error("self loop rejection test failed")
	}
	secondCity.ConnectCity("south", &firstCity)
	alien1.InvadeCity(&firstCity)

	if alien1.CurrentCity.Name != "firstcity" {
		t.Error("city name do not match")
	}
	if alien1.CurrentCity.IfDestroyed() {
		t.Error("just initialized city destoryed")
	}

	if firstCity.Links["north"].Name != secondCity.Name {
		t.Error("links do not work for second city")
	}
	if secondCity.Links["south"].Name != firstCity.Name {
		t.Error("links do not work for first city")
	}

	firstCity.DestoryCity()

	if !alien1.IsDead() {
		t.Error("alien must be destroyed as the city has been destroyed")
	}
	if len(firstCity.Links) != 0 {
		t.Error("destoryed cities should not have links to other cities")
	}
	if len(secondCity.Links) != 0 {
		t.Error("as the first city is destoryed second city should have no links")
	}

}
