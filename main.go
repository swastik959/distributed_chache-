package main

import (
	cache "github.com/swastik959/distributed_chache-/chache"
)

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	server := NewServer(opts, cache.New())
	server.Start()

}
