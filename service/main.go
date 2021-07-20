package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start server")

	http.HandleFunc("/health/", handler)
	http.ListenAndServe(":80", nil)
}

type Status struct {
	Status string
}

func handler(w http.ResponseWriter, r *http.Request) {
	status := Status{"OK"}

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
