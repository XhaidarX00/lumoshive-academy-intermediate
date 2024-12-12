package main

import (
	"be-golang-chapter-49/LA-Chapter-49D/config"
	"be-golang-chapter-49/LA-Chapter-49D/handler"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/send-otp", handler.OtpHandler)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
