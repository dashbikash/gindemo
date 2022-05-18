package main

import (
  "os"
	"github.com/gin-gonic/gin"
)

var( 
  APP_HOST=""
  APP_PORT="8088"  
)


func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go and Gin!")
	})

  configServer()
	router.Run(APP_HOST+":"+APP_PORT)
}
func configServer(){
  if os.Getenv("APP_HOST")!=""{
    APP_HOST=os.Getenv("APP_HOST")
  }
  if os.Getenv("APP_PORT")!=""{
    APP_PORT=os.Getenv("APP_PORT")
  }
}
