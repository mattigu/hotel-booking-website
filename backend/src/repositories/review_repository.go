package repositories

import (
	"bd2_projekt/database"
	"bd2_projekt/schemas"
	"context"
	"github.com/jackc/pgx/v5"
	"strconv"
)

type ReviewRepository struct{
	Db *database.Database
}

func (repository *ReviewRepository) PostReview(review *schemas.NewReview) error{
	query := `INSERT INTO reviews ("username", "hotel_id", "rating", "review_text") VALUES ('`+
		review.Username + `', ` +
		strconv.Itoa(review.HotelId) +`, `+
		strconv.Itoa(review.Rating)+`, '` + 
		review.ReviewText + `');`

	args := pgx.NamedArgs{
		"username": review.Username,
		"hotelId": review.HotelId,
		"rating": review.Rating,
		"reviewText": review.ReviewText,
	}

	repository.Db.Pool().Query(context.Background(), query, args)

	return nil
}