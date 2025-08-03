package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Implement a counter structure that can be incremented in a concurrent environment

type MutexCounter struct {
	mu    sync.Mutex
	value int64
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup

	mutexCounter := &MutexCounter{}
	atomicCounter := &AtomicCounter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutexCounter.Inc()
			atomicCounter.Inc()
		}()
	}
	wg.Wait()

	fmt.Println(mutexCounter.Value())
	fmt.Println(atomicCounter.Value())
}
