package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"runtime"
)

type conf struct {
	BasicName string `yaml:"name"`
	BasicHost string `yaml:"host"`
	MySQLUrl  string `yaml:"mysql"`
	RedisUrl  string `yaml:"redis"`
	//RegistryAddr = "192.168.10.130:8500"
	RegistryAddr string `yaml:"registry"`
	TracerAddr   string `yaml:"tracer"`
}
type service struct {
	Name   string                 `yaml:"name"`
	Addr   string                 `yaml:"addr"`
	Others map[string]interface{} `yaml:"others"`
}

type Config struct {
	BasicConfig *conf               `yaml:"basic"`
	Services    map[string]*service `yaml:"services"`
}

var Conf *Config

func init() {
	var path string
	if runtime.GOOS == "linux" {
		path = "config/conf.yaml"
	} else {
		path = "config/conf-win.yaml"
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("read config file error")
	}
	err = yaml.Unmarshal(file, &Conf)
	if err != nil {
		log.Println("unmarshal yaml file error")
	}
}
