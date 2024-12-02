// middleware
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware untuk logging
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log informasi sebelum handler
		fmt.Println("Request received")

		// Melanjutkan ke handler berikutnya
		c.Next()

		// Log informasi setelah handler
		fmt.Println("Response sent")
	}
}

func main() {
	r := gin.Default()

	// Menambahkan middleware global
	r.Use(Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}

// // minddleware routing group
// package main

// import (
//     "net/http"

//     "github.com/gin-gonic/gin"
// )

// // Middleware untuk otentikasi
// func AuthMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         token := c.GetHeader("Authorization")
//         if token != "valid-token" {
//             c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
//             c.Abort()
//             return
//         }
//         c.Next()
//     }
// }

// func main() {
//     r := gin.Default()

//     // Grup routing dengan middleware otentikasi
//     authGroup := r.Group("/auth", AuthMiddleware())
//     {
//         authGroup.GET("/profile", func(c *gin.Context) {
//             c.JSON(http.StatusOK, gin.H{"message": "User profile"})
//         })

//         authGroup.POST("/update", func(c *gin.Context) {
//             c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
//         })
//     }

//     r.Run(":8080")
// }

// // http client GET
// package main

// import (
//     "fmt"
//     "io"
//     "net/http"
// )

// func main() {
//     client := &http.Client{}

//     req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)
//     if err != nil {
//         fmt.Println("Error creating request:", err)
//         return
//     }

//     resp, err := client.Do(req)
//     if err != nil {
//         fmt.Println("Error making request:", err)
//         return
//     }
//     defer resp.Body.Close()

//     body, err := io.ReadAll(resp.Body)
//     if err != nil {
//         fmt.Println("Error reading response:", err)
//         return
//     }

//     fmt.Println("Response:", string(body))
// }

// // http client with param
// package main

// import (
//     "fmt"
//     "io"
//     "net/http"
//     "net/url"
// )

// func main() {
//     client := &http.Client{}

//     // Menambahkan parameter query ke URL
//     baseURL, _ := url.Parse("https://jsonplaceholder.typicode.com/posts")
//     params := url.Values{}
//     params.Add("userId", "1")
//     baseURL.RawQuery = params.Encode()

//     req, err := http.NewRequest("GET", baseURL.String(), nil)
//     if err != nil {
//         fmt.Println("Error creating request:", err)
//         return
//     }

//     resp, err := client.Do(req)
//     if err != nil {
//         fmt.Println("Error making request:", err)
//         return
//     }
//     defer resp.Body.Close()

//     body, err := io.ReadAll(resp.Body)
//     if err != nil {
//         fmt.Println("Error reading response:", err)
//         return
//     }

//     fmt.Println("Response:", string(body))
// }

// // http client post
// package main

// import (
//     "bytes"
//     "encoding/json"
//     "fmt"
//     "io"
//     "net/http"
// )

// func main() {
//     client := &http.Client{}

//     // Data yang akan dikirim dalam body
//     data := map[string]interface{}{
//         "title":  "foo",
//         "body":   "bar",
//         "userId": 1,
//     }

//     // Encode data ke JSON
//     jsonData, err := json.Marshal(data)
//     if err != nil {
//         fmt.Println("Error marshalling data:", err)
//         return
//     }

//     req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(jsonData))
//     if err != nil {
//         fmt.Println("Error creating request:", err)
//         return
//     }

//     req.Header.Set("Content-Type", "application/json")

//     resp, err := client.Do(req)
//     if err != nil {
//         fmt.Println("Error making request:", err)
//         return
//     }
//     defer resp.Body.Close()

//     body, err := io.ReadAll(resp.Body)
//     if err != nil {
//         fmt.Println("Error reading response:", err)
//         return
//     }

//     fmt.Println("Response:", string(body))
// }

// // http client post with form
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// )

// func main() {
// 	client := &http.Client{}

// 	// Data yang akan dikirim sebagai form data
// 	formData := url.Values{}
// 	formData.Set("name", "John Doe")
// 	formData.Set("email", "john.doe@example.com")

// 	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBufferString(formData.Encode()))
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error making request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return
// 	}

// 	fmt.Println("Response:", string(body))
// }
