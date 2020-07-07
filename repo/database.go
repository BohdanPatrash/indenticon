package repo

import (
	"github.com/BohdanPatrash/indenticon/dto"
	"github.com/BohdanPatrash/indenticon/service"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

const databasePath = "tmp/emails"

type DataBase struct {
	Users map[string]dto.User
}

var db *pg.DB

//DB returns current database
func DB() *pg.DB {
	return db
}

//Init initializes database
func Init() *pg.DB {
	if service.Testing() {
		db = pg.Connect(&pg.Options{
			Addr:     ":5432",
			User:     "postgres",
			Password: "password",
			Database: "test_indenticon",
		})
		return db
	}

	db = pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "password",
		Database: "indenticon",
	})
	return db
}

func DatabaseSetup() {
	models := []interface{}{
		(*dto.User)(nil),
	}
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
			Temp:        service.Testing(),
		})
		if err != nil {
			panic(err)
		}
	}
}
