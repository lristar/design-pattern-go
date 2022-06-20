package strategy

import "testing"

func TestStrategy(t *testing.T) {
	lfu := &lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")
}