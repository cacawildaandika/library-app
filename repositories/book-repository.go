package repositories

import (
	"github.com/cacawildaandika/go-gin-gorm-example/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookrepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (bookRepository *BookRepository) Save(book *models.Book) (error, *models.Book) {
	err := bookRepository.db.Save(book).Error

	if err != nil {
		return err, nil
	}

	return nil, book
}

func (bookRepository *BookRepository) FindAll() *models.Books {
	var books models.Books
	bookRepository.db.Find(&books)

	return &books
}

func (bookRepository *BookRepository) FindOne(id uint) (error, *models.Book) {
	var book models.Book

	err := bookRepository.db.Take(&book, id).Error

	if err != nil {
		return err, nil
	}

	return nil, &book
}

func (bookRepository *BookRepository) Delete(book *models.Book) error {
	err := bookRepository.db.Delete(&book).Error

	if err != nil {
		return err
	}

	return nil
}
