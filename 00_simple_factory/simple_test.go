package simple_factory

import "testing"

func TestNewFruit(t *testing.T) {
	f1 := NewFruit(1)
	if f1 == nil {
		t.Fatal("new apple failed")
	}
	f1.Name()

	f2 := NewFruit(2)
	if f2 == nil {
		t.Fatal("new banana failed")
	}
	f2.Name()
}
