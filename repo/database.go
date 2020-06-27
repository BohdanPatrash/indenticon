package repo

import "github.com/BohdanPatrash/indenticon/dto"

const databasePath = "tmp/emails"

type DataBase struct {
	Users map[string]dto.User
}

var db *DataBase

//DB returns current database
func DB() *DataBase {
	return db
}

//Init initializes database
func Init() *DataBase {
	db = &DataBase{
		Users: make(map[string]dto.User, 0),
	}
	return db
}
