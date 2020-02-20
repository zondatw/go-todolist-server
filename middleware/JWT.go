package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/zondaTW/go-todolist-server/lib"
)

func GetJWTMiddleware(key string, gwtAuthFunc lib.AuthFunc) *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(
		&jwt.GinJWTMiddleware{
			Realm:         "Login Required",
			Key:           []byte(key),
			Timeout:       time.Hour * 12,
			MaxRefresh:    time.Hour * 24,
			Authenticator: gwtAuthFunc,
			Unauthorized:  jwtUnAuthFunc,
		})

	if err != nil {
		panic(err)
	}

	return authMiddleware
}

func jwtUnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
