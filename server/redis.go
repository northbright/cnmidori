package server

import (
	"encoding/json"
	//"fmt"
	"path"
	//"github.com/garyburd/redigo/redis"
	"io/ioutil"
)

const (
	RedisSettingsFile string = "./redis-settings.json"
)

var (
// Redis Pool Map
// Key: Redis Server Name, Value: *redis.Pool.
//redisPools map[string]*redis.Pool
)

type RedisNode struct {
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

type RedisSettings struct {
	Servers []RedisNode `json:"servers"`
}

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

func init() {

}
