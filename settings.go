package cnmidori

import (
	"encoding/json"
	"io/ioutil"
)

type RedisNode struct {
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

type Settings struct {
	RedisServers []RedisNode `json:"redis-servers"`
}

func NewSettingsFromBuffer(buf []byte) (settings *Settings, err error) {
	s := Settings{}
	if err = json.Unmarshal(buf, &s); err != nil {
		return &s, err
	}

	return &s, nil

}

func NewSettings(settingsFile string) (settings *Settings, err error) {
	buf := []byte{}
	if buf, err = ioutil.ReadFile(settingsFile); err != nil {
		return &Settings{}, err
	}

	return NewSettingsFromBuffer(buf)
}
