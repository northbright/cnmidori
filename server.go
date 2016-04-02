package cnmidori

import (
	"github.com/garyburd/redigo/redis"
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

type Server struct {
	Settings
	RedisPools map[string]*redis.Pool
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

func NewServer(settingsFile string) (server *Server, err error) {
	server = &Server{Settings{}, make(map[string]*redis.Pool)}
	settings, err := NewSettings(settingsFile)
	if err != nil {
		return nil, err
	}

	server.Settings = *settings
	for _, v := range server.RedisServers {
		server.RedisPools[v.Name] = NewRedisPool(v.Addr, v.Password, DefMaxIdle, DefMaxActive, DefIdleTimeoutSec)
	}

	return server, nil
}
