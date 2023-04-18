package database

import (
	"main/config"
	"main/middlewares"
	"main/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUser(user models.User) (interface{}, error) {

	if e := config.DB.Save(&user).Error; e != nil {
		return nil, e
	}

	return user, nil

}

func GetUserById(id int) (interface{}, error) {
	var user models.User
	if e := config.DB.First(&user, id).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if e := config.DB.Delete(&user, id).Error; e != nil {
		return nil, e
	}

	return user, nil

}

func UpdateUser(id int, data models.User) (interface{}, error) {
	if e := config.DB.Model(&data).Where("ID = ?", id).Updates(data).Error; e != nil {
		return nil, e
	}

	return data, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var e error
	if e = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; e != nil {
		return nil, e
	}

	user.Token, e = middlewares.CreateToken(int(user.ID))
	if e != nil {
		return nil, e
	}

	if e = config.DB.Save(user).Error; e != nil {
		return nil, e

	}

	return user, nil
}
