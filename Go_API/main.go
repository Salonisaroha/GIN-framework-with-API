package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	type list struct {
//		Groceries  string `json:"groceries"`
//		Stationary string `json:"stationary"`
//		Items      int    `json:"items"`
//	}
type user struct {
	Id    int    `uri:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// var todos = []list{
// 	{Groceries: "potatoes", Stationary: "pencil", Items: 12},
// 	{Groceries: "Tomatoes", Stationary: "pen", Items: 18},
// 	{Groceries: "capsicum", Stationary: "scales", Items: 10},
// }

func main() {
	//r := gin.Default()
	// r.GET("/users", func(c *gin.Context) {
	// 	c.IndentedJSON(http.StatusOK, todos)
	// })
	// r.GET("/users", func(c *gin.Context) {
	// 	var val = c.Param("users")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"value": val,
	// 	})
	//c.String(http.StatusOK, "hello world")
	// })
	// r.Run("localhost:9090")

	// Bind data from JSON
	r := gin.Default()
	r.GET("/user/:id", func(c *gin.Context) {
		var userObj user
		if err := c.ShouldBindUri(&userObj); err == nil {
			fmt.Printf("user obj - %+v\n", userObj)
		} else {
			fmt.Printf("error-%+v\n", err)
		}
		if err1 := c.ShouldBindJSON(&userObj); err1 == nil {
			fmt.Printf("user obj json binding- %+v \n", userObj)
		} else {
			fmt.Printf("error-%+v \n", err1)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	r.Run("localhost:8080")
}
