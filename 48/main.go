package main

import (
	"belajar-golang-chapter-48/helper"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helper.HandleHome)
	http.HandleFunc("/ws", helper.HandleConnections)

	go helper.HandleMessages()

	log.Println("Server starting at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// package main

// import (
// 	"belajar-golang-chapter-48/server"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	hub := server.NewHub()
// 	go hub.Run()

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "client/index.html")
// 	})

// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		server.ServeWs(hub, w, r)
// 	})

// 	log.Println("Server started on :8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
