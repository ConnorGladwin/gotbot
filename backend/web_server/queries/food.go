package queries

import (
	"database/sql"
	"log"
)

type FoodInterface struct {
  Id string
  Name string
  Desc string
  Price string
}

func FoodQueries(db *sql.DB, query map[string]string) any {
  var result any
  log.Println(query)

  switch query["type"] {
  case "get":
    result = "get"
  case "update":
    result = "update" 
  case "delete":
    result = "delete"
  default:
    result = GetAll(db)
  }

  return result
}

func GetAll(db *sql.DB) map[int]FoodInterface {
  var (
    id string 
    name string 
    desc string
    price string
  )

  items :=  make(map[int]FoodInterface)

	rows, err := db.Query("select * from food")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    i := 0
	  err := rows.Scan(&id, &name, &desc, &price)
	  if err != nil {
		  log.Fatal(err)
	  }
	  log.Println(id, name, desc, price)
    items[i] = FoodInterface{id, name, desc, price}
  }
  err = rows.Err()

  if err != nil {
	  log.Fatal(err)
  }

  return items
}
