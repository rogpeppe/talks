package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/kr/pretty"
	"google.golang.org/grpc"
)

func doClient() error {
	serverAddr := fmt.Sprintf("localhost:%d", *port)
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("dial: %v", err)
	}
	defer conn.Close()
	client := NewChatClient(conn)
	ctx := context.Background()
	users, err := client.Who(ctx, &Empty{})
	if err != nil {
		return fmt.Errorf("cannot get users: %v", err)
	}
	pretty.Println("users: ", users)
	go func() {
		myName := "joebloggs"
		tok, err := client.NewUser(ctx, &Username{
			Name: myName,
		})
		if err != nil {
			log.Printf("cannot make new user: %v", err)
			return
		}
		pretty.Println("NewUserResponse:", tok)
		for i := 0; i < 30; i++ {
			_, err := client.Say(ctx, &User2Server{
				Username: myName,
				Token:    tok.Token,
				Text:     fmt.Sprint("whee", i),
			})
			if err != nil {
				log.Printf("cannot send say message: %v", err)
				return
			}
			time.Sleep(time.Second / 2)
		}
	}()
	go func() {
		listener, err := client.Listen(ctx, &Empty{})
		if err != nil {
			log.Printf("cannot listen: %v", err)
		}
		for {
			m, err := listener.Recv()
			if err != nil {
				if err == io.EOF {
					log.Printf("end of stream")
					return
				}
				log.Printf("got error receiving message: %v", err)
				return
			}
			log.Printf("got message from server: %#v", m)
		}
	}()
	select {}
}
