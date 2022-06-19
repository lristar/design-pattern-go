package adapter

import "testing"

func TestAdapter(t *testing.T) {
	StartComputer(&Mac{})
	adapter := &AdapterUsb{WindowMachine:Window{}}
	StartComputer(adapter)
}