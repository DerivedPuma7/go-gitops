package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := "Gustavo"
	age := "24"
	fmt.Fprintf(w, "Hellllllo! I'm %s, I'm %s\n Test ArgoCD 2.0\n", name, age)
}