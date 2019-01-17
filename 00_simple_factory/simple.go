package simple_factory

import "fmt"

// 抽象产品：水果
type Fruit interface {
	Name()
}

// 具体产品
// 苹果
type Apple struct{}

func (a *Apple) Name() {
	fmt.Println("Apple")
}

// 香蕉
type Banana struct{}

func (b *Banana) Name() {
	fmt.Println("Banana")
}

// 工厂
func NewFruit(t int) Fruit {
	switch t {
	case 1:
		return &Apple{}
	case 2:
		return &Banana{}
	}
	return nil
}
