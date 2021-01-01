package user

import (
	"database/sql"
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/zondaTW/go-todolist-server/lib"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GetJWTAuthFunc(db *sql.DB) lib.AuthFunc {
	return func(context *gin.Context) (interface{}, error) {
		var login Login
		if err := context.ShouldBind(&login); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		if _, err := getUser(db, login.Username, login.Password); err == nil {
			return nil, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}

func AddUser(db *sql.DB) func(c *gin.Context) {
	return func(context *gin.Context) {
		var login Login
		var user User
		if err := context.ShouldBind(&login); err == nil {
			user, err = addUser(db, login.Username, login.Password)
			if err != nil {
				panic(err)
			}
			context.JSON(http.StatusOK, user)
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
