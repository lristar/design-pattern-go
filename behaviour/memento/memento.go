package memento

import "fmt"

type memento struct {
	version int
	content string
}

func (m *memento) getContent() string {
	return m.content
}

func (m *memento) getVersion() int {
	return m.version
}

// 这是一个状态机,产生记录
type originator struct {
	memento
}

func (e *originator) createMemento() *memento {
	return &memento{
		version: e.version,
		content: e.content,
	}
}

func (e *originator) restoreMemento(m *memento) {
	e.memento.content = m.getContent()
}

func (e *originator) setContent(content string) {
	e.content = content
	e.version++
}

// 保存记录
type caretaker struct {
	mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMemento() *memento {
	if !c.IsEmpty() {
		return c.mementoArray[len(c.mementoArray)-1]
	}
	return nil
}

func (c *caretaker) getCurContent() string {
	if !c.IsEmpty() {
		return c.mementoArray[len(c.mementoArray)-1].getContent()
	}
	return ""
}

func (c *caretaker) getCurVersion() int {
	if !c.IsEmpty() {
		return c.mementoArray[len(c.mementoArray)-1].getVersion()
	}
	return 0
}

func (c *caretaker) getMementoByIndex(index int) *memento {
	return c.mementoArray[index]
}

func (c *caretaker) IsEmpty() bool {
	if len(c.mementoArray) > 0 {
		return false
	}
	return true
}

func (c *caretaker) rollBack() {
	if len(c.mementoArray) > 1 {
		c.mementoArray = c.mementoArray[:len(c.mementoArray)-1]
	} else {
		fmt.Println("已到第一个版本,无法回滚")
	}
}
