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
