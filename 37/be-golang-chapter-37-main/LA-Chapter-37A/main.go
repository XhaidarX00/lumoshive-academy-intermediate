package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	r := gin.Default()
	// r := gin.New()

	if r != nil {
		fmt.Println("success initialization gin")
	}

}
