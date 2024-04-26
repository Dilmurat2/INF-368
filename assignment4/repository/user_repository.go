package repository

import (
	"assingment4/config"
	"assingment4/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

type UserRepository interface {
	CreateUser(user *models.User) (int32, error)
	GetUserById(id int32) (*models.User, error)
	GetUsersList() (*[]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(config *config.Config) (UserRepository, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.Name)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}
	return &userRepository{db: db}, nil
}

func (u userRepository) CreateUser(user *models.User) (int32, error) {
	var id int32
	query := `INSERT INTO users (name, email) VALUES (?, ?) RETURNING id`
	result := u.db.Raw(query, user.Name, user.Email).Scan(&id)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			return 0, fmt.Errorf("пользователь с таким email уже существует")
		}
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("запись не была добавлена")
	}
	return id, nil
}
func (u userRepository) GetUserById(id int32) (*models.User, error) {
	var user models.User
	result := u.db.First(&user, id)
	if result.Error != nil {
		return &models.User{}, result.Error
	}
	return &user, nil
}

func (u userRepository) GetUsersList() (*[]models.User, error) {
	var users []models.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}
