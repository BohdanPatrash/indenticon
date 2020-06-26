package repos

import (
	"log"

	"github.com/dgraph-io/badger"
)

const databasePath = "tmp/emails"

var db *badger.DB

//DB returns current database
func DB() *badger.DB {
	return db
}

//Init initializes database
func Init() *badger.DB {
	opts := badger.DefaultOptions(databasePath)
	opts.Logger = nil

	var err error
	db, err = badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
