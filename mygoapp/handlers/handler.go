// package handlers

// import (
// 	"fmt"
// 	"net/http"
// )

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "this is the server expecting")
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "H, Go Server!")
// }

//	func a(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "a is forapple ")
//	}
//
// handlers/handler.go
// handlers/handler.go
package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Go Server!")
}
func SortHandler(w http.ResponseWriter, r *http.Request) {
	// Sample array to sort
	arr := []string{"banana", "apple", "orange", "grape"}

	// Sort the array
	sort.Strings(arr)

	// Display the sorted array
	fmt.Fprint(w, "Sorted Array: "+strings.Join(arr, ", "))
}
