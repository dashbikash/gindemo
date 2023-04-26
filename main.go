package main

import (
	"fmt"
	"log"
	"main/appdb"
	"main/handler"
	"main/middleware"
	"main/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	APP_HOST = ""
	APP_PORT = "8080"
)

func main() {
	go appdb.InitDB() // Initialize DB Schemas
	go initConfig()   // Initialize Config

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go and Gin!")
	})
	// Simple group: v1
	crud := router.Group("/crud")
	{
		crud.GET("/", handler.UserGet)
		crud.POST("/", handler.UserPost)
		crud.PUT("/:id", handler.UserPut)
		crud.DELETE("/:id", handler.UserDelete)
	}

	account := router.Group("/account")
	{
		account.POST("/login", handler.AccountLogin)
	}

	router.GET("/secured-token", middleware.AuthToken(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "This is token secured page"})
	})
	router.GET("/secured-jwt", middleware.AuthJWT(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "This is JWT secured page"})
	})

	router.GET("/jwt-new-token", func(ctx *gin.Context) {
		// Create a new token object, specifying signing method and the claims
		// you would like it to contain.
		claims := model.JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
				Issuer:    "Bikash",
			},
			JwtData: map[string]string{"User": "Bikash", "Role": "user"},
		}

		// Sign and get the complete encoded token as a string using the secret
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token_str, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		fmt.Println(token_str, err)
		ctx.JSON(http.StatusOK, map[string]string{"token": token_str})
	})
	router.GET("/jwt-validate", func(ctx *gin.Context) {
		token_str := ctx.Query("token")
		token, err := jwt.ParseWithClaims(token_str, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.JwtData, claims.RegisteredClaims.Issuer)
			ctx.JSON(http.StatusOK, claims.JwtData)
		} else {
			fmt.Println(err)
		}

		fmt.Println(token)
	})

	router.Run(APP_HOST + ":" + APP_PORT)
}
func initConfig() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}
