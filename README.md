# Country API 

A simple Go-based API to fetch country details, with caching for better performance.  

## Features
- Fetch country details like name, capital, currency, and population
- In-memory caching to reduce API calls
- Concurrency-safe cache handling
- Well-structured project with clean Go idioms

## Installation
1. Clone the repository  
   ```sh
   git clone https://github.com/yourusername/country-api.git
   cd country-api

2. Install dependencies
    go mod tidy
3. Run the application
    go run main.go

## Endpoints

1. Get country details
    curl -X GET "http://localhost:8080/api/countries/search?name=India"
2. Response
    {"name":"India","capital":"New Delhi","currency":"₹","population":1380004385}

## Technologies Used
1. Go
2. RESTCountries API
3. In-memory cache