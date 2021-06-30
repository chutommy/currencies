package main

import (
	"fmt"
	"log"
	"net"
	"os"

	config "github.com/chutommy/currencies/config"
	data "github.com/chutommy/currencies/data"
	currency "github.com/chutommy/currencies/protos/currency"
	server "github.com/chutommy/currencies/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log := log.New(os.Stdin, "[CURRENCY SERVICE]", log.LstdFlags)

	// configuration
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		log.Fatalf("[error] could not find config file: %v", err)
	}
	var addr = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	// data service
	ds := data.New()
	err = ds.Update("https://markets.businessinsider.com/currencies")
	if err != nil {
		log.Fatalf("[error] could not update data: %v", err)
	}

	// server
	cs := server.New(log, ds)
	gs := grpc.NewServer()
	// register
	currency.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("[error] listening on %s: %v", addr, err)
	}

	log.Printf("[start] listening on %s", addr)
	err = gs.Serve(listen)
	if err != nil {
		log.Panicf("[error] running server: %v", err)
	}
}
