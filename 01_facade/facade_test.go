package facade

import "testing"

func TestMainSystemImpl_Run(t *testing.T) {
	m := NewMainSystemImpl()
	m.Run()
}
