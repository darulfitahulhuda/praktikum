package database

import (
	"main/config"
	"main/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func CreateBook(book models.Book) (interface{}, error) {

	if e := config.DB.Save(&book).Error; e != nil {
		return nil, e
	}

	return book, nil

}

func GetBookById(id int) (interface{}, error) {
	var book models.Book
	if e := config.DB.First(&book, id).Error; e != nil {
		return nil, e
	}

	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book models.Book
	if e := config.DB.Delete(&book, id).Error; e != nil {
		return nil, e
	}

	return book, nil

}

func UpdateBook(id int, data models.Book) (interface{}, error) {
	if e := config.DB.Model(&data).Where("ID = ?", id).Updates(data).Error; e != nil {
		return nil, e
	}

	return data, nil
}
