package main

import (
	"fmt"

	"github.com/mhaatha/go-learn/functional-option-pattern/challenge"
)

type Server struct {
	Host     string
	Port     int
	Timeout  int
	MaxConn  int
	Protocol string
}

type Option func(*Server)

func WithTimeout(t int) Option {
	return func(s *Server) {
		s.MaxConn = t
	}
}

func WithMaxConn(c int) Option {
	return func(s *Server) {
		s.MaxConn = c
	}
}

func WithProtocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func NewServer(host string, port int, opts ...Option) *Server {
	svr := &Server{
		Host:    host,
		Port:    port,
		Timeout: 30,
		MaxConn: 100,
	}

	for _, opt := range opts {
		opt(svr)
	}

	return svr
}

func main() {
	s1 := NewServer("localhost", 8080)
	fmt.Printf("Server 1: %+v\n", s1)

	s2 := NewServer("localhost", 9000, WithTimeout(60))
	fmt.Printf("Server 2: %+v\n", s2)

	s3 := NewServer("localhost", 80, WithMaxConn(500), WithTimeout(10))
	fmt.Printf("Server 3: %+v\n", s3)

	s4 := NewServer("localhost", 80, WithMaxConn(500), WithTimeout(10), WithProtocol("HTTPS"))
	fmt.Printf("Server 3: %+v\n", s4)

	r1 := challenge.NewRequest("https://api.google.com")
	fmt.Printf("R1: %+v\n", r1)

	r2 := challenge.NewRequest("https://api.google.com/users",
		challenge.WithMethod("POST"),
		challenge.WithBody(`{"name": "user"}`),
	)
	fmt.Printf("R2: Method=%s, Body=%s\n", r2.Method, r2.Body)

	r3 := challenge.NewRequest("https://api.google.com/secret",
		challenge.WithHeader("Authorization", "Bearer 123"),
		challenge.WithHeader("Content-Type", "application/json"),
	)
	fmt.Printf("R3 Headers: %+v\n", r3.Headers)
}
