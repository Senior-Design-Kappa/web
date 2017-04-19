package logic

import (
	"github.com/Senior-Design-Kappa/web/backend"
	"github.com/Senior-Design-Kappa/web/config"
)

type Logic interface {
  CreateRoom(ownerId int64, videoLink string) (int64, error)
  GetVideoId(roomId int64) (string, error)
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

func (l logic) CreateRoom(ownerId int64, videoLink string) (int64, error) {
  return l.backend.CreateRoom(ownerId, videoLink)
}

func (l logic) GetVideoId(roomId int64) (string, error) {
  return l.backend.GetVideoId(roomId)
}
