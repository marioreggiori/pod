package store

import (
	"os"
	"path"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

//var dbOpenErr error
var isAvailable = false

func init() {
	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	db, err = bolt.Open(path.Join(dir, ".pod.db"), 0600, nil)
	if err != nil {
		return
	}

	tx, err := db.Begin(true)
	if err != nil {
		return
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("custom"))
	if err != nil {
		return
	}

	if err := tx.Commit(); err != nil {
		return
	}
	isAvailable = true
}

func Close() {
	db.Close()
}
