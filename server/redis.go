package server

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"path"
	"time"
)

const (
	// Redis Settings File.
	RedisSettingsFile string = "./settings/redis-settings.json"

	// Default Redis Pool Settings.
	// For more information, see:
	// http://godoc.org/github.com/garyburd/redigo/redis#Pool
	DefMaxIdle        int = 3
	DefMaxActive      int = 1000
	DefIdleTimeoutSec int = 60 * 3
)

var (
	// Redis Pool Map
	// Key: Redis Server Name, Value: *redis.Pool.
	RedisPools map[string]*redis.Pool = map[string]*redis.Pool{}
)

// Redis Node in settings file.
type RedisNode struct {
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

// Redis Servers in settings file.
// Ex:
//{
//    "servers":[
//        {"name":"user", "addr":"localhost:6379", "password":"123456"},
//        {"name":"data", "addr":"localhost:6380", "password":"123456"}
//    ]
//}
type RedisSettings struct {
	Servers []RedisNode `json:"servers"`
}

// LoadRedisSettings() loads Redis settings from file.
func LoadRedisSettings() (s RedisSettings, err error) {
	buf := []byte{}
	s = RedisSettings{}

	p := path.Join(Dirs["ServerRoot"], RedisSettingsFile)

	if buf, err = ioutil.ReadFile(p); err != nil {
		return s, err
	}

	if err = json.Unmarshal(buf, &s); err != nil {
		return s, err
	}

	return s, nil
}

// NewRedisPool() creates a new Redis pool which matains a pool of connections.
func NewRedisPool(addr, password string, maxIdle, maxActive, idleTimeoutSec int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeoutSec) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Initialize Redis pools.
func InitRedisPools() (err error) {
	// Load Redis settings from file.
	s, err := LoadRedisSettings()
	if err != nil {
		fmt.Printf("LoadRedisSettings() error: %v\n", err)
		return err
	}

	// Create Redis Connection Pools
	for _, v := range s.Servers {
		RedisPools[v.Name] = NewRedisPool(v.Addr, v.Password, DefMaxIdle, DefMaxActive, DefIdleTimeoutSec)
	}

	return nil
}
