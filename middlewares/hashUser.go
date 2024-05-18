package middlewares

import (
	"goserver/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var db *sqlx.DB

func SetDB(database *sqlx.DB) {
	db = database
}

func PasswordHash(password string) []byte {
	passwordHash := []byte(password)
	cost := 12
	hash, _ := bcrypt.GenerateFromPassword(passwordHash, cost)
	return hash
}

func AuthorizationAcc(c *gin.Context) {
	var user database.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Login == "" || user.Password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error: login or password == nil"})
		log.Println("Error: login or password == nil")
		return
	}
	var dbUser database.User
	err := db.Get(&dbUser, "SELECT login,password FROM userinfo WHERE login=$1", user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User authorized successfully"})
}
