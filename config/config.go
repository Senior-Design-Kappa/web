package config

import "os"

type Config struct {
	Addr string
}

func NewConfig() Config {
	c := Config{
		Addr: os.Getenv("SERVER_ADDRESS") + ":8000",
	}
	return c
}
