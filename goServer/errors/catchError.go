package errorHandle

import (
	"database/sql"
	"log"
)

func CheckSQLErr(err error) {
	if err != nil && err != sql.ErrNoRows {
		log.Fatalln(err)
	}
}

func ThrowFatal(err error) { log.Fatalln(err) }
