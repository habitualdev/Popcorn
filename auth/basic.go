package auth

import (
	"errors"
	"github.com/dgraph-io/badger/v3"
	"log"
)

type User struct {
	Username string
	Password 	 string
}

func CheckValue(user User) error {
	var value string

	db, err := badger.Open(badger.DefaultOptions("auth.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(user.Username))
		if err != nil {
			return err
		}
		item.Value(func(val []byte) error {
			value = string(val)
			return nil
		})
		return nil
	}); err != nil {
		return err
	}

	if value != user.Password {
		return errors.New("Invalid username or password")
	}

	return nil
}

func AddUser(user User) error {
	db, err := badger.Open(badger.DefaultOptions("auth.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(user.Username), []byte(user.Password))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}