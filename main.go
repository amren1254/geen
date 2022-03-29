// main file for launching the gin server
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "Pong",
	})
}
func add(c *gin.Context) {
	a, err := strconv.Atoi(c.Param("a"))
	if err != nil {
		log.Println(err.Error())
	}
	b, err := strconv.Atoi(c.Param("b"))
	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(200, a+b)
}
func user(c *gin.Context) {
	c.JSON(200, "Amrendra")
}
func main() {
	log.Println("Initialising the router")
	fmt.Println("Hello Gin server")
	router := gin.Default()
	router.GET("/ping", ping) //1
	router.GET("add/:a/:b", add)
	router.GET("/user", user)
	router.Run()
}
