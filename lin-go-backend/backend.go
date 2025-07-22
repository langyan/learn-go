package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1]
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from backend %s", port)
	})
	fmt.Println("Backend listening on port:", port)
	http.ListenAndServe(":"+port, nil)
}
