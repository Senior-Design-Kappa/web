package config

type Config struct {
	Addr string
}

func NewConfig() Config {
	c := Config{
		Addr: "localhost:8080",
	}
	return c
}
