package services

import (
	"bd2_projekt/database"
	"bd2_projekt/repositories"
	"bd2_projekt/schemas"
)

type ReviewService struct {
	repository *repositories.ReviewRepository
}

func NewReviewService(db *database.Database) *ReviewService {
	reviewRepository := &repositories.ReviewRepository{Db: db}
	return &ReviewService{repository: reviewRepository}
}

func (service *ReviewService) PostReview(reviewData *schemas.NewReview) error {
	return service.repository.PostReview(reviewData);
}
