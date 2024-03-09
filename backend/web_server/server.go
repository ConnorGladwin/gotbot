package server

import (
	"database/sql"
	"strings"
	// "encoding/json"
	"log"
	"net/http"

	// "net/url"

	"github.com/gin-gonic/gin"

	// "gotbot/backend/web_server/auth"
	"gotbot/backend/web_server/database"
	"gotbot/backend/web_server/queries"
  // "gotbot/backend/web_server/auth"
)

type Data struct {
  Name string
}

func Start() {
  r := gin.Default() 

  var db = database.Connect()

  // auth handler
  r.Use(AuthHandler)

  r.GET("auth/sign-up", SignUp(db))

  r.GET("/ping", DbConnection(db))
  r.GET("/a", handler)

  // user query handler

  // food query handler
  r.GET("/food", FoodQuery(db))

  // inventory query handler

  r.Run(":1234")
}

func handler(ctx *gin.Context) {
  data := ctx.Request.URL.Query()
  log.Println(data.Get("value"))

  ctx.String(http.StatusOK, "handler")
}

func AuthHandler(ctx *gin.Context) {
  request := ctx.Request.URL.Path
  
  if !strings.Contains(request, "auth") {
    ctx.String(http.StatusUnauthorized, "Unauthorized")
    ctx.Abort()
  } else {
    switch request {
    case "/auth/sign-up":
      ctx.Next()
    case "/auth/login":
      ctx.String(http.StatusOK, "login")
    default:
      ctx.String(http.StatusNotFound, "Not Found")
    }
  }
}

func DbConnection(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    conn := database.Connected(db)
    ctx.String(http.StatusOK, conn)
  }
}

// sign-up
func SignUp(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    request := ctx.Request.Body
    log.Println(request)
  }
}

// login
func Login(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    //
  }
}

// user query function
func UserQuery(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    //
  }
}

// food query function
func FoodQuery(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    request := ctx.Request.URL.Query()
    query := make(map[string]string)

    if request.Get("type") == "all" {
      query["type"] = "all"
    } else {
      query["type"] = request.Get("type")
      query["value"] = request.Get("value")
    }

    res := queries.FoodQueries(db, query)
    ctx.JSON(http.StatusOK, res)
  }
}

// inventory
func InventoryQuery(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    //
  }
}
