package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/taufiqade/gowallet/models"
	httpRequest "github.com/taufiqade/gowallet/models/http/request"
)

// UserRepository godoc
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository godoc
func NewUserRepository(db *gorm.DB) models.IUserRepository {
	return &UserRepository{DB: db}
}

// GetUserByID godoc
func (repo *UserRepository) GetUserByID(id int) (models.Users, error) {
	var user models.Users
	query := repo.DB.Where("id=?", id).First(&user)
	return user, query.Error
}

// Create godoc
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

// GetUserByEmail godoc
func (repo *UserRepository) GetUserByEmail(email string) (models.Users, error) {
	var user models.Users
	query := repo.DB.Where("email=?", email).First(&user)
	return user, query.Error
}
