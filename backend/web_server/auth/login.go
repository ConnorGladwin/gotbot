package auth

import (
	"database/sql"
	"fmt"

)

func Login(db *sql.DB, user map[string]string) map[string]string {
  loginResponse := make(map[string]string)

  userPassword := GetPassword(db, user["id"])

  if userPassword["error"] == "" {
    inputPassword := HashPassword(user["password"])   
    if userPassword["password"] == string(inputPassword) {
      loginResponse["status"] = "success"
			loginResponse["id"] = userPassword["id"]
    } else {
      loginResponse["status"] = "error"
      loginResponse["message"] = "password incorrect"
    }
  } else {
    loginResponse["status"] = "error"
    loginResponse["message"] = userPassword["error"]
  }

  return loginResponse
}

func GetPassword(db *sql.DB, id string) map[string]string {
  response := make(map[string]string)
  var userPassword string 
	var userId string

  usernameQuery := fmt.Sprintf("select id, password from users where username = '%v'", id)
  emailQuery := fmt.Sprintf("select id, password from users where email = '%v'", id)

  usernameErr := db.QueryRow(usernameQuery).Scan(&userId, &userPassword)
  emailErr := db.QueryRow(emailQuery).Scan(&userPassword)
  if usernameErr != nil && emailErr != nil {
    response["error"] = "username/email does not exist"
  } else {
    response["password"] = userPassword
		response["id"] = userId
  }

  return response
}
