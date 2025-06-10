package routers

import (
	user "api-go/handlers/user"
	"api-go/handlers/work"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Routers(db *sql.DB) *gin.Engine {
	r := gin.Default()

	//users
	r.GET("/users", func(c *gin.Context) {
		user.GetUsers(c, db)
	})
	r.GET("/user/:id", func(c *gin.Context) {
		user.GetUserById(c, db)
	})
	r.POST("/create-user", func(c *gin.Context) {
		user.CreateUser(c, db)
	})
	r.PUT("/update-user/:id", func(c *gin.Context) {
		user.UpdateUser(c, db)
	})
	r.DELETE("/delete-user/:id", func(c *gin.Context) {
		user.DeleteUser(c, db)
	})

	//works
	r.GET("/works", func(c *gin.Context) {
		work.GetWorks(c, db)
	})
	r.GET("/work/:id", func(c *gin.Context) {
		work.GetWorkById(c, db)
	})
	r.POST("/create-work", func(c *gin.Context) {
		work.CreateWork(c, db)
	})
	r.PUT("update-work/:id", func(c *gin.Context) {
		work.UpdateWork(c, db)
	})
	r.DELETE("delete-work/:id", func(c *gin.Context) {
		work.DeleteWork(c, db)
	})

	return r
}
