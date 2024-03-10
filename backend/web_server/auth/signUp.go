package auth

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/sha3"
)

func SignUp(db *sql.DB, user map[string]string) string {
    message := ""

    user["id"] = CreateUserId()
    user["hashedPassword"] = HashPassword(user["password"])

    query := fmt.Sprintf("insert into users (id, firstname, lastname, username, email, password) values ('%v', '%v', '%v', '%v', '%v', '%v')", user["id"], user["firstName"], user["lastName"], user["username"], user["email"], user["hashedPassword"])

    _, err := db.Exec(query)
    if err != nil {
      insertError := fmt.Sprintf("%q", err)
      if strings.Contains(insertError, "users_username_key") {
        message = "username exists"
      } else if strings.Contains(insertError, "users_email_key") {
        message = "email exists"
      }
    } else {
      message = "success"
    }

    return message
 
}

func CreateUserId() string {
  id := uuid.New().String()
  return id
}

func HashPassword(password string) string {
  hash := sha3.New256()

  _, _ = hash.Write([]byte(password))

  sha3 := hash.Sum(nil)

  return fmt.Sprintf("%x", sha3)
}
