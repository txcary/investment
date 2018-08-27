package config

import (
	"github.com/go-ini/ini"
	"github.com/txcary/investment/common/utils"
)

const (
	configFile string = "investment.ini"
)

type Config struct {
	configObj *ini.File	
}

var obj *Config

func (obj *Config)GetString(section string, key string) string {
	return obj.configObj.Section(section).Key(key).String()
}

func Instance() *Config {
	var err error
	if obj == nil {
		obj = new(Config)
		obj.configObj, err = ini.Load(utils.Gopath() + utils.Slash() + configFile)
		if err != nil {
			panic(err)
		}
	}
	return obj
}