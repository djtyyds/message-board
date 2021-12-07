package api

import (
	"github.com/gin-gonic/gin"
	"message-board/tool"
)

func auth(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		tool.RespErrorWithData(c, "请登录后再操作")
		c.Abort()
	}
	c.Set("username", username)
	c.Next()
}
