package middleware

import (
	"fmt"
	"gin-demo/common"
	"gin-demo/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JwtAuthMiddleware 登录验证中间件
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			response.Response(c, http.StatusUnauthorized, 401, "请求的auth为空", gin.H{})
			c.Abort()
			return
		}

		reqToken, err := common.ParseToken(authHeader)
		if err != nil {
			response.Response(c, http.StatusUnauthorized, 401, "token验证失败", gin.H{})
			c.Abort()
			return
		}
		c.Set("username", reqToken.Username)
		c.Set("userId", reqToken.UserId)
		c.Next()

	}
}
