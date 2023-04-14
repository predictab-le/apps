package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"go.flipt.io/flipt/rpc/flipt"
	sdk "go.flipt.io/flipt/sdk/go"
)

func main() {
	var (
		features = sdk.New(nil).Flipt()
		db       *sql.DB
		openAI   *http.Client
	)

	http.HandleFunc("/predict", func(w http.ResponseWriter, r *http.Request) {
		resp, err := features.Evaluate(r.Context(), &flipt.EvaluationRequest{
			FlagKey: "predictionSource",
			Context: map[string]string{
				"organization": getOrganization(r.Context()),
			},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var prediction Prediction
		switch resp.Value {
		case "fromDB":
			prediction, err = predictionFromDB(db)
		case "fromOpenAI":
			prediction, err = predictionFromOpenAI(openAI)
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&prediction); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
