package main

import (
"fmt"
"log"
"net/http"
)
func helloHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Gophers")
}
func main() {
	http.HandleFunc("\hello", helloHandler)
	if err := http.ListenAndServer(":8080", nil); err != nil {
		log.Fatal (err)
	}
}