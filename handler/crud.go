package handler

import (
	"main/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn = "postgres://zvxxcmnu:etJV4fcWltYMkOvwqSi2_jwY3HY-4qev@arjuna.db.elephantsql.com/zvxxcmnu"
)

func Get(ctx *gin.Context) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	var empList []model.Employee
	if err != nil {
		panic(err)
	} else {
		if result := db.Find(&empList); result.Error != nil {
			panic(result.Error)
		}
	}

	ctx.JSON(200, empList)
}
func Post(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"name": "bikash"})
}
func Put(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"name": "bikash"})
}
func Delete(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"name": "bikash"})
}
