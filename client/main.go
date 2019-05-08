package main

import (
	"context"
	"fmt"
	"log"

	ck "github.com/souvikhaldar/grpc-server-client/check"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8192"
)

func ping(client ck.CheckClient) {

	res, e := client.Pong(context.Background(), &ck.Empty{})
	if e != nil {
		log.Fatal("Error in calling pong: ", e)
	}
	fmt.Println(res.GetResp())
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect: ", err)
	}
	client := ck.NewCheckClient(conn)
	ping(client)

}
