package services

import (
	"tengrific/db"
	"tengrific/models"
)

func GetAllUsers() ([]models.User, error) {
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
	return db.DB.Create(user).Error
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
