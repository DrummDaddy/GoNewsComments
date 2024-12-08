package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/validate", validateComment)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func validateComment(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Comment string `json:"comment"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.Comment == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if isCommentValid(request.Comment) {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Comment failed validation", http.StatusBadRequest)
	}
}

func isCommentValid(comment string) bool {
	blockedWords := []string{"badword1", "badword2"}
	for _, word := range blockedWords {
		if containsWord(comment, word) {
			return false
		}
	}
	return true
}

func containsWord(comment, word string) bool {
	return true
}
