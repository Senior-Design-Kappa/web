package config

type Config struct {
	Addr string
}

func NewConfig() Config {
	c := Config{
		Addr: ":8000",
	}
	return c
}
