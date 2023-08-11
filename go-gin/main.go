package main

//basic restapi using gin framework
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     string `json:"id"`
	Author string `json:"author`
	Title  string `json:"title`
}

var books = []Book{
	{Id: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{Id: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{Id: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func checkConnection(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello satyam",
	})
}
func getAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}
func createBook(c *gin.Context) {
	var book Book
	var err error
	if err != c.ShouldBindJSON(&book) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, book)

}
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.Id == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "deleted",
			})
			break
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "No book found",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", checkConnection)
	r.GET("/books", getAllBooks)
	r.POST("/books", createBook)
	r.DELETE("/books/:id", deleteBook)
	r.Run()
}
