package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器错误",
	})
}
func RespErrorWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"info": data,
	})
}
func RespSuccessful(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"info": "成功！",
	})
}
func RespSuccessfulWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"info": "成功",
		"data": data,
	})
}
