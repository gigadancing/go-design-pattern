package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	asf := AppleSpinachFactory{}
	asf.MakeFruit()
	asf.MakeVegetable()
	bpf := BananaPeanutFactory{}
	bpf.MakeFruit()
	bpf.MakeVegetable()
}
