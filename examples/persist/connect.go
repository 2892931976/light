package persist

import (
	"database/sql"

	// import sql driver
	"github.com/gotips/log"
	_ "github.com/lib/pq"
)

const url = "postgres://postgres:@127.0.0.1:5432/meidb?sslmode=disable"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	// connect()
}

func connect() (err error) {
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Infof("successfully connect to %s", url)
	return nil
}
