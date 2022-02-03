package service

import (
	"github.com/red-life/simple-authentication-with-jwt/internal/dto"
	"github.com/red-life/simple-authentication-with-jwt/internal/entity/models"
	"github.com/red-life/simple-authentication-with-jwt/internal/repository"
)

type IUserService interface {
	AddUser(user *models.User) error
	Login(user *dto.LoginDTO) error
	DeleteUser(id int) error
}

type UserService struct {
	repo repository.IUserRepository
}

func (userService *UserService) AddUser(user *models.User) error{

}

func (userService UserService) Login(user *dto.LoginDTO) error{

}

