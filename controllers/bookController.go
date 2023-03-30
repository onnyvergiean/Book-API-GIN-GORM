package controllers

import (
	"errors"
	"net/http"
	"tugas8/database"
	"tugas8/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateBookInput struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

func CreateBook(ctx *gin.Context){
	var input CreateBookInput
	db := database.GetDB()
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook := models.Book{
		Title: input.Title,
		Author: input.Author,
	}
	err = db.Create(&newBook).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, newBook)
}


func GetBooks(ctx *gin.Context){
	var books []models.Book
	db := database.GetDB()
	db.Order("id").Find(&books)
	ctx.JSON(http.StatusOK, books)
}

func GetBook(ctx *gin.Context){
	var book models.Book
	db := database.GetDB()
	id := ctx.Param("id")
	err := db.Where("id = ?", id).First(&book).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, book)
}

func UpdateBook(ctx *gin.Context){
	var book models.Book
	var input UpdateBookInput
	db := database.GetDB()
	id := ctx.Param("id")

	err := db.Where("id = ?", id).First(&book).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateBook := models.Book{
		Title: input.Title,
		Author: input.Author,
	}

	err = db.Model(&book).Updates(updateBook).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)

}

func DeleteBook(ctx *gin.Context){
	var book models.Book
	db := database.GetDB()
	id := ctx.Param("id")

	err := db.Where("id = ?", id).First(&book).Error

	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Delete(&book).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}