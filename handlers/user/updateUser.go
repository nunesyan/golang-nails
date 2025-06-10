package user

import (
	"api-go/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "dados inválidos",
		})
		return
	}

	query := `
		UPDATE users
		SET name = $1, email = $2, password = $3
		WHERE id = $4
		RETURNING id, name, email, password
	`

	var update models.User
	err := db.QueryRow(query, input.Name, input.Email, input.Password, id).Scan(&id, &update.Name, &update.Email, &update.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao atualizar usuário",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"name":     update.Name,
		"email":    update.Email,
		"password": update.Password,
	})
}
