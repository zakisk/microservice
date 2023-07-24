package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"

	protos "github.com/zakisk/microservice/currency/protos/currency"
	"github.com/zakisk/microservice/product-api/data"
	"github.com/zakisk/microservice/product-api/handlers"
)

func main() {
	l := hclog.Default()
	v := data.NewValidation()

	currencyServiceName := "currency-service"

	// Resolve the currency service using the Kubernetes DNS-based service discovery
	currencyServiceAddress := fmt.Sprintf("%s.meshery.svc.cluster.local:9092", currencyServiceName)

	// Connect to the gRPC server running in the currency service
	conn, err := grpc.Dial(currencyServiceAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	cc := protos.NewCurrencyClient(conn)

	// create data.ProductsDB instance
	pdb := data.NewProductsDB(cc, l)

	//create the new handler
	ph := handlers.NewProducts(l, v, pdb)

	//creating our new ServeMux
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/products", ph.ListProducts).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products", ph.ListProducts)

	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingelProduct).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingelProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", ph.Update)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.POST)
	postRouter.Use(ph.MiddlewareValidateProduct)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DELETE)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	//creating my own server in order to set the fields as per my requirement
	s := &http.Server{
		Addr:         ":9090",
		Handler:      ch(sm),
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Info("Got signal: ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
