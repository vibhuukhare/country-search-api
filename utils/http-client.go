package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://restcountries.com/v3.1/name/"

var Logger = log.New(os.Stdout, "country-api: ", log.LstdFlags)

type CountryResponse struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Currency   string `json:"currency"`
	Population int    `json:"population"`
}

// FetchCountryDataFromExternalAPI will be called when the cache does not have the country's detail
// it will call the REST API countries
func FetchCountryDataFromExternalAPI(name string) (*CountryResponse, error) {
	client := http.Client{Timeout: 5 * time.Second}
	url := fmt.Sprintf("%s%s", baseURL, name)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching country data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response: %d", resp.StatusCode)
	}

	var countries []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if len(countries) == 0 {
		return nil, fmt.Errorf("no country data found")
	}

	country := countries[0]
	return &CountryResponse{
		Name:       country["name"].(map[string]interface{})["common"].(string),
		Capital:    country["capital"].([]interface{})[0].(string),
		Currency:   extractCurrencySymbol(country["currencies"]),
		Population: int(country["population"].(float64)),
	}, nil

}

func extractCurrencySymbol(currencies interface{}) string {
	for _, data := range currencies.(map[string]interface{}) {
		return data.(map[string]interface{})["symbol"].(string)
	}
	return ""
}
