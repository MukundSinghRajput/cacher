package tests

import (
	"reflect"
	"testing"
	"time"

	"github.com/MukundSinghRajput/cacher"
)

func TestSet(t *testing.T) {
	cache := cacher.NewCacher[string, string]()

	cache.Set("pro", "mukund", 5*time.Second)
	g, _ := cache.Get("pro")

	if !reflect.DeepEqual("mukund", g) {
		t.Errorf("Get() wrong return value. excepter = %v, got = %v", "mukund", g)
	}
}

func TestHas(t *testing.T) {
	cache := cacher.NewCacher[string, string]()

	cache.Set("pro", "mukund")

	ok := cache.Has("pro")

	if !ok {
		t.Error("Has() returned false excepted true")
	}
}

func TestGetAll(t *testing.T) {
	cache := cacher.NewCacher[string, string]()

	cache.Set("pro", "mukund")

	m := cache.GetAll()

	if len(m) == 0 {
		t.Error("GetAll() returned map with 0 length")
	}
}
