package server

import (
	"database/sql"
	"strconv"

	// "fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

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
  r.Use(cors.Default())

  var db = database.Connect()
  
  r.GET("/auth/validate", ValidateRequest)

  // auth handler
  authGroup := r.Group("/auth")
  authGroup.Use(AuthGroupHandler) 

  authGroup.POST("/sign-up", SignUp(db))
  authGroup.POST("/login", Login(db))

  r.GET("/ping", DbConnection(db))

  apiGroup := r.Group("/api")
  apiGroup.Use(VerifyToken)

  // user query handler
	apiGroup.GET("/user", UserQuery(db))
	apiGroup.PATCH("/user", UserQuery(db))
	apiGroup.DELETE("/user", UserQuery(db))

  // food query handler
  apiGroup.GET("/food", FoodQuery(db))
  apiGroup.POST("/food", FoodQuery(db))
  apiGroup.PATCH("/food", FoodQuery(db))
  apiGroup.DELETE("/food", FoodQuery(db))

  // inventory query handler
  apiGroup.GET("/inventory", InventoryQuery(db))

  r.Run(":8080")
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

// create HMAC secret instance
var secretKey = []byte("pleasehireme")

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

func ValidateRequest(ctx *gin.Context) {
  tokenString, _ := ctx.GetQuery("token")
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
   })
  
   if err != nil {
      log.Println(err)
   }
  
   if !token.Valid {
    log.Println("invalid token")
    ctx.AbortWithStatus(500)
   } else {
    log.Println("token valid")
    ctx.Status(http.StatusOK)
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
    res["token"], _ = CreateToken(res["id"])
    ctx.JSON(http.StatusOK, res)
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
    var res any
    requestMethod := ctx.Request.Method
    query := make(map[string]string)
		query["id"], _ = ctx.GetQuery("id")
		query["username"], _ = ctx.GetQuery("username")
    query["firstName"], _ = ctx.GetQuery("firstName")
    query["lastName"], _ = ctx.GetQuery("lastName")
		query["email"], _ = ctx.GetQuery("email")
		query["password"], _ = ctx.GetQuery("password")

		switch requestMethod {
			case "GET":
				res = queries.GetUser(db, query)
			case "PATCH":
				res = queries.UpdateUser(db, query)
			case "DELETE":
				res = queries.DeleteUser(db, query)
			default:
				res = http.StatusNotFound
		}

		ctx.JSON(http.StatusOK, res)
  }
}

// food query function
func FoodQuery(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    var res any
    requestMethod := ctx.Request.Method
    query := make(map[string]any)

    switch requestMethod {
    case "GET":
      request := ctx.Request.URL.Query()
      if request.Get("type") == "all" {
        query["type"] = "all"
      } else {
        query["type"] = request.Get("type")
        query["id"] = request.Get("id")
      }
      res = queries.FoodQueries(db, requestMethod, query)
    case "POST", "PATCH":
      query["id"], _ = ctx.GetPostForm("id")
			query["name"], _ = ctx.GetPostForm("name")
      query["desc"], _ = ctx.GetPostForm("desc")
      query["price"], _ = ctx.GetPostForm("price")
			query["stock"], _ = ctx.GetPostForm("stock")

      res = queries.FoodQueries(db, requestMethod, query)
    case "DELETE":
			query["id"], _ = ctx.GetPostForm("id")
			res = queries.FoodQueries(db, requestMethod, query)
    }

    ctx.JSON(http.StatusOK, res)
  }
}

// inventory
func InventoryQuery(db *sql.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    id, _ := ctx.GetQuery("id")
    requestValue, _ := ctx.GetQuery("value")
    value, _ := strconv.Atoi(requestValue)

    res := queries.InventoryUpdate(db, id, value)

    ctx.JSON(http.StatusOK, res)
 }
}
