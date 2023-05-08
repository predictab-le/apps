package main

import (
	"context"
	"database/sql"
	"net/http"
)

type Prediction struct{}

func predictionFromDB(*sql.DB) (Prediction, error) {
	return Prediction{}, nil
}

func predictionFromClosedAI(*http.Client) (Prediction, error) {
	return Prediction{}, nil
}

func getOrganization(ctx context.Context) string {
	return ""
}
