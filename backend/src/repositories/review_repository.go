package repositories

import (
	"bd2_projekt/database"
	"bd2_projekt/schemas"
	"context"
	//"github.com/jackc/pgx/v5"
	"strconv"
	"time"
)

type ReviewRepository struct{
	Db *database.Database
}

func (repository *ReviewRepository) PostReview(review *schemas.NewReview) error{
	now := time.Now()
	currentDate := now.Truncate(24 * time.Hour)
	formattedDate := currentDate.Format("2006-01-02")

	query := `INSERT INTO reviews ("username", "hotel_id", "rating", "review_text", "upload_date") VALUES ('`+
		review.Username + `', ` +
		strconv.Itoa(review.HotelId) +`, `+
		strconv.Itoa(review.Rating)+`, '` + 
		review.ReviewText + `', '`+ 
		formattedDate +`');`

	repository.Db.Pool().Query(context.Background(), query)

	return nil
}