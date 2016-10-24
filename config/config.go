package config

import (
  "os"
)

type Config struct {
	Addr string
  SyncAddr string
  ClientPath string
}

func NewDefaultConfig() Config {
	c := Config{
		Addr: ":8000",
    SyncAddr: "localhost:8001",
    ClientPath: "./",
	}
	return c
}

func (c * Config) UpdateFromEnvironment() {
  if os.Getenv("CLIENT_PATH") != "" {
    c.ClientPath = os.Getenv("CLIENT_PATH")
  }
  if os.Getenv("SYNC_ADDR") != "" {
    c.SyncAddr = os.Getenv("SYNC_ADDR")
  }
  if os.Getenv("ADDR") != "" {
    c.Addr = os.Getenv("ADDR")
  }
}
