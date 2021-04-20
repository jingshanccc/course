package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
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
	path := "https://gitee.com/jingshanccc/course/raw/main/public/config/"
	is := false
	if runtime.GOOS == "windows" {
		path += "conf-win.yaml"
	} else {
		is = true
		path += "conf.yaml"
	}
	file, err := http.Get(path)
	if err != nil {
		log.Println("read config file error, err is " + err.Error())
	}
	defer file.Body.Close()
	bytes, err := ioutil.ReadAll(file.Body)
	err = yaml.Unmarshal(bytes, &Conf)
	if err != nil {
		log.Println("unmarshal yaml file error, err is " + err.Error())
	}
	if is { // edit file path for not windows
		dir, _ := os.Getwd()
		Conf.Services["file"].Others["filePath"] = strings.Replace(dir, "gateway", "file", 1) + Conf.Services["file"].Others["filePath"].(string)
		Conf.Services["user"].Others["emailTemplatePath"] = dir + Conf.Services["user"].Others["emailTemplatePath"].(string)
	}
}
