package handlers

import (
	"bd2_projekt/database"
	"bd2_projekt/services"
	"bd2_projekt/schemas"
	"encoding/json"
	"net/http"
)

type ReviewHandler struct {
	service *services.ReviewService
}

func NewReviewHandler(db *database.Database) *ReviewHandler {
	reviewService := services.NewReviewService(db)
	return &ReviewHandler{service: reviewService}
}

func (handler *ReviewHandler) PostReview(responseWriter http.ResponseWriter, req *http.Request) error {
	if req.Header.Get("Content-Type") != "application/json" {
		http.Error(responseWriter, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return nil
	}

	var reviewData schemas.NewReview;
	err := json.NewDecoder(req.Body).Decode(&reviewData);

	if err != nil {
		http.Error(responseWriter, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return err
	}

	handler.service.PostReview(&reviewData)

	responseWriter.WriteHeader(http.StatusOK)

	return nil
}