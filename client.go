package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "./pb"
)

func main() {
	conn, err := grpc.Dial(":8989", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	msg := "hello123"
	cli := pb.NewGreetClient(conn)
	req := pb.HelloReq{
		Msg: &msg,
	}

	rsp, err := cli.Greet(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(*rsp.Rsp, rsp.Amount)
}
