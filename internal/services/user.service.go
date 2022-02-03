package service

import (
	"errors"
	"github.com/red-life/simple-authentication-with-jwt/internal/dto"
	"github.com/red-life/simple-authentication-with-jwt/internal/entity/models"
	"github.com/red-life/simple-authentication-with-jwt/internal/repository"
	"github.com/red-life/simple-authentication-with-jwt/internal/utils/password"
)

type IUserService interface {
	AddUser(user dto.RegisterDTO) error
	Login(user dto.LoginDTO) error
	DeleteUser(id int) error
}

func NewUserService(userRepo repository.IUserRepository) IUserService{
	return &UserService{
		repo: userRepo,
	}
}

type UserService struct {
	repo repository.IUserRepository
}

func (userService *UserService) AddUser(user dto.RegisterDTO) error{
	userModel := &models.User{
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
	}
	err := userService.repo.CreateUser(userModel)
	if err != nil{
		return err
	}
	return nil
}

func (userService UserService) Login(user dto.LoginDTO) error{
	userFound, err := userService.repo.GetUserByEmail(user.Email)
	if err != nil{
		return err
	}
	comparedPassword := password.ComparePasswords(user.Password,userFound.Password)
	if !comparedPassword{
		return errors.New("Password doesn't match")
	}
	return nil

}

func (userService UserService) DeleteUser(id int) error{
	err := userService.repo.DeleteUser(id)
	if err != nil{
		return err
	}
	return nil
}