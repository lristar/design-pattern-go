package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	consumer := NewConsumer("tv")
	watcher1 := NewWatcher("adaf")
	watcher2 := NewWatcher("zxcv")
	consumer.register(watcher2)
	consumer.register(watcher1)
	consumer.notifyAll()
	for {
		fmt.Println("请输入内容:")
		reader := bufio.NewReader(os.Stdin)

		//ReadBytes
		//当读取到‘\n’时，即一行完
		res, err := reader.ReadBytes('\n')
		if err != nil {
			panic("出错咯~~")
		}
		watcher1.Send(string(res))
		watcher2.Send(string(res))
	}
}
