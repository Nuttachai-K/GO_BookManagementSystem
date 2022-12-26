package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/korn/learning/book_store/pkg/models"
	"github.com/korn/learning/book_store/pkg/utils"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	c.JSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		log.Println("error while parsing book ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	bookDetails, _ := models.GetBookById(ID)
	c.JSON(http.StatusOK, bookDetails)
}

func CreateBook(c *gin.Context) {
	CreateBook := &models.Book{}
	utils.ParseBody(c, CreateBook)
	b := CreateBook.CreaterBook()
	c.JSON(http.StatusOK, b)

}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var updateBook = &models.Book{}
	utils.ParseBody(c, UpdateBook)
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	if err := db.Save(&bookDetails).Error; err != nil {
		log.Printf("error updating book: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, bookDetails)
}
