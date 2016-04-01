package main

import (
	"fmt"
	"github.com/northbright/cnmidori/server"
)

func main() {
	var err error
	fmt.Printf("ServerRoot: %v\n", server.ServerRoot)

	// Load Redis Settings.
	redisSettings, err := server.LoadRedisSettingsFromFile(server.RedisSettingsFile)
	if err != nil {
		fmt.Printf("LoadRedisSettingsFromFile(%v) error: %v\n", server.RedisSettingsFile, err)
	} else {
		fmt.Printf("LoadRedisSettingsFromFile(%v) OK: %v\n", server.RedisSettingsFile, redisSettings)
	}

	// Create Redis Pools.
	server.RedisPools, err = server.CreateRedisPools(redisSettings)
	if err != nil {
		fmt.Printf("CreateRedisPools(%v) error: %v\n", redisSettings, err)
	} else {
		fmt.Printf("CreateRedisPools() OK\n")
	}
}
