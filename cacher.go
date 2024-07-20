package cacher

import (
	"fmt"
	"sync"
	"time"
)

// Cache is a generic cache structure with a mutex for concurrent access.
type Cache[C comparable, T any] struct {
	mu   *sync.RWMutex
	data map[C]T
}

// NewCacher initializes a new Cache instance.
func NewCacher[C comparable, T any]() *Cache[C, T] {
	return &Cache[C, T]{
		mu:   &sync.RWMutex{},
		data: make(map[C]T),
	}
}

// Set adds a key-value pair to the cache with an optional time-to-live (ttl).
func (c *Cache[C, T]) Set(key C, value T, ttl ...time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value

	// If ttl is greater than 0, start a goroutine to delete the key after ttl expires.
	if len(ttl) > 0 && ttl[0] > 0 {
		go func() {
			<-time.After(ttl[0])
			c.mu.Lock()
			defer c.mu.Unlock()
			delete(c.data, key)
		}()
	}
	return nil
}

// Get retrieves a value from the cache by key. Returns an error if the key is not found.
func (c *Cache[C, T]) Get(key C) (T, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if val, ok := c.data[key]; !ok {
		var zero T
		return zero, fmt.Errorf("key %v not found", key)
	} else {
		return val, nil
	}
}

// Has checks if a key exists in the cache.
func (c *Cache[C, T]) Has(key C) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.data[key]
	return ok
}

// Delete removes a key from the cache. Returns an error if the key does not exist.
func (c *Cache[C, T]) Delete(key C) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[key]; !ok {
		return fmt.Errorf("the key %v doesn't exist", key)
	}
	delete(c.data, key)
	return nil
}

// GetAll returns all the data
func (c *Cache[C, T]) GetAll() map[C]T {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.data) == 0 {
		return map[C]T{}
	}

	return c.data
}
