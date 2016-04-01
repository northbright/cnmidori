package server_test

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/northbright/cnmidori/server"
	"os"
)

var (
	RedisSettingsStr = `
{
    "servers":[
        {"name":"user", "addr":"localhost:6379", "password":"123456"},
        {"name":"data", "addr":"localhost:6380", "password":"123456"}
    ]
}`
)

// Run "go test"
func Example() {
	settings, err := server.LoadRedisSettingsFromStr(RedisSettingsStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "server.LoadSettingsFromStr() error: %v\n", err)
		return
	}

	pools, err := server.CreateRedisPools(settings)
	if err != nil {
		fmt.Fprintf(os.Stderr, "server.CreateRedisPools() error: %v\n", err)
		return
	}

	// Get Reids connection
	conn := pools["data"].Get()
	defer conn.Close()

	cmd := "\"SET server.name cnmidori\""
	if _, err = conn.Do("SET", "server.name", "cnmidori"); err != nil {
		fmt.Fprintf(os.Stderr, "%s error: %v\n", cmd, err)
	} else {
		fmt.Fprintf(os.Stderr, "%s OK\n", cmd)
	}

	cmd = "\"GET server.name\""
	if v, err := redis.String(conn.Do("GET", "server.name")); err != nil {
		fmt.Fprintf(os.Stderr, "%s error: %v\n", cmd, err)
	} else {
		fmt.Fprintf(os.Stderr, "%s OK: %v\n", cmd, v)
	}

	cmd = "\"DEL server.name\""
	if n, err := redis.Int64(conn.Do("DEL", "server.name")); err != nil {
		fmt.Fprintf(os.Stderr, "%s error: %v\n", cmd, err)
	} else {
		fmt.Fprintf(os.Stderr, "%s OK: deleted count = %v\n", cmd, n)
	}

	// Output:
}
