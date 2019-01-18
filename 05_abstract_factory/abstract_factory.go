package abstract_factory

import "fmt"

// 抽象产品
// 水果
type Fruit interface {
	GrewFruit()
}

// 蔬菜
type Vegetable interface {
	GrewVegetable()
}

// 抽象工厂
type Factory interface {
	MakeFruit() Fruit
	MakeVegetable() Vegetable
}

// 具体产品
// 苹果
type Apple struct{}

func (a *Apple) GrewFruit() {
	fmt.Println("grew apple")
}

// 香蕉
type Banana struct{}

func (b *Banana) GrewFruit() {
	fmt.Println("grew banana")
}

// 菠菜
type Spinach struct{}

func (s *Spinach) GrewVegetable() {
	fmt.Println("grew spinach")
}

// 花生
type Peanut struct{}

func (p *Peanut) GrewVegetable() {
	fmt.Println("grew peanut")
}

// 具体工厂
// 苹果菠菜工厂
type AppleSpinachFactory struct{}

func (asf *AppleSpinachFactory) MakeFruit() Fruit {
	fmt.Println("make apples")
	return &Apple{}
}

func (asf *AppleSpinachFactory) MakeVegetable() Vegetable {
	fmt.Println("make spinach")
	return &Spinach{}
}

// 香蕉花生工厂
type BananaPeanutFactory struct{}

func (bpf *BananaPeanutFactory) MakeFruit() Fruit {
	fmt.Println("make bananas")
	return &Banana{}
}

func (bpf *BananaPeanutFactory) MakeVegetable() Vegetable {
	fmt.Println("make peanuts")
	return &Peanut{}
}
