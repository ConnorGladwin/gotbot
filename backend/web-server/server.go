package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/", func(ctx *gin.Context) {
    ctx.String(http.StatusOK, "Hello, World")
  })

  r.GET("/a", handler)

  r.Run(":8080")
}

func handler(ctx *gin.Context) {
  ctx.String(http.StatusOK, "This is a route")
}
