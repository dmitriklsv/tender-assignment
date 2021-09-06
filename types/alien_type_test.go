package types

import "testing"

func TestAlien(t *testing.T) {
	name := "alienx"

	newAlien := InitAlien(name)
	// city := InitCity("somecity")

	if newAlien.Name != name {
		t.Errorf("wanted >%s, got >%s", newAlien.Name, name)
	}
	if newAlien.Trapped {
		t.Errorf("aliens should not be trapped in initial")
	}

	// newAlien.InvadeCity(&city)
	// if newAlien.CurrentCity.Name != city.Name {
	// 	t.Errorf("current and invaded city does not match")
	// }

}
