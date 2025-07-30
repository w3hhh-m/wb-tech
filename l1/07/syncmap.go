package main

import (
	"fmt"
	"sync"
)

// Create a map data structure where multiple goroutines can safely write
// and read from it at the same time without causing data races or crashes

// SyncMap is a custom sync.Map
type SyncMap struct {
	sync.Mutex
	m map[any]any
}

func NewSyncMap() *SyncMap {
	return &SyncMap{m: make(map[any]any)}
}

func (sm *SyncMap) Store(key any, value any) {
	sm.Lock()
	defer sm.Unlock()
	sm.m[key] = value
}

func (sm *SyncMap) Load(key any) (any, bool) {
	sm.Lock()
	defer sm.Unlock()
	value, ok := sm.m[key]
	return value, ok
}

func main() {
	custom := NewSyncMap()
	builtin := sync.Map{}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				custom.Store(j, "custom")
				val, ok := custom.Load(j)
				if ok {
					fmt.Printf("Custom map[%v] value: %v\n", j, val)
				} else {
					fmt.Printf("Key %v not found in custom map\n", j)
				}
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				builtin.Store(j, "builtin")
				val, ok := builtin.Load(j)
				if ok {
					fmt.Printf("Builtin map[%v] value: %v\n", j, val)
				} else {
					fmt.Printf("Key %v not found in builtin map\n", j)
				}
			}
		}()
	}

	wg.Wait()
}
