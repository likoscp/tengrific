package services

import (
	"tengrific/db"
	"tengrific/models"
	"fmt"
	"regexp"
	"errors"
)

func GetAllUsers() ([]models.User, error) {
	if db.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	var users []models.User
	result := db.DB.Find(&users)
	return users, result.Error
}


func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	return &user, result.Error
}

func CreateUser(user *models.User) error {
	if !IsValidEmail(user.Email) {
		return errors.New("Invalid email")
}
return db.DB.Create(user).Error}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func UpdateUser(id uint, updatedData *models.User) error {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return db.DB.Model(&user).Updates(updatedData).Error
}

func DeleteUser(id uint) error {
	return db.DB.Delete(&models.User{}, id).Error
}
