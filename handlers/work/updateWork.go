package work

import (
	"api-go/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateWork(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var input models.Work
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "dados inválidos",
		})
		return
	}

	query := `
		UPDATE services
		SET name = $1, description = $2, duration_minutes = $3
		WHERE id = $4
		RETURNING id, name, description, duration_minutes
	`

	var update models.Work
	err := db.QueryRow(query, input.Name, input.Description, input.Duration_Minutes, id).Scan(&id, &update.Name, &update.Description, &update.Duration_Minutes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao atualizar serviço",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               id,
		"name":             update.Name,
		"description":      update.Description,
		"duration_minutes": update.Duration_Minutes,
	})
}
