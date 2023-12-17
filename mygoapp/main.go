// this is the main file from where prgram running begin
// it contains the entry point of our applicatoin
package main

// import (
// 	"fmt" //to display the textual data on console
// 	"net/http"

// 	/* This imports the "net/http" package, which provides functions and types for implementing HTTP server and client functionality in Go. This is essential for building web applications in Go as it allows handling requests, responses, and communication with HTTP clients.  */
// 	"mygoapp/handlers"
// ) //For accessing functionalities defined in the "handlers" package within the "mygoapp" directory.

// main.go

// import (
// 	"/LEARNGO/mygoapp/handlers"
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", handlers.HomeHandler)
// 	http.HandleFunc("a/", handlers.a)
// 	fmt.Println("Server is running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

// main.go
// package main
// "mygoapp/handlers"

import (
	"fmt"

	"net/http"

	"example.com/hello/mygoapp/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/sort", handlers.SortHandler) // New route for sorting array
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
