package controller

import (
	"last/dto"
	"last/helper"
	"last/models"
	"last/services"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

type AuthorController interface {
	CreateAuthor(context *gin.Context)
	GetAllAuthor(context *gin.Context)
	GetAuthorById(context *gin.Context)
	UpdateAuthor(context *gin.Context)
	DeleteAuthor(context *gin.Context)
}
type authorController struct {
	authorService services.AuthorService
}

func NewAuthorController(authorServ services.AuthorService) AuthorController {
	return &authorController{
		authorService: authorServ,
	}
}

func (c *authorController) CreateAuthor(context *gin.Context) {
	var authorCreate dto.AuthorDTO
	errDTO := context.ShouldBindJSON(&authorCreate)
	if errDTO != nil {
		res := helper.BuildErorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		result := c.authorService.CreateAuthor(authorCreate)
		res := helper.BuildResponse(true, "Author Created", result)
		context.JSON(http.StatusCreated, res)
	}

}

func (c *authorController) UpdateAuthor(context *gin.Context) {
	var authorUpdate dto.AuthorDTO
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErorResponse("Invalid ID", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	errDTO := context.ShouldBindJSON(&authorUpdate)
	if errDTO != nil {
		res := helper.BuildErorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		result := c.authorService.UpdateAuthor(id, authorUpdate)
		res := helper.BuildResponse(true, "Author Updated", result)
		context.JSON(http.StatusOK, res)
	}

}

func (c *authorController) GetAuthorById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErorResponse("Invalid ID", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	var author models.Author = c.authorService.GetAuthorById(id)
	if (author == models.Author{}) {
		res := helper.BuildErorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", author)
		context.JSON(http.StatusOK, res)
	}

}

func (c *authorController) GetAllAuthor(context *gin.Context) {
	var authors []models.Author = c.authorService.GetAllAuthor()
	context.JSON(http.StatusOK, authors)
}

func (c *authorController) DeleteAuthor(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErorResponse("Invalid ID", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	result := c.authorService.DeleteAuthor(id)
	context.JSON(http.StatusOK, result)
}
