package services

import (
	"github.com/cacawildaandika/go-gin-gorm-example/dtos"
	"github.com/cacawildaandika/go-gin-gorm-example/models"
	"github.com/cacawildaandika/go-gin-gorm-example/repositories"
)

func CreateBook(book *models.Book, bookRepository repositories.BookRepository) dtos.Response {
	err, data := bookRepository.Save(book)

	if err != nil {
		return dtos.Response{
			Status: "Error",
			Error:  err.Error(),
		}
	}

	return dtos.Response{
		Status:  "Ok",
		Data:    data,
		Message: "Success create a book",
	}
}

func GetAllBooks(bookRepository repositories.BookRepository) dtos.Response {
	datas := bookRepository.FindAll()

	return dtos.Response{
		Status:  "Ok",
		Data:    datas,
		Message: "Success get all books",
	}
}

func GetById(id uint, bookRepository repositories.BookRepository) dtos.Response {
	err, result := bookRepository.FindOne(id)

	if err != nil {
		return dtos.Response{
			Status:  "Error",
			Error:   err.Error(),
			Message: "Failed get book",
		}
	}

	return dtos.Response{
		Status:  "Ok",
		Data:    result,
		Message: "Success get book",
	}
}
