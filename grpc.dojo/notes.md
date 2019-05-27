# gRPC

The aim of this workshop is to provide a hands-on introduction to using
gRPC. First we'll write some code to act as a gRPC client for a very
simple chat server I've written; then you'll write your own server and
(hopefully!) talk to other servers that others in the workshop have
written.

## 1. Download the protobuf compiler

First ensure that you have Go 1.12 installed. You
*can* use an earlier version but Go modules make things easier.

You'll need to install the protobuf compiler.
Google for `golang protobuf quick start` or go to
https://grpc.io/docs/quickstart/go/.

You'll need both the protoc compiler (protoc) and the Go plugin
for it (github.com/golang/protobuf/protoc-gen-go)

## 2. Get the protobuf description of the server protocol

You can download this at:

https://raw.githubusercontent.com/rogpeppe/talks/master/grpc.dojo/chatserver/chat.proto

Short link: https://bit.ly/2YOBwws

## 3. Generate client code

You'll need to create a new directory for your gRPC module.
If you're using Go 1.12, either create the directory outside
of your GOPATH or set `GO111MODULE=on`.

Run `go mod init grpcexample` (or choose some other module name),
and create a file containing a `main` function so that the protobuf
code knows what package name to use on the generated files:

	package main
	
	func main() {
	}

Then copy the `.proto` file to that directory and run:

	go:generate protoc -I . --go_out=plugins=grpc:. ./chat.proto

So avoid typing that line every time, you can create a file with
a `go:generate` line in:

	package main
	
	//go:generate protoc -I . --go_out=plugins=grpc:. ./chat.proto

Then you can use `go generate` to re-generate the file.

You should now have some gRPC code generated for you!

## 4. Create a new user

If you look at the `chat.proto` file, you'll find a description of the chat
server. Your first objective is to write some client code that will
interact with the chat server. To do that, first create a new user
by calling the NewUser call.

That requires a client connection, which you can make with:

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := NewChatClient(conn)

You'll need to import the grpc package:

	import "google.golang.org/grpc"

Note that for the purposes of this workshop, all connections
are insecure. Don't do this in a real server!

A handy way to print out the results of RPC calls without
writing lots of code is to use Keith Rarick's `github.com/kr/pretty`
package.

The result of the `NewUser` call holds a token which you
need to use as a password to send messages with that
username. It also holds your IP address, which you'll be
using later to publish your own server.

## 5. Send a message

The `Say` call can be used to send a message to all users.  Write some
code to read lines from the standard input and send each line as a message
(use bufio.Scanner to read lines).

## 6. Listen to messages from other users

The `Listen` call returns a stream of messages.  Write some code to call
`Listen` and then repeatedly call `Recv` on the returned value to receive
messages and then print each message when it's received.

## 7. Define your own service

Now it's time to write your own RPC service.  Invent a very simple RPC
service, and write a `.proto` file that defines it. You'll be implementing
this, so best to keep to only one or two entry points to start with.

Run `protoc` to generate code for your service.

Notice that the generated code contains an `XXXServer` interface,
where `XXX` is the base name of your proto file.

To implement your server, you'll need to implement a type
with all the methods in that interface, and then
start a server listening on a network port. Assuming
your server implementation is myServer:

	lis, err := net.Listen("tcp", fmt.Sprintf(":8765))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterChatServer(grpcServer, myServer)
	grpcServer.Serve(lis)

That listens on port 8765 and serves your protocol.

## 8. Write some client code for your server

Check that your server works by writing some code to
make some calls to it.

## 9. Make your server available.

There is an AnnounceRPCServer in the chat server protocol
that you can use to send the address of your server
and the contents of the `.proto` file that you've written.

Use this to publish your server implementation to
the other workshop participants.

## 10. Use other people's servers.

Use the `Who` call to find out any other servers that
people have written, and try to call them.

If there are no other servers yet, leave yours running
and see if you can help out some other the other
participants!
