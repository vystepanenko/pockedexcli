package pokedexcache

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Hour)

	if cache.cache == nil {
		t.Errorf("Expected cache to be initialized")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Hour)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("Expected %s to be found in cache", c.inputKey)
		}

		if string(actual) != string(c.inputVal) {
			t.Errorf("Expected %s, got %s", c.inputVal, actual)
		}
	}
}

func TestPurgeCache(t *testing.T) {
	interval := time.Millisecond * 10

	cache := NewCache(interval)

	key := "key1"
	cache.Add(key, []byte("val1"))

	time.Sleep(interval + time.Millisecond*5)

	if _, ok := cache.Get(key); ok {
		t.Errorf("Expected %s to be purged", key)
	}
}

func TestPurgeCacheFail(t *testing.T) {
	interval := time.Millisecond * 10

	cache := NewCache(interval)

	key := "key1"
	cache.Add(key, []byte("val1"))

	time.Sleep(interval / 2)

	if _, ok := cache.Get(key); !ok {
		t.Errorf("Expected %s to be in cache", key)
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
