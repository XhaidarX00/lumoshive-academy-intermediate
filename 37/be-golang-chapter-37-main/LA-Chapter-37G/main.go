// router group
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Membuat grup routing untuk /api
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "List of users"})
		})

		apiGroup.POST("/users", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{"message": "User created"})
		})

		apiGroup.GET("/products", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "List of products"})
		})
	}

	// Membuat grup routing untuk /admin
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/dashboard", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Admin dashboard"})
		})

		adminGroup.POST("/settings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Settings updated"})
		})
	}

	r.Run(":8080")
}

// // access static file
// package main

// import (
//     "net/http"

//     "github.com/gin-gonic/gin"
// )

// func main() {
//     r := gin.Default()

//     // Melayani file statis dari direktori "assets"
//     // URL prefix adalah "/static"
//     r.Static("/static", "./assets")

//     // Melayani satu file statis
//     r.StaticFile("/document", "./assets/example.pdf")

//     // Melayani file statis dari direktori "assets" menggunakan http.Dir
//     r.StaticFS("/public", http.Dir("./assets"))

//     r.Run(":8080")
// }
