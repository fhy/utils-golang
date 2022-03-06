package config

import (
	"fmt"
	"os"
)

type Server struct {
	Listen       string `yaml:"listen"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"readtimeout"`
	WriteTimeout int    `yaml:"writetimeout"`
	Domain       string `yaml:"domain"`
}

type LogConfig struct {
	Path          string `yaml:"path"`
	RotationCount int    `yaml:"rotation_count"`
	Level         string `yaml:"level"`
}

type DbConfig struct {
	Path string `yaml:"path"`
}

type WeChatConfig struct {
	AppId     string `yaml:"id"`
	AppSecret string `yaml:"secret"`
}

func LoadConf(configFile string, config *interface{}) error {
	fmt.Printf("loading configfile: %s", configFile)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Print("can't find config path")
		os.Exit(1)
	} else {
		if err != nil {
			fmt.Print("Decode Config Error", err)
			os.Exit(1)
		}
	}
	return LoadConfYaml(configFile, config)
}
