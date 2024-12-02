// request body
package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func main() {
	r := gin.Default()

	// Route untuk menerima JSON
	r.POST("/user-json", func(c *gin.Context) {
		var user User
		body, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusBadRequest, "Error: %s", err.Error())
		}

		if err := json.Unmarshal(body, &user); err != nil {
			c.String(http.StatusBadRequest, "Error: %s", err.Error())
			return
		}

		// Mengembalikan respons dalam format string
		c.String(http.StatusOK, "Success: Name=%s, Email=%s", user.Name, user.Email)
	})

	// Route untuk menerima XML
	r.POST("/user-xml", func(c *gin.Context) {
		var user User

		body, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusBadRequest, "Error: %s", err.Error())
		}

		if err := xml.Unmarshal(body, &user); err != nil {
			c.String(http.StatusBadRequest, "Error: %s", err.Error())
			return
		}

		// Mengembalikan respons dalam format string
		c.String(http.StatusOK, "Success: Name=%s, Email=%s", user.Name, user.Email)
	})

	r.Run(":8080")
}

// // body parser
// package main

// import (
//     "net/http"

//     "github.com/gin-gonic/gin"
// )

// type User struct {
//     Name  string `json:"name" xml:"name" form:"name" binding:"required"`
//     Email string `json:"email" xml:"email" form:"email"  binding:"required,email"`
// }

// func main() {
//     r := gin.Default()

//     r.POST("/user/json", func(c *gin.Context) {
//         var user User
//         if err := c.ShouldBindJSON(&user); err != nil {
//             c.String(http.StatusBadRequest, "Error: %s", err.Error())
//             return
//         }

//         c.String(http.StatusOK, "Name : %s Emaiil : %s", user.Name, user.Email)
//     })

//     r.POST("/user/xml", func(c *gin.Context) {
//         var user User
//         if err := c.ShouldBindXML(&user); err != nil {
//             c.String(http.StatusBadRequest, "Error: %s", err.Error())
//             return
//         }
//         c.String(http.StatusOK, "Name : %s Emaiil : %s", user.Name, user.Email)
//     })

//     r.POST("/user/form", func(c *gin.Context) {
//         var user User
//         if err := c.ShouldBind(&user); err != nil {
//             c.String(http.StatusBadRequest, "Error: %s", err.Error())
//             return
//         }
//         c.String(http.StatusOK, "Name : %s Emaiil : %s", user.Name, user.Email)
//     })

//     r.Run(":8080")
// }
