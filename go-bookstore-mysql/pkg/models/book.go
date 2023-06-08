package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json: "author"`
	Publication string `json:"publication"`
}

func CreateBook(db *gorm.DB, Book *Book) (err error) {
	err = db.Create(Book).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllBooks(db *gorm.DB, Book *[]Book) (err error) {
	err = db.Find(Book).Error
	if err != nil {
		return err
	}
	return nil
}

func GetBookById(db *gorm.DB, Book *Book, id int) (err error) {
	err = db.Where("id=?", id).First(Book).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, Book *Book) (err error) {
	db.Save(Book)
	return nil
}

func DeleteBook(db *gorm.DB, Book *Book, id int) (err error) {
	db.Where("id=?", id).Delete(Book)
	return nil
}
