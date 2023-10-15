package main

import (
	"fmt"
	"log"
	"net"

	cache "github.com/swastik959/distributed_chache-/chache"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	ServerOpts
	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}

}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		x := fmt.Errorf("listen error: %s", err)
		return x
	}
	log.Printf("server starting on port [%s]\n ", s.ListenAddr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept err: %s\n", err)
			continue
		}
		go s.handleConn(conn)
	}

}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("read err: %s\n", err)
			break
		}
		msg := buf[:n]
		fmt.Println(string(msg))
	}
}
