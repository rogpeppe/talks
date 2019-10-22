package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

//newUser(name, [server]) -> userhandle
//
//use userhandler as authentication for messages from that user
//
//task:
//	implement your own server for getting user info.
//	getInfo() -> user attrs
//

var port = flag.Int("p", 8888, "port to serve on")
var runClient = flag.Bool("client", false, "run client not server")

func main() {
	flag.Parse()
	if *runClient {
		if err := doClient(); err != nil {
			log.Fatalf("client: %v", err)
		}
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterChatServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
