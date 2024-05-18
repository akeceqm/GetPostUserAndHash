package handle

import (
	"goserver/database"
	"goserver/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func SetDB(database *sqlx.DB) {
	db = database
}

func HandleUsersGET(c *gin.Context) {
	var users []database.User

	users = make([]database.User, 0)

	err := db.Select(&users, `SELECT * FROM  userinfo`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func HandleUserPOST(c *gin.Context) {
	var user database.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(`INSERT INTO userinfo (login, password) VALUES($1,$2) RETURNING id`, user.Login, middlewares.PasswordHash(user.Password))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}
