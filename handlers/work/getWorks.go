package work

import (
	"api-go/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWorks(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT name, description, duration_minutes FROM services")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var works []models.Work
	for rows.Next() {
		var work models.Work
		if err := rows.Scan(&work.Name, &work.Description, &work.Duration_Minutes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		works = append(works, work)
	}

	c.JSON(http.StatusOK, works)
}
