package work

import (
	"api-go/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWork(c *gin.Context, db *sql.DB) {
	var work models.Work

	if err := c.ShouldBindBodyWithJSON(&work); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "json inválido",
		})
		return
	}

	if work.Name == "" || work.Description == "" || work.Duration_Minutes == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "é necessário preencher todo json",
		})
		return
	}

	err := db.QueryRow(`
		INSERT INTO services (name, description, duration_minutes) VALUES ($1, $2, $3) RETURNING id
	`, work.Name, work.Description, work.Duration_Minutes).Scan(&work.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               work.ID,
		"name":             work.Name,
		"description":      work.Description,
		"duration_minutes": work.Duration_Minutes,
	})
}
