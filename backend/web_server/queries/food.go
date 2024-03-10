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
    result = "update" 
  case "DELETE":
    result = "delete"
  }

  return result
}

func GetAll(db *sql.DB) []FoodInterface {
  var (
    id string 
    name string 
    desc string
    price int 
  )

  var items []FoodInterface

	rows, err := db.Query("select * from food")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
	  err := rows.Scan(&id, &name, &desc, &price)
	  if err != nil {
		  log.Fatal(err)
	  }
    items = append(items, FoodInterface{id, name, desc, price})
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
  )

  var item []FoodInterface
  
  query := fmt.Sprintf("select * from food where id = '%v'", id)
  err := db.QueryRow(query).Scan(&id, &name, &desc, &price)
  if err != nil {
    log.Println(err)
  }

  item = append(item, FoodInterface{id, name, desc, price})

  return item
}

func InsertNew(db *sql.DB, item map[string]any) map[string]any {
  result := make(map[string]any)
  item["id"] = uuid.New().String()

  query := fmt.Sprintf("insert into food (id, name, description, price) values ('%v', '%v', '%v', %v)", item["id"], item["name"], item["desc"], item["price"])

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
