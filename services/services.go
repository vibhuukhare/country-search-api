package services

import (
	"github.com/vibhu.khare/country-api/cache"
	"github.com/vibhu.khare/country-api/utils"
)

type CountryService struct {
	cache *cache.Cache
}

func NewCountryService(c *cache.Cache) *CountryService {
	return &CountryService{cache: c}
}

func (s *CountryService) GetCountryData(name string) (*utils.CountryResponse, error) {
	
	// Check in cache data, if found then will return the data stored in the cache
	if cachedData, found := s.cache.Get(name); found {
		utils.Logger.Println("Cache hit for", name)
		return cachedData.(*utils.CountryResponse), nil
	}

	// The cache does not have the specific country details
	// So we will fetch from the third party API
	country, err := utils.FetchCountryDataFromExternalAPI(name)
	if err != nil {
		return nil, err
	}
	utils.Logger.Println("Cache miss, fetching data for", name)
	// Store in cache after fetching from the third party api
	s.cache.Set(name, country)
	return country, nil
}
