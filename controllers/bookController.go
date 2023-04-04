package controllers

import (
	"errors"
	"net/http"
	"tugas8/database"
	"tugas8/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//
type CreateBookInput struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

// CreteBook godoc
// @Summary Post a book for the given id
// @Description Post details of a book corresponding to the given id
// @Tags book
// @Accept  json
// @Produce  json
// @Param models.Book body CreateBookInput true "Create Book"
// @Success 200 {object} models.Book
// @Failure 400 {object} string
// @Router /books [post]
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

// GetBooks godoc
// @Summary Get details
// @Description Get details of all books
// @Tags book
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Book
// @Failure 400 {object} string
// @Router /books [get]
func GetBooks(ctx *gin.Context){
	var books []models.Book
	db := database.GetDB()
	db.Order("id").Find(&books)
	ctx.JSON(http.StatusOK, books)
}

// GetBook godoc
// @Summary Get details of a book fot the given id
// @Description Get details of a book corresponding to the given id
// @Tags book
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {object} string
// @Router /books/{id} [get]
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

// UpdateBook godoc
// @Summary Update details of a book for the given id
// @Description Update details of a book corresponding to the given id
// @Tags book
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param models.Book body UpdateBookInput true "Update Book"
// @Success 200 {object} models.Book
// @Failure 400 {object} string
// @Router /books/{id} [put]
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

// DeleteBook godoc
// @Summary Delete a book for the given id
// @Description Delete details of a book corresponding to the given id
// @Tags book
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /books/{id} [delete]
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