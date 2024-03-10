package server

import (
	"database/sql"
	// "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	// "gotbot/backend/web_server/auth"
	"gotbot/backend/web_server/auth"
	"gotbot/backend/web_server/database"
	"gotbot/backend/web_server/queries"
)

type Data struct {
  Name string
}

func Start() {
  err := godotenv.Load()

  if err != nil {
    log.Fatal("Error loading .env file")
  }

  r := gin.Default() 

  var db = database.Connect()

  // auth handler
  authGroup := r.Group("/auth")
  authGroup.Use(AuthGroupHandler) 

  authGroup.POST("/sign-up", SignUp(db))
  authGroup.POST("/login", Login(db))

  r.GET("/ping", DbConnection(db))

  apiGroup := r.Group("/api")
  apiGroup.Use(VerifyToken)
  apiGroup.GET("/a", handler)

  // user query handler

  // food query handler
  apiGroup.GET("/food", FoodQuery(db))

  // inventory query handler

  r.Run(":1234")
}

func handler(ctx *gin.Context) {
  ctx.String(http.StatusOK, "handler")
}

func AuthGroupHandler(ctx *gin.Context) {
  request := ctx.Request.URL.Path
  
    switch request {
    case "/auth/sign-up":
      ctx.Next()
    case "/auth/login":
      ctx.Next()
    default:
      ctx.String(http.StatusNotFound, "Not Found")
    }
}

func DbConnection(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    conn := database.Connected(db)
    ctx.String(http.StatusOK, conn)
  }
}

var secretKey = []byte("secret-key")

// create JWT token
func CreateToken(username string) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": username,
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  })

  tokenString, err := token.SignedString(secretKey)
  if err != nil {
    log.Println(err)
    return "", err
  }

  return tokenString, nil
}

// verify JWT token 
func VerifyToken(ctx *gin.Context) {
  tokenString := ctx.GetHeader("Authorization")
  tokenString = tokenString[len("Bearer "):]
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
   })
  
   if err != nil {
      log.Println(err)
   }
  
   if !token.Valid {
     log.Println("invalid token")
     ctx.AbortWithStatus(401)
   } else {
     ctx.Next()
   }
}

// sign-up
func SignUp(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    user := make(map[string]string)

    user["username"], _ = ctx.GetPostForm("username")
    user["firstName"], _ = ctx.GetPostForm("firstName")
    user["lastName"], _ = ctx.GetPostForm("lastName")
    user["email"], _ = ctx.GetPostForm("email")
    user["password"], _ = ctx.GetPostForm("password")

    res := auth.SignUp(db, user)
    ctx.String(http.StatusOK, res)
  }
}

// login
func Login(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    user := make(map[string]string)

    user["id"], _ = ctx.GetPostForm("id")
    user["password"], _ = ctx.GetPostForm("password")

    res := auth.Login(db, user)
    res["token"], _ = CreateToken(user["id"])
    ctx.JSON(http.StatusOK, res)
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
