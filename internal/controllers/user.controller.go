package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/red-life/simple-authentication-with-jwt/internal/dto"
	service "github.com/red-life/simple-authentication-with-jwt/internal/services"
	"github.com/red-life/simple-authentication-with-jwt/pkg/jwt"
	"net/http"
)

type IUserController interface {
	AddUser(c *gin.Context)
	Login(c *gin.Context)
	DeleteUser(C *gin.Context)
}

func NewUserController(userService service.IUserService, jwt jwt.IJWTPackage) IUserController{
	return &UserController{
		userService: userService,
		jwt: jwt,
	}
}

type UserController struct {
	userService service.IUserService
	jwt jwt.IJWTPackage
}


func (userController UserController)AddUser(c *gin.Context){
	var registerDto dto.RegisterDTO
	if err := c.ShouldBindJSON(&registerDto); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err := validate.Struct(registerDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userController.userService.AddUser(registerDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtToken, err := userController.jwt.CreateToken(registerDto.Email)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User successfully added","access_token":jwtToken})
}

func (userController UserController)Login(c *gin.Context){
	var loginDto dto.LoginDTO
	if err := c.ShouldBindJSON(&loginDto); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err := validate.Struct(loginDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userController.userService.Login(loginDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtToken, err := userController.jwt.CreateToken(loginDto.Email)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully","access_token":jwtToken})
}

func (userController UserController)DeleteUser(c *gin.Context){
	email, ok := c.Get("Email")
	if !ok{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing email in token"})
		return
	}
	err := userController.userService.DeleteUser(email.(string))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
}
