package tests

import (
	"testing"

	"github.com/MukundSinghRajput/cacher"
)

// BenchmarkCacheSet benchmarks the Set method of the Cache.
func BenchmarkCacheSet(b *testing.B) {
	cache := cacher.NewCacher[int, int]()
	for i := 0; i < b.N; i++ {
		cache.Set(i, i, 0)
	}
}

// BenchmarkCacheGet benchmarks the Get method of the Cache.
func BenchmarkCacheGet(b *testing.B) {
	cache := cacher.NewCacher[int, int]()
	for i := 0; i < b.N; i++ {
		cache.Set(i, i, 0)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i)
	}
}

// BenchmarkCacheDelete benchmarks the Delete method of the Cache.
func BenchmarkCacheDelete(b *testing.B) {
	cache := cacher.NewCacher[int, int]()
	for i := 0; i < b.N; i++ {
		cache.Set(i, i, 0)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Delete(i)
	}
}
