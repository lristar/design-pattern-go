package adapter

import "fmt"

type AdapterUsb struct {
	WindowMachine Window
}

func (a *AdapterUsb)insertIntoLightningPort(){
	fmt.Println("Adapter converts Lightning signal to USB.")
	a.WindowMachine.insertIntoUSBPort()
}

type computer interface {
	insertIntoLightningPort()
}

type Mac struct {

}

func (m *Mac)insertIntoLightningPort(){
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Window struct {

}

func (w *Window)insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

func StartComputer(m computer){
	m.insertIntoLightningPort()
}