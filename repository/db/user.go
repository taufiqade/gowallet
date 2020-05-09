package db

import (
	"github.com/jinzhu/gorm"
	"test/models"
	httpRequest "test/models/http/request"
	"time"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.IUserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) GetUserByID(id int) (models.Users, error) {
	var user models.Users
	query := repo.DB.Where("id=?", id).First(&user)
	return user, query.Error
}

func (repo *UserRepository) Create(data *httpRequest.UserRequest) (models.Users, error) {
	var user = models.Users{
		Email:     data.Email,
		Password:  data.Password,
		Name:      data.Name,
		UpdatedAt: time.Time{},
		CreatedAt: time.Time{},
	}
	query := repo.DB.Create(&user)
	return user, query.Error
}

func (repo *UserRepository) GetUserByEmail(email string) (models.Users, error) {
	var user models.Users
	query := repo.DB.Where("email=?", email).First(&user)
	return user, query.Error
}
