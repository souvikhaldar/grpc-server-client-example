# grpc-server-client-example
This is an example grpc server and client written in golang for learning purpose. Others can use it as a reference for learning as well.

# Prerequisite
Follow [this link](https://grpc.io/docs/quickstart/go/) for set up basic grpc setup. You can ignore the example mentioned.

# Operation
1) `cd /server`
2) `go run main.go`
3) `cd /client`
4) `go run main.go`
You can see a `pong` response, since the client in background made a `ping` call to the server.
