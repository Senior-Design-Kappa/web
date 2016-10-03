package auth

import (
  "database/sql"
  "log"

  "gopkg.in/authboss.v0"

  _ "github.com/mattn/go-sqlite3"
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

func NewDbUserStorer() *DbUserStorer {
  db, err := sql.Open("sqlite3", "./db/users.db")
  if err != nil {
    log.Printf("error: could not open db connection (%+v)\n", err)
  }
  return &DbUserStorer{
    db: db,
  }
}

func (d DbUserStorer) Create(key string, attr authboss.Attributes) error {
  row := d.db.QueryRow("SELECT email FROM users WHERE email=?", key)
  if err := row.Scan(); err == sql.ErrNoRows {
    d.Put(key, attr)
    return nil
  } else {
    return authboss.ErrUserFound
  }
}

func (d DbUserStorer) Put(key string, attr authboss.Attributes) error {
  var user User
  if err := attr.Bind(&user, true); err != nil {
    return err
  }
  stmt, err := d.db.Prepare("INSERT INTO users (id, email, password, name) VALUES (?, ?, ?, ?)")
  if err != nil {
    return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(user.ID, user.Email, user.Password, user.Name)
  if err != nil {
    return err
  }
  return nil
}

func (d DbUserStorer) Get(key string) (result interface{}, err error) {
  row := d.db.QueryRow("SELECT id, email, password, name FROM users WHERE email=?", key)
  var user User
  if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name); err != nil {
    if err == sql.ErrNoRows {
      return nil, authboss.ErrUserNotFound
    } else {
      return nil, err
    }
  }
  return &user, nil
}
