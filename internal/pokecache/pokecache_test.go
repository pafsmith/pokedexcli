package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		name string
		key  string
		val  []byte
	}{
		{
			name: "basic URL",
			key:  "https://testdomain.com",
			val:  []byte("testdata"),
		},
		{
			name: "URL with path",
			key:  "https://testdomain.com/path",
			val:  []byte("moretestdata"),
		},
		{
			name: "empty value",
			key:  "https://testdomain.com/empty",
			val:  []byte(""),
		},
		{
			name: "large value",
			key:  "https://testdomain.com/large",
			val:  []byte("this is a much longer piece of data that simulates a larger API response"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key %q", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected value %q, got %q", string(c.val), string(val))
				return
			}
		})
	}
}

func TestGetNonExistent(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	_, ok := cache.Get("nonexistent")
	if ok {
		t.Errorf("expected to not find nonexistent key")
	}
}

func TestCacheOverwrite(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	key := "https://testdomain.com"
	cache.Add(key, []byte("original"))
	cache.Add(key, []byte("updated"))

	val, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	if string(val) != "updated" {
		t.Errorf("expected updated value, got %q", string(val))
	}
}

func TestCacheConcurrency(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(n int) {
			key := fmt.Sprintf("key-%d", n)
			val := []byte(fmt.Sprintf("value-%d", n))
			cache.Add(key, val)
			retrieved, ok := cache.Get(key)
			if !ok || string(retrieved) != string(val) {
				t.Errorf("concurrent operation failed for %s", key)
			}
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://testdomain.com", []byte("testdata"))

	_, ok := cache.Get("https://testdomain.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://testdomain.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
