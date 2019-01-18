package factory_method

import "fmt"

// 工厂方法模式是一种常用的类创建型设计模式,此模式的核心精神是封装类中不变的部分，提取其中个性化善变的部分为独立类，
// 通过依赖注入以达到解耦、复用和方便后期维护拓展的目的。

// 针对每一个产品提供一个工厂，通过不同的工厂来创建不同的产品。

// 以灯泡(Bulb)和电子管(Tube)为例
// 将公共的方法封装到接口中
type Light interface {
	// 相同的方法
	TurnOn()
	TurnOff()
	// 不同的方法
	UseFor()
}

// 将相同的部分：如需要玻璃和钨丝制造，封装到一个类中。
// 材料
type Material struct {
	glass    string // 玻璃
	tungsten string // 钨丝
}

func (m *Material) TurnOn() {
	fmt.Println("turn on")
}

func (m *Material) TurnOff() {
	fmt.Println("turn off")
}

// 分别实现各自特殊的部分
// 灯泡
type Bulb struct {
	Material
}

func (b *Bulb) UseFor() {
	fmt.Println("Bulbs use for lighting")
}

// 电子管
type Tube struct {
	Material
}

func (t *Tube) UseFor() {
	fmt.Println("Tubes use for machine")
}

// 抽象工厂
type Factory interface {
	Create() Light
}

// 灯泡工厂
type BulbFactory struct{}

func (bf *BulbFactory) Create() Light {
	return &Bulb{}
}

// 电子管工厂
type TubeFactory struct{}

func (tf *TubeFactory) Create() Light {
	return &Tube{}
}

// 四种角色：
// 抽象工厂：Factory
// 具体工厂：BulbFactory、TubeFactory
// 抽象产品：Light
// 具体产品：Bulb、Tube
//
