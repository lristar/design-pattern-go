package status

import "testing"

func TestStatus(t *testing.T) {
	cola := &Good{
		name:  "可乐",
		value: 3,
		count: 3,
	}
	coffee := &Good{
		name:  "咖啡",
		value: 6,
		count: 1,
	}
	mike := &Good{
		name:  "牛奶",
		value: 4,
		count: 0,
	}
	machine := NewMachine()
	machine.AddGood(cola)
	machine.AddGood(coffee)
	machine.AddGood(mike)
	machine.Prepare()
	machine.SelectGood(1, 1)
	machine.Run()
	machine.GiveMoney(10)
	machine.Run()
	machine.Run()
	machine.GetAllGood()
}
