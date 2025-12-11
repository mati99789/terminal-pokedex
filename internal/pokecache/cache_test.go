package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestAddCache(t *testing.T) {
	cache := NewCache(time.Minute * 5)

	key := "test-key"
	val := []byte("test")

	cache.Add(key, val)

	retrived, ok := cache.Get(key)

	if !ok {
		t.Errorf("can't find key: %s in cache", key)
	}

	if !bytes.Equal(val, retrived) {
		t.Errorf("Retrived data: %v is not equal added %v ", retrived, val)
	}

	_, ok = cache.Get("non-existent-key")
	if ok {
		t.Error("Expected to not find non-existent key, but found it")
	}

}

func TestReapLoop(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)

	_, ok := cache.Get(key)
	if !ok {
		t.Error("Expected to find fresh entry")
	}

	time.Sleep(interval + 5*time.Millisecond)

	_, ok = cache.Get(key)
	if ok {
		t.Error("Expected entry to be reaped, but it still exists")
	}
}
