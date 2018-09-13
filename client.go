package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/litao44/rpctest/pb"
)

func main() {
	conn, err := grpc.Dial(":8989", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := pb.NewGreetClient(conn)
	req := pb.HelloReq{
		Msg:    "hello123",
		Amount: 2,
	}

	rsp, err := cli.Greet(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Rsp)
}
