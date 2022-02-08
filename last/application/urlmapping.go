package application

import (
	"last/controller"
	"last/database"
	"last/middleware"
	"last/repository"
	"last/services"

	"gorm.io/gorm"
)

var (
	db               *gorm.DB                    = database.SetupDataBaseConnection()
	authorRepository repository.AuthorRepository = repository.NewAuthorRepository(db)
	userRepository   repository.UserRepository   = repository.NewUserRepository(db)
	jwtService       services.JWTService         = services.NewJWTService()
	authService      services.AuthService        = services.NewAuthService(userRepository)
	authorService    services.AuthorService      = services.NewAuthorService(authorRepository)
	authorController controller.AuthorController = controller.NewAuthorController(authorService)
	authController   controller.AuthController   = controller.NewAuthController(authService, jwtService)
)

func mapUrls() {

	authorRoute := r.Group("ap/author")
	authorRoute.Use(middleware.CustomErrors)
	{
		authorRoute.POST("/cauthor", authorController.CreateAuthor)
		authorRoute.GET("/getauthor", authorController.GetAllAuthor)
		authorRoute.PUT("/upauthor/:id", authorController.UpdateAuthor)
		authorRoute.DELETE("/delauthor/:id", authorController.DeleteAuthor)
		authorRoute.GET("/geteauthor/:id", authorController.GetAuthorById)
	}
	authRoutes := r.Group("api/auth")

	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
}
