package main

import (
	"gin/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
  router := gin.Default()
  
  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  
  router.Run() // listen and serve on 0.0.0.0:8000 ✅ (Corrected)
}
