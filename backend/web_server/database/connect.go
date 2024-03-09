package database

import (
	"database/sql"
	"os"
  "log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB) {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  dbConnString := os.Getenv("DB_CONN_STRING")

  var connectionString string = dbConnString
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
