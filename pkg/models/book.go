package models

import (
	"github.com/gyurebalint/golang_bookstore_api/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title         string `gorm:"" json:"title"`
	Description   string `gorm:"" json:"description"`
	CoverImageUrl string `gorm:"" json:"coverImageUrl"`
	Authors       string `gorm:"" json:"author"`
	Publication   string `gorm:"" json:"publication"`
	Link          string `gorm:"" json:"link"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) *Book {
	var book Book
	db := db.Where("ID=?", ID).Find(&book)

	db.Delete(&book, ID)
	return &book
}
