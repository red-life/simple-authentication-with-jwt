package router

import (
	"github.com/gin-gonic/gin"
	"github.com/red-life/simple-authentication-with-jwt/internal/config"
	controller "github.com/red-life/simple-authentication-with-jwt/internal/controllers"
	"github.com/red-life/simple-authentication-with-jwt/internal/repository"
	service "github.com/red-life/simple-authentication-with-jwt/internal/services"
	"github.com/red-life/simple-authentication-with-jwt/pkg/jwt"
	"gorm.io/gorm"
)

func InitRouter() *gin.Engine{

	var postgresConfig *config.PostgresConfig
	err := config.ReadYAML("config.yaml", postgresConfig)
	if err != nil{
		panic(err)
	}
	var(
		jwtAuth jwt.IJWTPackage = jwt.NewJWTPackage("SECRET_KEY")
		postgresConn *gorm.DB = repository.PostgresConnection(postgresConfig)
		userRepository repository.IUserRepository = repository.NewUserRepository(postgresConn)
		userService service.IUserService = service.NewUserService(userRepository)
		userController controller.IUserController = controller.NewUserController(userService,jwtAuth)
	)

	r := gin.New()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/register",userController.AddUser)
		auth.POST("/login",userController.Login)
	}
	/*
	user := v1.Group("/user", middleware.isAuthorized())
	{
		user.Post("/me", handler)
	}
	*/
	return r

}