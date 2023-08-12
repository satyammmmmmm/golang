package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// return simple text message
func helloHandler(c *gin.Context) {

	c.String(http.StatusOK, "hello satyam")
}

// extract queryparameter from url i.e http://localhost:8080/query?name=John&age=30
func queryHandler(c *gin.Context) {
	name := c.Query("name")            //If the parameter doesn't exist, it returns an empty string.
	age := c.DefaultQuery("age", "25") //If the parameter doesn't exist, it returns the provided default value
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

// Content-Type:application/x-www-form-urlencoded
func formHandler(c *gin.Context) {
	message := c.PostForm("message")
	nickname := c.DefaultPostForm("nickname", "bhand")
	c.JSON(http.StatusOK, gin.H{
		"message":  message,
		"nickname": nickname,
	})

}

// retrieves the value of the "User-Agent" header from the request
func headerHandler(c *gin.Context) {
	userAgent := c.GetHeader("User-Agent")
	c.String(http.StatusOK, "user-agent:%s", userAgent)
}
func errorHandler(c *gin.Context) {
	err := errors.New("an error occurred")
	c.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
func custommiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("customkey", "customval")
		c.Next()
	}
}
func middleware(c *gin.Context) {
	value := c.Value("customkey").(string)
	c.String(http.StatusOK, "Value from middleware: %s", value)
}
func main() {
	fmt.Println("starting server")
	r := gin.Default()

	r.GET("/hello", helloHandler)
	r.GET("/query", queryHandler)
	r.POST("/form", formHandler)
	r.GET("/header", headerHandler)
	r.GET("/error", errorHandler)
	r.Use(custommiddleware())
	r.GET("/middle", middleware)
	r.Run()

}
