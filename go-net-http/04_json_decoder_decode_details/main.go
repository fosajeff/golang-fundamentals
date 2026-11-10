package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

type TestRequest struct {
	Name string `json:"name"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Only method POST is allowed",
		})
		return
	}
	defer r.Body.Close()

	var req TestRequest

	// json.parse equivalent (don't forget to handle parsing errors)
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid json format",
		})
		return
	}
	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Name must not be empty",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"data":      req,
		"timeStamp": time.Now().UTC(),
	})
}

func main() {

	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
}
