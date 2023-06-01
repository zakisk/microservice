package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/zakisk/microservice/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api:", log.LstdFlags)

	//create the new handler
	ph := handlers.NewProducts(l)


	//creating our new ServeMux
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)
	
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	//creating my own server in order to set the fields as per my requirement
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	} ()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	
	sig := <-c
	l.Println("Got signal: ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	s.Shutdown(ctx)
}