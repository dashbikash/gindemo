package main

import (
	"main/handler"
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
  // Simple group: v1
	crud := router.Group("/crud")
	{
		crud.GET("/", handler.Get)
		crud.POST("/", handler.Post)
		crud.PUT("/", handler.Put)
    crud.DELETE("/", handler.Delete)
	}

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
