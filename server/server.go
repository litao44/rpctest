package main

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/litao44/rpctest/pb"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterGreetServer(s, &GreetServer{})
	listen, err := net.Listen("tcp", ":8989")
	if err != nil {
		panic(err)
	}

	s.Serve(listen)
}

type GreetServer struct{}

func (s *GreetServer) Greet(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	fmt.Printf("received msg: %s, amount: %d\n", req.Msg, req.Amount)
	return &pb.HelloResp{
		Rsp: strings.Repeat(req.Msg, int(req.Amount)),
	}, nil
}
