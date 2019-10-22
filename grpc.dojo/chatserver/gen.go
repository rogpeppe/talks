package main

//go:generate protoc -I . --go_out=plugins=grpc:. ./chat.proto
