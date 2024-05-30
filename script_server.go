package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/robertkrimen/otto"
)

func handleScript(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		// Create a new JavaScript runtime
		vm := otto.New()

		// Load and execute the received script
		script := string(body)
		_, err = vm.Run(script)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Script executed successfully.")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/execute-script", handleScript)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
