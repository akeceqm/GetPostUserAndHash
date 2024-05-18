package main

import (
	handle "goserver/Handle"
	"goserver/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB
var server *gin.Engine
var err error

const connectionString = "host=127.0.0.1 port=5432 user=postgres password=akeceqm dbname=users sslmode=disable"

func main() {

	server = gin.Default()
	middlewares.LogFile(server)
	db, err = sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("failed connection db")
	}
	defer db.Close()

	handle.SetDB(db)
	middlewares.SetDB(db)

	server.GET("/users", handle.HandleUsersGET)
	server.POST("/users", handle.HandleUserPOST)
	server.POST("/authorization", middlewares.AuthorizationAcc)

	server.Run(":8080")

}
