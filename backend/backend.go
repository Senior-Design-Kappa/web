package backend

import (
	"database/sql"
	"log"

	"github.com/Senior-Design-Kappa/web/config"
	_ "github.com/mattn/go-sqlite3"
)

type Backend interface {
  CreateRoom(ownerId int64, videoId string) (int64, error)
  GetVideoId(roomId int64) (string, error)
}

type backend struct {
	Config config.Config
	db     *sql.DB
}

func NewBackend(conf config.Config) (Backend, error) {
	db, err := sql.Open("sqlite3", "./db/data.db")
	if err != nil {
		log.Printf("error: could not open db connection (%+v)\n", err)
	}
	b := &backend{
		Config: conf,
		db:     db,
	}
	return b, nil
}

func (b backend) CreateRoom(ownerId int64, videoId string) (int64, error) {
  stmt, err := b.db.Prepare("INSERT INTO rooms (owner_id, video_link) VALUES (?, ?)")
  if err != nil {
    return -1, err
  }
  res, err := stmt.Exec(ownerId, videoId)
  if (err != nil) {
    return -1, err
  }
  return res.LastInsertId()
}

func (b backend) GetVideoId(roomId int64) (string, error) {
  // Sad Machine
  defaultVideoId := "HAIDqt2aUek"

  row := b.db.QueryRow("SELECT video_link FROM rooms WHERE id=?", roomId)
  var videoId string
  if err := row.Scan(&videoId); err != nil {
    if err == sql.ErrNoRows {
      return defaultVideoId, nil
    } else {
      log.Printf("%s\n", err)
      return "", err
    }
  }
  return videoId, nil
}
