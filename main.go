package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	config *ini.File
	db     *sql.DB
	router *gin.Engine
)

func init() {
	godotenv.Load(".env")
}

func main() {
	var err error
	var (
		ip         = os.Getenv("SYSTEM_IP")
		port       = os.Getenv("SYSTEM_PORT")
		authKey    = os.Getenv("SYSTEM_AUTH_KEY")
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
	)

	connectionString := fmt.Sprintf(
		"host=%s port=%s "+
			"user=%s password=%s "+
			"dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	urlSetting := fmt.Sprintf("%s:%s", ip, port)
	router = gin.Default()
	router.GET("/", health)
	router.GET("/health", health)
	initRoute(authKey)
	router.Run(urlSetting)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
