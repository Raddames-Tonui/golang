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
  
  router.Run() 
}
