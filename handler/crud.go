package handler

import "github.com/gin-gonic/gin"

func Get(ctx *gin.Context){
  ctx.JSON(200,map[string]string{"name":"bikash"})
}
func Post(ctx *gin.Context){
  ctx.JSON(200,map[string]string{"name":"bikash"})
}
func Put(ctx *gin.Context){
  ctx.JSON(200,map[string]string{"name":"bikash"})
}
func Delete(ctx *gin.Context){
  ctx.JSON(200,map[string]string{"name":"bikash"})
}