package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	_ "github.com/lib/pq"
)

var (
	config *ini.File
	db     *sql.DB
	router *gin.Engine
)

func init() {
	var err error
	config, err = ini.Load("env.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}

func main() {
	var err error
	var (
		ip         = config.Section("System").Key("ip").String()
		port       = config.Section("System").Key("port").String()
		authKey    = config.Section("System").Key("auth_key").String()
		dbHost     = config.Section("Database").Key("host").String()
		dbPort     = config.Section("Database").Key("port").Value()
		dbUser     = config.Section("Database").Key("user").String()
		dbPassword = config.Section("Database").Key("password").String()
		dbName     = config.Section("Database").Key("dbname").String()
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
	initRoute(authKey)
	router.Run(urlSetting)
}
