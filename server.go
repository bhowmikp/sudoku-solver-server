package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Ex: http://localhost:8000/?board=test
func handler(w http.ResponseWriter, r *http.Request) {
	board := r.URL.Query().Get("board")

	statusBool, board := solvePuzzle(board)

	var statusString string
	if statusBool {
		statusString = "success"
	} else {
		statusString = "fail"
	}

	if board == "" {
		fmt.Fprintf(w, "Example valid link /?board=..529.6......753.99...3...8.896.......79.28.......7.9.6...4...5..472......1.692..")
	} else {
		json.NewEncoder(w).Encode(map[string]string{"status": statusString, "board": board})
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Server listening at port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
