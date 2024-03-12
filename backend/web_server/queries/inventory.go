package queries

import (
	"database/sql"
	"fmt"
)

func InventoryUpdate(db *sql.DB, id string, value int) map[string]any {
  result := make(map[string]any)

  query := fmt.Sprintf("update food set stock = %v where id = '%v'", value, id)
  _, err := db.Exec(query)
  if err != nil {
    result["status"] = "error"
    result["message"] = err
  } else {
    result["status"] = "success"
    result["message"] = "inventory successfully updated"
  }

  return result
}
