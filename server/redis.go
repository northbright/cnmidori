package server

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"time"
)

const (
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

// LoadRedisSettings() loads Redis settings from []byte(JSON string).
func LoadRedisSettings(buf []byte) (s RedisSettings, err error) {
	s = RedisSettings{}
	if err = json.Unmarshal(buf, &s); err != nil {
		return s, err
	}

	return s, nil
}

// LoadRedisSettingsFromStr() loads Redis settings from JSON string.
func LoadRedisSettingsFromStr(str string) (s RedisSettings, err error) {
	return LoadRedisSettings([]byte(str))
}

// LoadRedisSettingsFromFile() loads Redis settings from JSON file.
func LoadRedisSettingsFromFile(file string) (s RedisSettings, err error) {
	buf := []byte{}
	s = RedisSettings{}

	if buf, err = ioutil.ReadFile(file); err != nil {
		return s, err
	}

	return LoadRedisSettings(buf)
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

// CreateRedisPools() creates the Redis pools.
func CreateRedisPools(settings RedisSettings) (pools map[string]*redis.Pool, err error) {
	pools = map[string]*redis.Pool{}

	// Create Redis Connection Pools
	for _, v := range settings.Servers {
		pools[v.Name] = NewRedisPool(v.Addr, v.Password, DefMaxIdle, DefMaxActive, DefIdleTimeoutSec)
	}

	return pools, nil
}
