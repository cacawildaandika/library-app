package handlers

import (
	"net/http"

	"github.com/cacawildaandika/go-gin-gorm-example/services"
	"gorm.io/gorm"

	"github.com/cacawildaandika/go-gin-gorm-example/repositories"

	"github.com/cacawildaandika/go-gin-gorm-example/models"

	"github.com/cacawildaandika/go-gin-gorm-example/dtos"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	db, ok := c.MustGet("databaseConnection").(*gorm.DB)
	if !ok {
		panic("Can't connect to database")
	}

	bookRepository := repositories.NewBookrepository(db)

	response := services.GetAllBooks(*bookRepository)

	c.JSON(http.StatusOK, response)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		response := dtos.Response{
			Status:  "Invalid Data",
			Error:   err.Error(),
			Message: "Please check request",
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	db, ok := c.MustGet("databaseConnection").(*gorm.DB)
	if !ok {
		panic("Can't connect to database")
	}

	bookRepository := repositories.NewBookrepository(db)

	response := services.CreateBook(&book, *bookRepository)

	if response.Error != nil {
		c.JSON(http.StatusInternalServerError, response)
	}

	c.JSON(http.StatusOK, response)
}
