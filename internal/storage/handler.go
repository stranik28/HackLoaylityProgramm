package storage

import (
	"github.com/stranik28/HackLoaylityProgramm/internal/models"
)

func CreateUser(user *models.User) (id int, err error) {
	if err = DB.Create(user).Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}

func UserId(username string) (exists bool, id int) {
	var user models.User
	// Check if exist and if exist return id
	DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		exists = false
		return exists, 0
	}
	id = user.Id
	return exists, id
}

func IsAdmin(id int) bool {
	var user models.User
	DB.Where("id = ?", id).First(&user)
	if user.Id == 0 {
		return false
	}
	if user.Role != "admin" {
		return false
	}
	return true
}
