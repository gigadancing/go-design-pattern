package adapter

import "fmt"

// 适配器模式
// 将一个类的接口转换成客户希望的另外一个接口。Adapter模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。

// 客户期待的目标（可以是类或接口）
type Target interface {
	Request()
}

// 需要适配的类
type Adaptee struct{}

func NewAdaptee() *Adaptee {
	return &Adaptee{}
}

func (a *Adaptee) SpecificRequest() {
	fmt.Println("Adaptee: SpecificRequest")
}

// 适配器
// 通过在内部包装一个Adaptee对象，将源接口转换成目标接口
type Adaptor struct {
	adaptee *Adaptee
}

func (a *Adaptor) Request() {
	a.adaptee.SpecificRequest()
}

func NewAdaptor(adaptee *Adaptee) *Adaptor {
	return &Adaptor{
		adaptee: adaptee,
	}
}
