package user

import (
	models "api-go/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT name, email, password FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Name, &user.Email, &user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
