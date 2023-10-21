package main

import (
	"log"
	"net"
	"time"

	cache "github.com/swastik959/distributed_chache-/chache"
)

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", opts.ListenAddr)
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("Set Foo Bar 2500"))
	}()
	server := NewServer(opts, cache.New())
	server.Start()

}
