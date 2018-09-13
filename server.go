package main

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "./pb"
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
	return &pb.HelloResp{
		Rsp:    req.Msg,
		Amount: req.Amount,
	}, nil
}
