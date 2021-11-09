package main

import (
	"context"
	"fmt"

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

	res, err := cli.Set(context.Background(), &pbv.SetRequest{Key: key, Value: value})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BlockchainDB Set done.")
	fmt.Println(res.Tx)

	res1, err := cli.Get(context.Background(), &pbv.GetRequest{Key: key})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Blockchain Get done.")
	fmt.Println(string(res1.Value))
}