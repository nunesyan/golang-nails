package user

import (
	models "api-go/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var user models.User
	err := db.QueryRow("SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "usuário não encontrado",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
