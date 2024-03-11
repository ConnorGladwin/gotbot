package queries

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type FoodInterface struct {
  Id string
  Name string
  Desc string
  Price int 
	Stock int
}

func FoodQueries(db *sql.DB, method string, query map[string]any) any {
  var result any

  switch method {
  case "GET":
    if query["type"] == "all" {
      result = GetAll(db)
    } else {
      result = GetById(db, query["id"].(string))
    }
  case "POST":
		result = InsertNew(db, query)
  case "PATCH":
    result = UpdateItem(db, query)
  case "DELETE":
    result = DeleteItem(db, query)
  }

  return result
}

func GetAll(db *sql.DB) []FoodInterface {
  var (
    id string 
    name string 
    desc string
    price int 
		stock int
  )

  var items []FoodInterface

	rows, err := db.Query("select * from food")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
	  err := rows.Scan(&id, &name, &desc, &price, &stock)
	  if err != nil {
		  log.Fatal(err)
	  }
    items = append(items, FoodInterface{id, name, desc, price, stock})
  }
  err = rows.Err()

  if err != nil {
	  log.Fatal(err)
  }

  return items
}

func GetById(db *sql.DB, id string) []FoodInterface {
  var (
    name string
    desc string
    price int
		stock int
  )

  var item []FoodInterface
  
  query := fmt.Sprintf("select * from food where id = '%v'", id)
  err := db.QueryRow(query).Scan(&id, &name, &desc, &price, &stock)
  if err != nil {
    log.Println(err)
  }

  item = append(item, FoodInterface{id, name, desc, price, stock})

  return item
}

func InsertNew(db *sql.DB, item map[string]any) map[string]any {
  result := make(map[string]any)
  item["id"] = uuid.New().String()

  query := fmt.Sprintf("insert into food (id, name, description, price, stock) values ('%v', '%v', '%v', %v, %v)", item["id"], item["name"], item["desc"], item["price"], item["stock"])
	log.Println(query)

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

func UpdateItem(db *sql.DB, item map[string]any) map[string]any {
	result := make(map[string]any)

	query := fmt.Sprintf("update food set name = '%v', description = '%v', price = %v where id = '%v'", item["name"], item["desc"], item["price"], item["id"])
	
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

func DeleteItem(db *sql.DB, item map[string]any) map[string]any {
	result := make(map[string]any)

	query := fmt.Sprintf("delete from food where id = '%v'", item["id"])
	
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
