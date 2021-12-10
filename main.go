package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"midepeter/devtest/api"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	l := log.New()
	l.Println("serving HTTP...")
	sm := mux.NewRouter()

	router := sm.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/price", api.Pricehandler)
	router.HandleFunc("/pricews", api.PricehandlerWs)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan

	l.Println("gracefully shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
