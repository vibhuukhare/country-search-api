package services

import (
	"sync"
	"testing"

	"github.com/vibhu.khare/country-api/cache"
	"github.com/vibhu.khare/country-api/utils"
)

func TestCountryServiceCacheHit(t *testing.T) {
	c := cache.NewCache()
	service := NewCountryService(c)

	// Preload cache
	c.Set("India", &utils.CountryResponse{Name: "India", Capital: "New Delhi", Currency: "₹", Population: 1380004385})

	country, err := service.GetCountryData("India")
	if err != nil || country.Name != "India" {
		t.Fatalf("Expected India, got %v", country)
	}
}

func TestCountryServiceCacheMiss(t *testing.T) {
	c := cache.NewCache()
	service := NewCountryService(c)

	country, err := service.GetCountryData("Brazil")
	if err != nil || country.Name != "Brazil" {
		t.Fatalf("Expected Brazil, got %v", country)
	}

	cachedCountry, _ := c.Get("Brazil")
	if cachedCountry == nil {
		t.Fatalf("Expected Brazil to be cached")
	}
}

// Race condition
func TestCountryServiceConcurrencyForRaceCondition(t *testing.T) {
	c := cache.NewCache()
	service := NewCountryService(c)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ { // 10 concurrent requests
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = service.GetCountryData("India")
		}()
	}
	wg.Wait()

	// Check if data is cached after concurrent calls
	if _, exists := c.Get("India"); !exists {
		t.Fatalf("Expected India to be cached")
	}
}
