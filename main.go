package main

import (
	"fmt"
	"github.com/northbright/cnmidori/server"
)

func main() {
	var err error
	fmt.Printf("ServerRoot: %v\n", server.ServerRoot)

	// Initialize Redis Pools
	err = server.InitRedisPools()
	if err != nil {
		fmt.Printf("InitRedisPools() error: %v\n", err)
	}
}
