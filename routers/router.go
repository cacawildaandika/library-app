package routers

import (
	"github.com/cacawildaandika/go-gin-gorm-example/database"
	"github.com/cacawildaandika/go-gin-gorm-example/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConnection", db)
		c.Next()
	}
}

func Run() {
	r := gin.Default()

	db := database.Connect()

	db.AutoMigrate(&models.Book{})

	r.Use(DatabaseMiddleware(db))

	v1 := r.Group("v1")
	InitializeBookRouter(v1)

	r.Run(":5000")
}
