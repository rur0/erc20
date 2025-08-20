package erc20

import (
	"sync"
	"time"
)

// item is a struct that holds the value and the last access time
type item struct {
	value      interface{}
	lastAccess int64
}

type TTLMap struct {
	m map[string]*item

	mu sync.Mutex
}

func NewTTLMap(maxTTL int) *TTLMap {
	m := &TTLMap{m: make(map[string]*item)}

	// cleanup
	go func() {
		for now := range time.Tick(time.Minute) {
			m.mu.Lock()
			for k, v := range m.m {
				if now.Unix()-v.lastAccess > int64(maxTTL) {
					delete(m.m, k)
				}
			}
			m.mu.Unlock()
		}
	}()

	return m
}

// Put adds a new item to the map or updates the existing one
func (m *TTLMap) Put(k string, v interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	it, ok := m.m[k]
	if !ok {
		it = &item{
			value: v,
		}
	}
	it.value = v
	it.lastAccess = time.Now().Unix()
	m.m[k] = it
}

// Get returns the value of the given key if it exists
func (m *TTLMap) Get(k string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if it, ok := m.m[k]; ok {
		it.lastAccess = time.Now().Unix()
		return it.value, true
	}

	return nil, false
}

// Delete removes the item from the map
func (m *TTLMap) Delete(k string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.m, k)
}
