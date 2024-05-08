package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Program started")
	fmt.Fprintf(w, "Welcome")
	fmt.Println("Program ended")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
