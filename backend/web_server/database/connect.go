package database

import (
	"database/sql"
	_ "github.com/lib/pq"

)

func Connect() (*sql.DB) {
  var connectionString string = "postgres://admin:vTuK3VqqzEroDwumrY91x9HLWXb2h4LN@dpg-cnm6h5mv3ddc73fkhung-a.frankfurt-postgres.render.com/gotbot"

  db, err := sql.Open("postgres", connectionString)
  if err != nil {
    panic(err)
  }

  return db
}

func Connected(db *sql.DB) string {
  var res = ""

  if err := db.Ping(); err != nil {
    res = "An error occured"
  } else {
    res = "connected"
  }

  return res
}
