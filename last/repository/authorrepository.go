package repository

import (
	"last/models"
	"log"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	CreateAuthor(author models.Author) models.Author
	UpdateAuthor(id uint64, author models.Author) models.Author
	GetAllAuthor() []models.Author
	DeleteAuthor(id uint64) models.Author
	GetAuthorById(id uint64) models.Author
}

type authorConnection struct {
	connection *gorm.DB
}

//
func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorConnection{
		connection: db,
	}
}

func (db *authorConnection) CreateAuthor(author models.Author) models.Author {

	err := db.connection.Create(&author).Error
	if err != nil {
		log.Fatalf("Failed to create author %v", err)

	}
	return author
}

func (db *authorConnection) UpdateAuthor(id uint64, author models.Author) models.Author {
	// var authorUpdate models.Author
	db.connection.Updates(&author).Where("id = ?", id)
	return author
}

func (db *authorConnection) GetAllAuthor() []models.Author {
	var authors []models.Author
	db.connection.Find(&authors)
	return authors
}

func (db *authorConnection) GetAuthorById(id uint64) models.Author {
	var authors models.Author
	db.connection.Where("id = ?", id).First(&authors)
	return authors
}

func (db *authorConnection) DeleteAuthor(id uint64) models.Author {
	var author models.Author
	db.connection.Where("id = ?", id).Delete(&author)
	return author

}
