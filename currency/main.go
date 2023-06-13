package main

import (
	"net"
	"os"
	"os/signal"

	"github.com/hashicorp/go-hclog"
	protos "github.com/zakisk/microservice/currency/protos/currency"
	"github.com/zakisk/microservice/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	// creating new gRPC server
	gs := grpc.NewServer()

	// creating  new currency server
	cs := server.NewCurrency(log)

	// registering gRPC service with reflection
	reflection.Register(gs)

	protos.RegisterCurrencyServer(gs, cs)
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
	}

	// listening to gRPC server on port 9092
	go func() {
		err := gs.Serve(l)
		if err != nil {
			log.Error("Unable to listen to server", "error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Info("Got signal: ", sig)

	gs.Stop()
}
