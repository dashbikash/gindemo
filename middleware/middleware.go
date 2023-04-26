package middleware

import (
	"fmt"
	"main/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		fmt.Println(token)
		if strings.Compare(token, "mytoken") == 0 {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "Not Authorized"})
		}

	}
}

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt_token := ctx.GetHeader("Authorization")
		token, err := jwt.ParseWithClaims(jwt_token, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})

		if claims, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.JwtData, claims.RegisteredClaims.Issuer)
			ctx.Next()
		} else {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "Not Authorized"})
		}

	}
}
