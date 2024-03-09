package auth

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/sha3"
)

func SignUp(db *sql.DB, user map[string]string) {
  id := CreateUserId()
  log.Println(id)
}

func EmailExsits(db *sql.DB, email string) {
  //
}

func UsernameExists(db *sql.DB, username string) {
  //
}

func CreateUserId() string {
  id := uuid.New().String()
  return id
}

func HashPassword(password string) []byte {
  hash := sha3.New256()

  _, _ = hash.Write([]byte(password))

  sha3 := hash.Sum(nil)

  return sha3
}
