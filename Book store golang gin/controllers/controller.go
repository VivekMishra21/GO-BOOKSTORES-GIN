package controllers

import (
	Models "book/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var book []Models.Book
	err := Models.GetALLBook(&book)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func CreateABook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.CreateABook(&book)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func GetABook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetABook(&book, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func UpdateABook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.GetABook(&book, id)
	if err != nil {
		c.JSON(http.StatusNotFound, book)
	}
	c.BindJSON(&book)
	err = Models.UpdateABook(&book, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}
func DeleteABook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.DeleteABook(&book, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
