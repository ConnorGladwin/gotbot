package queries

import (
	"database/sql"
	"log"
)

func GetUser(db *sql.DB, user map[string]string) map[string]string {
	result := make(map[string]string)
	log.Println(user["id"])
	return result
}

func UpdateUser(db *sql.DB, user map[string]string) map[string]string {
	result := make(map[string]string)
	log.Println(user)
	return result
}

func DeleteUser(db *sql.DB, user map[string]string) map[string]string {
	result := make(map[string]string)
	log.Println(user)
	return result
}
