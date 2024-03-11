package queries

import (
	"database/sql"
	"fmt"
	"log"
)

type UserInterface struct {
  Id string
  Username string
  FirstName string
  LastName string
  Email string
}

func GetUser(db *sql.DB, user map[string]string) []UserInterface {
  var (
    id string
    username string
    firstName string
    lastName string
    email string
    password string
  )

	var result []UserInterface

  query := fmt.Sprintf("select * from users where id = '%v'", user["id"])
  err := db.QueryRow(query).Scan(&id, &username, &firstName, &lastName, &email, &password)
  if err != nil {
    log.Println(err)
  }

  result = append(result, UserInterface{id, username, firstName, lastName, email})
  
	return result
}

func UpdateUser(db *sql.DB, user map[string]string) map[string]any {
	result := make(map[string]any)

	query := fmt.Sprintf("update users set username = '%v', firstname = '%v', lastname = '%v' where id = '%v'", user["username"], user["firstName"], user["lastName"], user["id"])
	
	_, err := db.Exec(query)
    if err != nil {
      result["status"] = false
      result["message"] = err
    } else {
      result["status"] = true
      result["message"] = "success"
		}


	return result
}

func DeleteUser(db *sql.DB, user map[string]string) map[string]any {
	result := make(map[string]any)

	query := fmt.Sprintf("delete from users where id = '%v'", user["id"])
	
	_, err := db.Exec(query)
  if err != nil {
    result["status"] = false
    result["message"] = err
  } else { 
    result["status"] = true
    result["message"] = "success"
  }

	return result

}
