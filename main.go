package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/vibhu.khare/country-api/cache"
	"github.com/vibhu.khare/country-api/handlers"
	"github.com/vibhu.khare/country-api/services"
)

func main() {

	l := log.New(os.Stdout, "country-api", log.LstdFlags)

	c := cache.NewCache()
	service := services.NewCountryService(c)
	handler := handlers.NewCountryHandler(service)

	sm := http.NewServeMux()

	sm.HandleFunc("/api/countries/search", handler.SearchCountryName)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  time.Second * 120,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan

	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
