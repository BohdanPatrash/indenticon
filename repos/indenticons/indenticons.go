package indenticons

import (
	"crypto/sha256"

	"github.com/BohdanPatrash/indenticon/repos"
	"github.com/dgraph-io/badger"
)

func GetByEmail(email string) (string, error) {
	hash := sha256.Sum256([]byte(email))
	return GetByHash(hash[:])
}

func GetByHash(hash []byte) (string, error) {
	var icon string
	db := repos.DB()

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(hash)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			icon = string(val)
			return nil
		})
		return err
	})

	return icon, err
}
