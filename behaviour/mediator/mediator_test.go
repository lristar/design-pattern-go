package mediator

import "testing"

func TestTrain_GetMediator(t *testing.T) {
	manager := newStationManger()
	passengerTrain1 := &Train{
		md:   manager,
		name: "G1023",
	}
	passengerTrain2 := &freightTrain{
		Train{
			md:   manager,
			name: "G1024",
		},
	}
	passengerTrain1.arrive()
	passengerTrain2.arrive()
	passengerTrain1.depart()

}
