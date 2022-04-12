package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/models"
	"net/http"
	"strconv"
)

func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.GetAllBooks())
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Data can not created"})
		return
	}

	models.AddNewBook(&newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func ShowBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	book, err := models.GetBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Data Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	book, err := models.RemoveBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Data Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"msg": "Data " + book.Title + " has been deleted"})
}

func UpdateBook(c *gin.Context) {
	var updateBook models.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := c.BindJSON(&updateBook); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Data can not created"})
		return
	}

	book, errUpdate := models.UpdateBookById(id, &updateBook)
	if errUpdate != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Data Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"msg": "Data " + book.Title + " has been updated"})
}
