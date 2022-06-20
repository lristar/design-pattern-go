package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	caretaker := &caretaker{mementoArray: make([]*memento, 0)}
	origin := &originator{}
	origin.setContent("修改第一次")
	caretaker.addMemento(origin.createMemento())
	origin.setContent("修改第二次")
	caretaker.addMemento(origin.createMemento())
	origin.setContent("修改第三次")
	caretaker.addMemento(origin.createMemento())
	fmt.Println(caretaker.getMemento().getContent(), "version :", caretaker.getMemento().getVersion())
	caretaker.rollBack()
	fmt.Println(caretaker.getMemento().getContent(), "version :", caretaker.getMemento().getVersion())
	caretaker.rollBack()
	fmt.Println(caretaker.getMemento().getContent(), "version :", caretaker.getMemento().getVersion())
	caretaker.rollBack()
	fmt.Println(caretaker.getMemento().getContent(), "version :", caretaker.getMemento().getVersion())
	caretaker.rollBack()
}
