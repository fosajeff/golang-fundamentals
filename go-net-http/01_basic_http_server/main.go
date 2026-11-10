package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// ResponseWriter is used to write response back to the client
	// then incoming request info e.g header

	if r.Method != http.MethodGet {
		http.Error(w, "Only Get is allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = w.Write([]byte("Hello from GO net/http server"))

}

func main() {

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("try going to 8080. port")

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
}
