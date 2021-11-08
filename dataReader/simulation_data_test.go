package dataReader

import "testing"

func TestReadCityData(t *testing.T) {
	path := "testData.txt"
	cities, err := ReadCityData(path)
	if err != nil {
		t.Error(err)
	}
	if len(*cities) != 3 {
		t.Error("all cities are not loaded")
	}

	for _, city := range *cities {
		for direction, toCity := range city.Links {
			if city.Name != toCity.Links[OppositeDirections[direction]].Name {
				t.Error("opposite direction link does not work")
			}
		}

	}

}
