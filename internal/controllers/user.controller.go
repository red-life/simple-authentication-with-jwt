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
}

type UserController struct {
	userService service.IUserService
	jwt jwt.IJWTPackage
}

var validate *validator.Validate

func (userController UserController)AddUser(c *gin.Context){
	var registerDto dto.RegisterDTO
	err := validate.Struct(registerDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.ShouldBindJSON(&registerDto); err != nil{
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
	c.SetCookie("token",jwtToken,15000,"/","localhost",true,true)
	c.JSON(http.StatusBadRequest, gin.H{"message": "User successfully added"})
}

func (userController UserController)Login(c *gin.Context){
	var loginDto dto.LoginDTO
	err := validate.Struct(loginDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.ShouldBindJSON(&loginDto); err != nil{
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
	c.SetCookie("token",jwtToken,15000,"/","localhost",true,true)
	c.JSON(http.StatusBadRequest, gin.H{"message": "Logged in successfully"})
}
