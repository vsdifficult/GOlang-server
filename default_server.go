package main

import (
	"fmt"
	"io"
	"net/http"
)

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Received script:\n%s", body)
}

func admin(aw http.ResponseWriter, ar *http.Request) {
	var headerName = "Admin"

	if ar.Header.Get(headerName) != "" {
		fmt.Fprint(aw, true, headerName)
	} else {
		fmt.Fprint(aw, false)
	}
}

func main() {
	http.HandleFunc("/testing", scriptHandler)
	http.HandleFunc("/admin", admin)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
