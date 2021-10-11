package main

import (
	"fmt"
	"net/http"
)

// servidor web
func main() {
	fmt.Println("Servidor web ...")
	http.HandleFunc("/", home)
	// puerto, handler
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
