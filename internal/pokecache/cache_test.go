package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	interval := 5 * time.Second
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
				t.Errorf("expected to find the key from cache")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find the value from cache")
				return
			}
		})
	}
}

func TestReadLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	const waitTime = baseTime + 5*time.Second
	cache := NewCache(baseTime)

	testCase := struct {
		key string
		val []byte
	}{
		key: "https://example.com",
		val: []byte("testdata"),
	}
	cache.Add(testCase.key, testCase.val)

	_, ok := cache.Get(testCase.key)
	if !ok {
		t.Error("expected to find value from cache")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(testCase.key)
	if ok {
		t.Errorf("expected to not find value cache")
		return
	}
}
