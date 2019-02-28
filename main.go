package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root!\n"))
	fmt.Println("Hello Root!")
}
func YourPubHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Public World!\n"))
	fmt.Println("Hello World!")
}
func YourPrivHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Private World!\n"))
	fmt.Println("Hello Private World!")
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/private", YourPrivHandler)
	r.HandleFunc("/public", YourPubHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8042", r))

}
