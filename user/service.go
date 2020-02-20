package user

import (
	"database/sql"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/zondaTW/go-todolist-server/lib"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GetJWTAuthFunc(db *sql.DB) lib.AuthFunc {
	return func(c *gin.Context) (interface{}, error) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		if _, err := getUser(db, login.Username, login.Password); err == nil {
			return nil, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}
