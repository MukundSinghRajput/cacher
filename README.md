# Cacher

The Cacher package provides a simple, generic caching structure for Go. It includes a mutex for concurrent access and functions for setting, getting, checking for the existence of, and deleting cache entries.

## Cache Structure

The `Cache` structure is a generic cache that takes two type parameters:

- `C`: the type of the keys in the cache
- `T`: the type of the values in the cache

Cache includes a sync.RWMutex for concurrent access and a map of key-value pairs.

## Functions

- `NewCacher[C comparable, T any]() *Cache[C, T]`: initializes a new `Cache` instance with a read-write mutex and an empty map of key-value pairs.
- `Set(key C, value T, ttl ...time.Duration)`: adds a key-value pair to the cache with an optional time-to-live (ttl). If ttl is specified and is greater than zero, a new goroutine is started to remove the key-value pair after ttl expires.
- `Get(key C) (T, error)`: retrieves a value from the cache by key. Returns an error if the key is not found.
- `Has(key C) bool`: checks if a key exists in the cache.
- `Delete(key C) error`: removes a key from the cache. Returns an error if the key does not exist.
- `GetAll() map[C]T`: retrieves all key-value pairs from the cache.

## Example Usage

```
package main

import (
 "fmt"
 "time"

 "github.com/[your-username]/cacher"
)

func main() {
 c := cacher.NewCacher[string, int]()

 err := c.Set("key1", 123, 5*time.Second)
 if err != nil {
  fmt.Println(err)
  return
 }

 val, err := c.Get("key1")
 if err != nil {
  fmt.Println(err)
  return
 }

 fmt.Println(val)

 time.Sleep(6 * time.Second)

 val, err = c.Get("key1")
 if err != nil {
  fmt.Println(err)
  return
 }

 fmt.Println(val)
}
```

This example sets the key "key1" with a value of 123 and a time-to-live of 5 seconds. After retrieving the value and printing it, the program waits for 6 seconds and then tries to retrieve the value again. This time, the key-value pair has expired, so the `Get` function returns an error.

## Testing

To run the tests, navigate to the `cacher` package directory and use the `make test` command.

```
cd cacher
make test
```

This will run the tests and output the results.

## Benchmark

To run the benchmark test, navigate to the `cacher` package directory and use the `make benchmark` command.

```
cd cacher
make benchmark
```

This will run the benchmark and output the results.