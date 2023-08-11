package main

//basic restapi using gin framework
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	Id     string `json:"id"`
	Author string `json:"author`
	Title  string `json:"title`
}

func initialize() {
	var err error
	dsn := "username:password@tcp(localhost:3306)/dbname"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to commect")

	}
	db.AutoMigrate(&Book{})

}
func createBookHandler(c *gin.Context) {
	var book Book
	var err error
	if err != c.ShouldBindJSON(&book) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return

	}
	db.Create(&book)
	c.JSON(http.StatusCreated, book)
}
func listBooksHandler(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)

}
func deleteBookHandler(c *gin.Context) {
	id := c.Param("id")
	var book Book
	var err error
	if err != db.First(&book, id).Error {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}
	db.Delete(&book)
	c.Status(http.StatusNoContent)
}

func main() {
	initialize()
	r := gin.New()
	r.GET("/books", listBooksHandler)
	r.POST("/books", createBookHandler)
	r.DELETE("/books/:id", deleteBookHandler)

	r.Run()
}
