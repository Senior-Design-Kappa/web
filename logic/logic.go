package logic

import (
	"github.com/Senior-Design-Kappa/web/backend"
	"github.com/Senior-Design-Kappa/web/config"
)

type Logic interface {
}
type logic struct {
	backend backend.Backend
	Config  config.Config
}

func NewLogic(conf config.Config, backend backend.Backend) (Logic, error) {
	l := &logic{
		backend: backend,
		Config:  conf,
	}
	return l, nil
}