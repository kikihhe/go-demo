package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ConfigInfo *Config
)

func InitConf() {
	const ConfigFile = "setting.yaml"
	c := &Config{}
	yamlConfig, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("配置文件读取错误: %s\n", ConfigFile))
	}
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		panic(fmt.Errorf("配置初始化失败, %s\n", err))
	}
	ConfigInfo = c
}

func GetConfig() *Config {
	return ConfigInfo
}
