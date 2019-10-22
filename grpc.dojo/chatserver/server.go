package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc/peer"
)

type user struct {
	name           string
	token          string
	serverAddr     string
	serverProtocol string
}

// server implements the ChatServer interface.
type server struct {
	mu        sync.Mutex
	users     map[string]*user
	listeners map[chan<- *Server2User]bool
}

func newServer() *server {
	return &server{
		users:     make(map[string]*user),
		listeners: make(map[chan<- *Server2User]bool),
	}
}

func (s *server) NewUser(ctx context.Context, u *Username) (*NewUserResponse, error) {
	peer, peerOK := peer.FromContext(ctx)
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.users[u.Name] != nil {
		return nil, fmt.Errorf("user name already registered; try another one")
	}
	var buf [16]byte
	if _, err := rand.Read(buf[:]); err != nil {
		return nil, err
	}
	tok := fmt.Sprintf("%x", buf)
	s.users[u.Name] = &user{
		name:  u.Name,
		token: tok,
	}
	var addr string
	if peerOK {
		addr = peer.Addr.(*net.TCPAddr).IP.String()
	}
	return &NewUserResponse{
		Token:  tok,
		IPAddr: addr,
	}, nil
}

func (s *server) AnnounceRPCServer(ctx context.Context, m *ServerAnnouncement) (*Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, err := s.authUser(m.Username, m.Token)
	if err != nil {
		return nil, err
	}
	u.serverAddr, u.serverProtocol = m.ServerAddr, m.ServerProtocol
	return &Empty{}, nil
}

// TODO return number of listeners that received the message
func (s *server) Say(ctx context.Context, m *User2Server) (*Empty, error) {
	_, err := s.say(m)
	if err != nil {
		return nil, err
	}
	return &Empty{}, nil
}

func (s *server) Listen(_ *Empty, srv Chat_ListenServer) error {
	ctx := srv.Context()
	ch, stop := s.listen()
	defer stop()
	for {
		select {
		case m := <-ch:
			if err := srv.Send(m); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *server) Who(context.Context, *Empty) (*UserList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var ul UserList
	for _, u := range s.users {
		ul.Users = append(ul.Users, &User{
			Name:           u.name,
			ServerAddr:     u.serverAddr,
			ServerProtocol: u.serverProtocol,
		})
	}
	return &ul, nil
}

func (s *server) listen() (<-chan *Server2User, func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ch := make(chan *Server2User)
	s.listeners[ch] = true
	return ch, func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		delete(s.listeners, ch)
	}
}

func (s *server) say(m *User2Server) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, err := s.authUser(m.Username, m.Token); err != nil {
		fmt.Printf("bad auth for %q\n", m.Username)
		return 0, err
	}
	inm := &Server2User{
		Text:     m.Text,
		Username: m.Username,
	}
	fmt.Printf("user %q says: %q\n", m.Username, m.Text)
	n := 0
	for lis := range s.listeners {
		select {
		case lis <- inm:
			n++
		default:
			log.Printf("discarded message")
		}
	}
	return n, nil
}

func (s *server) authUser(username, token string) (*user, error) {
	u, ok := s.users[username]
	if !ok {
		return nil, fmt.Errorf("unknown username")
	}
	if token != u.token {
		return nil, fmt.Errorf("bad username token")
	}
	return u, nil
}
