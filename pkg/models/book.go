package models

import (
	"github.com/cristianPanaite/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var deletedBook Book
	db.Where("ID=?", Id).Delete(&deletedBook)
	return deletedBook
}

func (b *Book) UpdateBook(ID int64) Book {
	var updatedBook Book
	db.Model(&updatedBook).Where("ID = ?", ID).Updates(Book{Name: b.Name, Author: b.Author, Publication: b.Publication})
	return updatedBook
}
