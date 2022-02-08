package services

import (
	"last/dto"
	"last/models"
	"last/repository"
	"log"

	// "github.com/go-playground/validator/v10"
	"github.com/mashingan/smapping"
)

// type DeleteRetrun func(int)
type AuthorService interface {
	CreateAuthor(a dto.AuthorDTO) models.Author
	GetAllAuthor() []models.Author
	UpdateAuthor(id uint64, author dto.AuthorDTO) models.Author
	DeleteAuthor(id uint64) models.Author
	GetAuthorById(id uint64) models.Author
}

type authorService struct {
	authorRepository repository.AuthorRepository
}

// var validate = validator.New()

//ne
func NewAuthorService(authorRepo repository.AuthorRepository) AuthorService {

	return &authorService{
		authorRepository: authorRepo,
	}
}

func (service *authorService) CreateAuthor(a dto.AuthorDTO) models.Author {
	author := models.Author{}
	err := smapping.FillStruct(&author, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed to map %v:", err)
	}
	res := service.authorRepository.CreateAuthor(author)
	return res
}

func (service *authorService) GetAllAuthor() []models.Author {
	return service.authorRepository.GetAllAuthor()

}
func (service *authorService) GetAuthorById(id uint64) models.Author {
	return service.authorRepository.GetAuthorById(id)

}

func (service *authorService) DeleteAuthor(id uint64) models.Author {
	return service.authorRepository.DeleteAuthor(id)
}

func (service *authorService) UpdateAuthor(id uint64, a dto.AuthorDTO) models.Author {
	author := models.Author{}
	err := smapping.FillStruct(&author, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed to map %v:", err)
	}
	res := service.authorRepository.UpdateAuthor(id, author)
	return res
}
