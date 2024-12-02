// http request
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/headers", func(c *gin.Context) {
		// Mengambil nilai dari header `Content-Type`
		contentType := c.GetHeader("Content-Type")

		// Mengambil nilai cookie `id`
		cookie, err := c.Cookie("id")
		if err != nil {
			// Jika cookie tidak ditemukan, kembalikan pesan error
			c.String(http.StatusBadRequest, "Cookie 'id' not found")
			return
		}

		// Menampilkan nilai dari Content-Type dan Cookie
		c.String(http.StatusOK, "Content-Type: %s, Cookie id: %s", contentType, cookie)
	})

	r.Run(":8080")
}

// // router parameter
// package main

// import (
//     "net/http"

//     "github.com/gin-gonic/gin"
// )

// func main() {
//     r := gin.Default()

//     // Rute dengan satu parameter `userID`
//     r.GET("/user/:userID", func(c *gin.Context) {
//         userID := c.Param("userID") // Mengambil nilai parameter `userID`
//         c.String(http.StatusOK, "User ID: %s", userID)
//     })

//     r.GET("/user/:userID/order/:orderID", func(c *gin.Context) {
//         userID := c.Param("userID")   // Mengambil parameter `userID`
//         orderID := c.Param("orderID") // Mengambil parameter `orderID`
//         c.String(http.StatusOK, "User ID: %s, Order ID: %s", userID, orderID)
//     })

//     r.Run(":8080")
// }
