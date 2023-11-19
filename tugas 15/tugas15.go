package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	http.HandleFunc("/index", index)
	fmt.Println("mulai web server pada localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	for i := 1; i < 101; i++ {
		fmt.Fprintln(w, "angka :", i)
	}

}
