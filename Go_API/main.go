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
	Id    int    `json:"id"`
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

	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		var userObj user
		if err := c.ShouldBindJSON(&userObj); err == nil {
			fmt.Printf("user obj - %+v\n", userObj)
		} else {
			fmt.Printf("error-%+v\n", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	r.Run("localhost:8080")
}
