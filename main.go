package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/chutified/currencies/data"
	"github.com/chutified/currencies/protos/currency"
	"github.com/chutified/currencies/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// logging
	log := log.New(os.Stdin, "[CURRENCY SERVICE]", log.LstdFlags)

	// data service
	ds := data.New()

	// server
	cs := server.New(log, ds)
	gs := grpc.NewServer()
	// register
	currency.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	var root = fmt.Sprintf("%s:%d", "localhost", "8080")
	listen, err := net.Listen("tcp", root)
	if err != nil {
		log.Fatalf("listening on %s: %w", root, err)
	}
	err := gs.Serve(listen)
	if err != nil {
		panic(err)
	}
}
