package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	helper "github.com/taufiqade/gowallet/utils/helper"
)

// JWTAuthMidlewareAdmin godoc
func (m *DefaultMiddleware) JWTAuthMidlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("admin route")
		authHeader := c.Request.Header["Authorization"]
		if len(authHeader) == 0 {
			c.AbortWithStatus(403)
		}
		token := strings.Split(authHeader[0], " ")[1]
		claim, status, err := helper.VerifyToken(token)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(401)
		}
		if status == 1 {
			c.AbortWithStatus(401)
		} else if status == 2 {
			c.AbortWithStatus(400)
		} else if claim.Type != "admin" {
			c.AbortWithStatus(401)
		}

		c.Next()
	}
}

// JWTAuthMidlewareGuest godoc
func (m *DefaultMiddleware) JWTAuthMidlewareGuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("guest route")
		authHeader := c.Request.Header["Authorization"]
		if len(authHeader) == 0 {
			c.AbortWithStatus(403)
			return
		}
		token := strings.Split(authHeader[0], " ")[1]
		_, status, err := helper.VerifyToken(token)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(401)
			return
		}
		if status == 1 {
			fmt.Println("status")
			c.AbortWithStatus(401)
			return
		} else if status == 2 {
			c.AbortWithStatus(400)
			return
		}

		c.Next()
	}
}
