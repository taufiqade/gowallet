package middleware

import (
	"fmt"
	"strconv"
	"test/utils/helper"

	"github.com/gin-gonic/gin"
)

func (m *DefaultMiddleware) JWTAuthMidlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("admin route")
		authHeader := c.Request.Header["Authorization"]
		if len(authHeader) == 0 {
			c.AbortWithStatus(403)
		}
		token := authHeader[0]
		claim, status, err := helper.VerifyToken(token)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(401)
		}
		if status == 1 {
			c.AbortWithStatus(401)
		} else if status == 2 {
			c.AbortWithStatus(400)
		} else if claim.UserType != "admin" {
			c.AbortWithStatus(401)
		} else {
			c.Request.Header.Set("Wallet-Uid", strconv.Itoa(claim.UserID))
			c.Request.Header.Set("Wallet-Utype", claim.UserType)
			c.Next()
		}
	}
}

func (m *DefaultMiddleware) JWTAuthMidlewareGuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("guest route")
		authHeader := c.Request.Header["Authorization"]
		if len(authHeader) == 0 {
			c.AbortWithStatus(403)
			return
		}
		token := authHeader[0]
		claim, status, err := helper.VerifyToken(token)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(401)
			return
		}
		if status == 1 {
			c.AbortWithStatus(401)
			return
		} else if status == 2 {
			c.AbortWithStatus(400)
			return
		} else {
			c.Request.Header.Set("Wallet-Uid", strconv.Itoa(claim.UserID))
			c.Request.Header.Set("Wallet-Utype", claim.UserType)
			c.Next()
		}
	}
}