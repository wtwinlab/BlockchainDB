package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/sbip-sg/BlockchainDB/node/config"
	"github.com/sbip-sg/BlockchainDB/node/service"
	pbv "github.com/sbip-sg/BlockchainDB/proto/blockchaindb"
	"google.golang.org/grpc"
)

func main() {

	configFile := flag.String("config", "config.toml", "The path to the config file")
	flag.Parse()
	var conf config.Options
	err := config.ReadConfig(&conf, *configFile) //default config file "config.toml"
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config: %v\n", err)
		os.Exit(1)
	}

	s := grpc.NewServer()

	svr, err := service.NewServerNode(&conf)
	if err != nil {
		//panic(err)
		fmt.Printf("%v", err)

	}
	pbv.RegisterBCdbNodeServer(s, svr)
	lis, err := net.Listen("tcp", conf.ServerNodeAddr)

	if err != nil {
		//panic(err)
		fmt.Printf("%v", err)
	} else {
		log.Println("Veritas Node listen address: " + conf.ServerNodeAddr)
	}

	go func() {
		log.Printf("Veritas Node Serving gRPC: %s", conf.ServerNodeAddr)
		s.Serve(lis)
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	sig := <-ch
	log.Printf("Received signal %v, quiting gracefully", sig)
}
