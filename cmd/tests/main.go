package main

import (
	"context"
	"fmt"
	"time"

	pbv "github.com/sbip-sg/BlockchainDB/proto/blockchaindb"
	"google.golang.org/grpc"
)

func main() {
	addr := "127.0.0.1:50001"
	key := "tianwen1108"
	value := "66666666666666666666666666"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cli := pbv.NewBCdbNodeClient(conn)

	start := time.Now()
	res, err := cli.Set(context.Background(), &pbv.SetRequest{Key: key, Value: value})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BlockchainDB Set done.")
	fmt.Println(res.Tx)
	counter := 0
	for {
		counter += 1
		res1, err := cli.Get(context.Background(), &pbv.GetRequest{Key: key})
		if err != nil {
			fmt.Println(err)
			break
		}

		if string(res1.Value) != "" {
			fmt.Println("read query round ", counter)
		} else {
			fmt.Println("Blockchain Get done.")
			fmt.Println(string(res1.Value))
			elapsed := time.Since(start)
			fmt.Printf("Tx set-get took %s", elapsed)
			break
		}
	}

}
