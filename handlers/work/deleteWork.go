package work

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteWork(c *gin.Context, db *sql.DB) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	result, err := db.Exec("DELETE FROM services WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao deletar serviço",
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao verificar serviço",
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "serviço não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "serviço deletado com sucesso!",
	})

}
