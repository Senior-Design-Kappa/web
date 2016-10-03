package auth

import (
  "database/sql"

  "gopkg.in/authboss.v0"
)

type DbUserStorer struct {
  db *sql.DB
}

type User struct {
  ID int
  Name string

  // Auth
  Email string
  Password string
}

func NewDbUserStorer(db *sql.Db) *DbStorer {
  return &DbUserStorer{
  }
}

func (d DbUserStorer) Put(key string, attr authboss.Attributes) error {
  var user User
  if err := attr.Bind(&user, true); err != nil {
    return err
  }
  return nil
}

func (d DbUserStorer) Get(key string) (result interface{}, err error) {
  return nil, nil
}
