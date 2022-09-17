package main

import (
  "os"
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "hello": "nodion",
    })
  })

  if os.Getenv("PORT") != "" {
    r.Run("0.0.0.0:" + os.Getenv("PORT"))
  } else {
    r.Run("0.0.0.0:8080")
  }
}
