package cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestCacheGetSet(t *testing.T) {
	c := NewCache()

	c.Set("test", "value")
	if val, found := c.Get("test"); !found || val != "value" {
		t.Errorf("Cache Get failed, expected 'value', got %v", val)
	}
}

// Race condition
func TestCacheConcurrentAccessForRaceCondition(t *testing.T) {
	c := NewCache()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if val, found := c.Get(fmt.Sprintf("key%d", i)); !found || val != i {
				t.Errorf("Concurrent Set or Get failed for key%d, expected %d, got %v", i, i, val)
			}
		}(i)
	}
	wg.Wait()
}
