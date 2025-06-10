package work

import (
	"api-go/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWorkById(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var work models.Work
	err := db.QueryRow("SELECT id, name, description FROM services WHERE id = $1", id).Scan(&work.ID, &work.Name, &work.Description)
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

	c.JSON(http.StatusOK, work)
}
