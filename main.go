package main

import (
	"fmt"
	"github.com/northbright/cnmidori/server"
)

func main() {
	fmt.Printf("ServerRoot: %v\n", server.ServerRoot)
	s, err := server.LoadRedisSettings()
	if err != nil {
		fmt.Printf("Load Redis Setings Error: %v\n", err)
	}
	fmt.Printf("Redis Settings: %v\n", s)
}
