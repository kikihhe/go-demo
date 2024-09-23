package configs

type Config struct {
	Postgresql *Postgresql `yaml:"postgresql"`
	Server     *Server     `yaml:"server"`
}
