package handler

import (
	"crypto/sha256"
	"fmt"
	"main/appdb"
	"main/model"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UserGet(ctx *gin.Context) {

	db, err := appdb.GetConnection()
	var users []appdb.User
	if err == nil {
		db.Find(&users)
		ctx.JSON(200, users)
		return
	}

}

func UserPost(ctx *gin.Context) {
	var reqUser model.RequestUser
	if err := ctx.ShouldBind(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashing := sha256.New()
	hashing.Write([]byte(reqUser.Passwd))
	passHash := fmt.Sprintf("%x", hashing.Sum(nil))
	dbUser := appdb.User{LoginID: reqUser.LoginID, Name: reqUser.Name, Role: reqUser.Role, Country: reqUser.Country, Passwd: passHash}
	db, err := appdb.GetConnection()
	if err == nil {
		db.Create(&dbUser)
	}

	ctx.JSON(200, dbUser.UserID)
}
func UserPut(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"name": "bikash"})
}

func UserDelete(ctx *gin.Context) {
	id := (ctx.Param("id"))
	db, err := appdb.GetConnection()

	if err == nil {
		db.Delete(&appdb.User{}, id)
	}
	ctx.JSON(http.StatusOK, "Deleted")
}

func AccountLogin(ctx *gin.Context) {
	var login map[string]string

	if err := ctx.ShouldBind(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, err := appdb.GetConnection()
	var user appdb.User
	if err == nil {
		db.First(&user, "login_id = ?", login["loginid"])
		claims := model.JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
				Issuer:    "System",
			},
			JwtData: map[string]string{"ID": strconv.Itoa(user.UserID), "LoginID": user.LoginID, "Name": user.Name, "Role": user.Role},
		}

		// Sign and get the complete encoded token as a string using the secret
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token_str, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		ctx.JSON(200, map[string]string{"token": token_str})
		return
	}

}
