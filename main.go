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
	err := ds.Update()
	if err != nil {
		log.Fatal("[error] could not update data: %v", err)
	}

	// server
	cs := server.New(log, ds)
	gs := grpc.NewServer()
	// register
	currency.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	var root = fmt.Sprintf("%s:%d", "localhost", 10502)
	listen, err := net.Listen("tcp", root)
	if err != nil {
		log.Fatalf("listening on %s: %v", root, err)
	}

	log.Printf("[start] listening on %s", root)
	err = gs.Serve(listen)
	if err != nil {
		panic(err)
	}
}
