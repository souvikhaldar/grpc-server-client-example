package main

import (
	"context"
	"log"
	"net"

	ck "github.com/souvikhaldar/grpc-server-client/check"
	"google.golang.org/grpc"
)

const (
	port = ":8192"
)

type server struct {
}

func (s *server) Pong(ctx context.Context, a *ck.Empty) (*ck.Response, error) {
	return &ck.Response{Resp: "pong"}, nil
}

func main() {
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Unable to listen for requests: ", err)
	}
	serv := grpc.NewServer()
	ck.RegisterCheckServer(serv, &server{})
	serv.Serve(conn)
}
