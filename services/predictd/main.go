package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func main() {
	var db *sql.DB

	http.HandleFunc("/predict", func(w http.ResponseWriter, r *http.Request) {
		prediction, err := predictionFromDB(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&prediction); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
