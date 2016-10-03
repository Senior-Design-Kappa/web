package router

import (
  "gopkg.in/authboss.v0"
)

type MemStorer struct {
  Users map[string]authboss.Attributes
}

func NewMemStorer() *MemStorer {
  return &MemStorer{
    Users: make(map[string]authboss.Attributes),
  }
}

func (s MemStorer) Put(key string, attr authboss.Attributes) error {
  s.Users[key] = attr
  return nil
}

func (s MemStorer) Get(key string) (result interface{}, err error) {
  result, ok := s.Users[key]
  if !ok {
    return nil, authboss.ErrUserNotFound
  }

  return &result, nil
}
