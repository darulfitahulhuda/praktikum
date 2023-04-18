package database

import (
	"main/config"
	"main/models"
)

func GetBlogs() (interface{}, error) {
	var blogs []models.Blog

	if e := config.DB.Find(&blogs).Error; e != nil {
		return nil, e
	}
	return blogs, nil
}

func CreateBlog(blog models.Blog) (interface{}, error) {

	if e := config.DB.Create(&blog).Error; e != nil {
		return nil, e
	}

	return blog, nil

}

func GetBlogById(id int) (interface{}, error) {
	var blog models.Blog
	if e := config.DB.First(&blog, id).Error; e != nil {
		return nil, e
	}

	return blog, nil
}

func DeleteBlog(id int) (interface{}, error) {
	var blog models.Blog
	if e := config.DB.Delete(&blog, id).Error; e != nil {
		return nil, e
	}

	return blog, nil

}

func UpdateBlog(id int, data models.Blog) (interface{}, error) {
	if e := config.DB.Model(&data).Where("ID = ?", id).Updates(data).Error; e != nil {
		return nil, e
	}

	return data, nil
}
