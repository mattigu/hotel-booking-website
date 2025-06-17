package repositories

import (
	"bd2_projekt/database"
	"bd2_projekt/schemas"
	"context"
	//"github.com/jackc/pgx/v5"
	"strconv"
)

type ReviewRepository struct{
	Db *database.Database
}

func (repository *ReviewRepository) PostReview(review *schemas.NewReview) error{
	query := `INSERT INTO reviews ("username", "hotel_id", "rating", "review_text", "upload_date") VALUES ('`+
		review.Username + `', ` +
		strconv.Itoa(review.HotelId) +`, `+
		strconv.Itoa(review.Rating)+`, '` + 
		review.ReviewText + `', '`+ 
		review.UploadDate +`');`

	repository.Db.Pool().Query(context.Background(), query)

	return nil
}