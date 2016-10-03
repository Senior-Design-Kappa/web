package main

import (
	"log"

	"github.com/Senior-Design-Kappa/web/auth"
	"github.com/Senior-Design-Kappa/web/backend"
	"github.com/Senior-Design-Kappa/web/config"
	"github.com/Senior-Design-Kappa/web/logic"
	"github.com/Senior-Design-Kappa/web/router"
)

func main() {
	conf := config.NewConfig()
	b := makeBackend(conf)
	l := makeLogic(conf, b)
  a := makeAuth()
	s := router.NewServer(conf, l, a)
	log.Fatal(s.ListenAndServe())
}

func makeAuth() auth.Auth {
  a, err := auth.NewAuth()
  if err != nil {
    log.Fatalf("error: auth layer could not be created (%+v)\n", err)
  }
  return a
}

func makeBackend(conf config.Config) backend.Backend {
	b, err := backend.NewBackend(conf)
	if err != nil {
		log.Fatalf("error: backend layer could not be created (%+v)\n", err)
	}
	return b
}

func makeLogic(conf config.Config, backend backend.Backend) logic.Logic {
	l, err := logic.NewLogic(conf, backend)
	if err != nil {
		log.Fatalf("error: logic layer could not be created (%+v)\n", err)
	}
	return l
}
