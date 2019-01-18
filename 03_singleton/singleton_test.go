package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	ins1 := GetInstance1()
	ins2 := GetInstance2()
	ins3 := GetInstance3()
	ins4 := GetInstance4()

	if ins1 != ins2 || ins1 != ins3 || ins3 != ins4 {
		t.Fatal("instance is not equal")
	}

}
