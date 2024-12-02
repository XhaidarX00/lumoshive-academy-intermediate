package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/get")
	r.POST("/post")
	r.PUT("/put")
	r.DELETE("/delete")
	r.PATCH("/patch")
	r.HEAD("/head")
	r.OPTIONS("/options")
	r.Run(":8080")
}

// // handler
// package main

// import (
//     "net/http"

//     "github.com/gin-gonic/gin"
// )

// func main() {
//     r := gin.Default()

//     r.GET("/get", func(c *gin.Context) {
//         c.String(http.StatusOK, "GET method")
//     })

//     r.POST("/post", func(c *gin.Context) {
//         c.String(http.StatusOK, "POST method")
//     })

//     r.PUT("/put", func(c *gin.Context) {
//         c.String(http.StatusOK, "PUT method")
//     })

//     r.DELETE("/delete", func(c *gin.Context) {
//         c.String(http.StatusOK, "DELETE method")
//     })

//     r.PATCH("/patch", func(c *gin.Context) {
//         c.String(http.StatusOK, "PATCH method")
//     })

//     r.HEAD("/head", func(c *gin.Context) {
//         c.String(http.StatusOK, "") // HEAD tidak mengembalikan body
//     })

//     r.OPTIONS("/options", func(c *gin.Context) {
//         c.String(http.StatusOK, "OPTIONS method")
//     })

//     r.Run(":8080")
// }
