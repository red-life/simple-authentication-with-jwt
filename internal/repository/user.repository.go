package repository

import (
	"gorm.io/gorm"
	"github.com/red-life/simple-authentication-with-jwt/internal/entity/models"
)

type IUserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(id int) error
}

func NewUserRepository(dbConn *gorm.DB) IUserRepository {
	return &UserRepository{
		dbConn,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (userRepo *UserRepository) CreateUser(user *models.User) error {
	result := userRepo.db.Create(&user)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func (userRepo *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	result := userRepo.db.Where("id = ?", id).Take(&user)
	if result.Error != nil{
		return nil,result.Error
	}
	return &user,nil
}

func (userRepo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := userRepo.db.Where("email = ?", email).Take(&user)
	if result.Error != nil{
		return nil,result.Error
	}
	return &user,nil
}

func (userRepo *UserRepository) DeleteUser(id int) error {
	var user models.User
	result := userRepo.db.Where("id = ?", id).Delete(&user)
	if result.Error != nil{
		return result.Error
	}
	return nil

}
