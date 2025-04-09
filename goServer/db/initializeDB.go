package db

import (
	"database/sql"
	"log"

	e "goServer/errors"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

//! this could be better in the future
// type ContactController struct {
// 	db  *db.Queries
// 	ctx context.Context
// }

// func NewContactController(db *db.Queries, ctx context.Context) *ContactController {
// 	return &ContactController{db, ctx}
// }

const (
	Path_To_DB string = "./db/tasks.db"
)

func InitializeDb() {
	db, err := sql.Open("sqlite3", Path_To_DB)

	if err != nil {
		log.Fatalln(err)
	}
	Db = db

}

func CreateTables() {

	if err := Db.Ping(); err != nil {
		e.ThrowFatal(err)
	}

	_, err := Db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
    	id integer primary key autoincrement unique not null,
    	name varchar(64) not null check(length(name) > 3),
    	description varchar(128) null,
    	status varchar(16) not null default('pending') CHECK (status IN ('done', 'pending', 'cancelled')),
    	photo text null,
    	hex_color char(7) not null default('#000000') CHECK(hex_color LIKE '#%')
	);
	`)

	if err != nil {
		log.Fatalln(err)
	}

}
