package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/johnson7543/pricefetcher/server"
	"github.com/johnson7543/pricefetcher/service"

	"github.com/johnson7543/pricefetcher/client"
	"github.com/johnson7543/pricefetcher/proto"
)

func main() {

	// set default port number for -json and -grpc
	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4000", "listen address of the grpc transport")
	)
	flag.Parse()

	svc := service.NewLoggingService(service.NewPriceService())
	ctx := context.Background()

	grpcClient, err := client.NewGRPCClient(*grpcAddr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			time.Sleep(2 * time.Second)
			resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", resp)
		}

	}()

	go server.MakeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := server.NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()

}
