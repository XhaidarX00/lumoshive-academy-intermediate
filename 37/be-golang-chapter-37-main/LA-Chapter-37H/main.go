// error handling
package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/data", func(c *gin.Context) {
		err := someFunctionThatMightFail()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			// c.String(http.StatusBadRequest, "Error: %s", err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "some data"})
	})

	r.Run(":8080")
}

func someFunctionThatMightFail() error {
	// return nil
	return errors.New("something error")
}

// // template engine
// package main

// import (
//     "net/http"

//     "github.com/gin-gonic/gin"
// )

// func main() {
//     r := gin.Default()

//     // Memuat template dari direktori "templates"
//     r.LoadHTMLGlob("templates/*")

//     r.GET("/", func(c *gin.Context) {
//         data := gin.H{
//             "Title":   "Welcome to Gin",
//             "Message": "This is a simple example of using templates with Gin.",
//         }
//         c.HTML(http.StatusOK, "index.html", data)
//     })

//     r.GET("/about", func(c *gin.Context) {
//         data := gin.H{
//             "Title":   "Lumoshive Academy",
//             "Message": "Selamat datang di lumoshive academy",
//         }
//         c.HTML(http.StatusOK, "about.html", data)
//     })

//     r.Run(":8080")
// }
