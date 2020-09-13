package routers

import (
	"github.com/cacawildaandika/go-gin-gorm-example/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeBookRouter(rg *gin.RouterGroup) {
	bookRouter := rg.Group("/book")

	bookRouter.GET("/", handlers.GetAllBooks)
	bookRouter.GET("/:id", handlers.GetBookById)
	bookRouter.POST("/", handlers.CreateBook)
}
