// request form
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/submit", func(c *gin.Context) {
		// Mengambil data form field `name`
		name := c.PostForm("name")
		// Mengambil data form field `email` dengan default value
		email := c.DefaultPostForm("email", "example@example.com")

		c.String(http.StatusOK, "Name: %s, Email: %s", name, email)
	})

	r.Run(":8080")
}

// // request multipart form
// package main

// import (
//     "net/http"

//     "github.com/gin-gonic/gin"
// )

// func main() {
//     r := gin.Default()

//     r.POST("/upload", func(c *gin.Context) {
//         // Mengambil field teks dari form
//         title := c.PostForm("title")

//         // Mengambil file dari form field `file`
//         file, err := c.FormFile("file")
//         if err != nil {
//             c.String(http.StatusBadRequest, "error %s", err.Error())
//             return
//         }

//         // Menyimpan file ke direktori `uploads`
//         err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
//         if err != nil {
//             c.String(http.StatusInternalServerError, "File %s uploaded error", err.Error())
//             return
//         }

//         c.String(http.StatusOK, "File %s uploaded successfully dan titlte %s", file.Filename, title)
//     })

//     r.Run(":8080")
// }
