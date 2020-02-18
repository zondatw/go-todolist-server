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
		host     = config.Section("Database").Key("host").String()
		port     = config.Section("Database").Key("port").Value()
		user     = config.Section("Database").Key("user").String()
		password = config.Section("Database").Key("password").String()
		dbname   = config.Section("Database").Key("dbname").String()
	)

	connectionString := fmt.Sprintf(
		"host=%s port=%s "+
			"user=%s password=%s "+
			"dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router = gin.Default()
	initRoute()
	router.Run(":5000")
}
