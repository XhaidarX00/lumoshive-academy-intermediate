// http response
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
			"status":  "success",
		})
	})

	response := gin.H{
		"message": "Hello, World!",
		"status":  "success",
	}

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, response)
	})

	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.lumoshive.com")
	})

	r.Run(":8080")
}

// // download file
// package main

// import (
//     "github.com/gin-gonic/gin"
// )

// func main() {
//     r := gin.Default()

//     r.GET("/review-file", func(c *gin.Context) {
//         filePath := "./files/example.pdf"
//         c.File(filePath)
//     })

//     r.GET("/download-file", func(c *gin.Context) {
//         filePath := "./files/example.pdf"
//         fileName := "dokumen_lumoshive.pdf"
//         c.FileAttachment(filePath, fileName)
//     })

//     r.Run(":8080")
// }
