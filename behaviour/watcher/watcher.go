package main

import "log"

type ISubscriber interface {
	register(os observer)
	deregister(os observer)
	notifyAll()
}

type observer interface {
	watching(string)
	getID() string
}

// 具体订阅者
type consumer struct {
	obs     map[string]observer
	name    string
	inStock bool
}

func NewConsumer(name string) *consumer {
	return &consumer{name: name, obs: make(map[string]observer, 0)}
}

func (c *consumer) register(os observer) {
	c.obs[os.getID()] = os
}

func (c *consumer) deregister(os observer) {
	delete(c.obs, os.getID())
}

func (c *consumer) notifyAll() {
	for _, observer := range c.obs {
		go observer.watching(c.name)
	}
}

type watcher struct {
	id  string
	msg chan string // 如果要多个消费者要多通道,这个需要重新编写
}

func (w *watcher) watching(itemName string) {
	for s := range w.msg {
		log.Printf("Sending email to customer %s for item %s-----msg %s\n", w.id, itemName, s)
	}
}

func (w *watcher) Send(msg string) {
	w.msg <- msg
}

func (w *watcher) getID() string {
	return w.id
}

func NewWatcher(id string) *watcher {
	return &watcher{
		id:  id,
		msg: make(chan string, 10),
	}
}
