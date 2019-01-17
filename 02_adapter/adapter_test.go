package adapter

import "testing"

func TestAdaptor(t *testing.T) {
	adaptee := NewAdaptee()
	adaptor := NewAdaptor(adaptee)
	adaptor.Request()
}
