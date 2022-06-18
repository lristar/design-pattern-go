package create

import (
	"testing"
	"time"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 10; i++ {
		go GetSingleInstance()
	}
	time.Sleep(20 * time.Second)
}

func TestSingletonOne(t *testing.T) {
	for i := 0; i < 10; i++ {
		go GetSingleInstanceOnce()
	}
	time.Sleep(20 * time.Second)
}
