package facade

import "fmt"

// facade将外部的请求分给适当的子系统处理

// 主系统对外统一的"外观"
// 运行接口
type MainSystem interface {
	Run()
}

// 主系统
type MainSystemImpl struct {
	sa *subSystemA
	sb *subSystemB
}

func NewMainSystemImpl() *MainSystemImpl {
	return &MainSystemImpl{
		sa: newSubsystemA(),
		sb: newSubsystemB(),
	}
}

// 主系统运行要依赖子系统A和B
// 子系统A先读数据，系统B再写数据
func (m *MainSystemImpl) Run() {
	m.sa.Read()
	m.sb.Write()
}

// 子系统A
type subSystemA struct{}

func newSubsystemA() *subSystemA {
	return &subSystemA{}
}

func (sa *subSystemA) Read() {
	fmt.Println("subsystem A start reading")
}

// 子系统B
type subSystemB struct{}

func newSubsystemB() *subSystemB {
	return &subSystemB{}
}

func (sb *subSystemB) Write() {
	fmt.Println("subsystem B start writing")
}
