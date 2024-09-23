package configs

import "strconv"

type Server struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func GetAddress() string {
	ip := GetConfig().Server.IP
	port := GetConfig().Server.Port
	if ip == "" {
		ip = "localhost"
	}
	return ip + ":" + strconv.Itoa(port)
}
