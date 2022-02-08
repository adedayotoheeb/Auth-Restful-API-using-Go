package controller

import (
	"last/dto"
	"last/helper"
	"last/models"
	"last/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDTO
	errDTO := ctx.ShouldBindJSON(&loginDto)
	if errDTO != nil {
		response := helper.BuildErorResponse("Failed t process reuest", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(models.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "Login Successfully", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErorResponse("Please Check your Credentials", "Invalid Credential", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBindJSON(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErorResponse("Failed to process reuest", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErorResponse("Failed to process reuest","Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest,response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token =token
		response := helper.BuildResponse(true,"Successfully registered",createdUser)
		ctx.JSON(http.StatusCreated,response)
	}
}
