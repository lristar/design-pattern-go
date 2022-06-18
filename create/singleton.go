package create

import (
	"fmt"
	"sync"
)

var mx = &sync.Mutex{}

type Single struct {
}

var singleInstance *Single

func GetSingleInstance() *Single {
	if singleInstance == nil {
		mx.Lock()
		defer mx.Unlock()
		if singleInstance == nil {
			fmt.Println("更新实例")
			singleInstance = &Single{}
		} else {
			fmt.Println("实例已存在")
		}
	}
	return singleInstance
}

var one = &sync.Once{}

func GetSingleInstanceOnce() *Single {
	if singleInstance == nil {
		one.Do(
			func() {
				fmt.Println("新建实例")
				singleInstance = &Single{}
			})
	} else {
		fmt.Println("实例已存在")
	}
	return singleInstance
}
