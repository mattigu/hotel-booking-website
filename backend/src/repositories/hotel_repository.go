package repositories

import (
	"bd2_projekt/database"
	"bd2_projekt/schemas"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type HotelRepository struct {
	Db *database.Database
}

func (hotelRepository *HotelRepository) GetAll() ([]schemas.Hotel, error) {
	query := "select id, name, address_id, description, star_standard from hotels"
	rows, err := hotelRepository.Db.Pool().Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unable to query hotels: %w", err)
	}
	defer rows.Close()
	hotels := []schemas.Hotel{}
	for rows.Next() {
		hotel := schemas.Hotel{}
		err := rows.Scan(
			&hotel.Id,
			&hotel.Name,
			&hotel.AddressId,
			&hotel.Description,
			&hotel.StarStandard)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (repository *HotelRepository) getAmenitiesFor(id int) []schemas.Amenities{
	query := `SELECT a.id, a.name, a.description 
	FROM amenities a INNER JOIN hotel_amenities ha on ha.amenity_id=a.id
	WHERE ha.hotel_id=@id`

	args := pgx.NamedArgs{
		"id": id,
	}

	rows, err := repository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var amenities []schemas.Amenities
	for rows.Next() {
		var amenity schemas.Amenities
		err := rows.Scan(
			&amenity.Id,
			&amenity.Name,
			&amenity.Description,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		amenities = append(amenities, amenity)
	}
	return amenities
}

func (repository *HotelRepository) getSomeReviewsFor(id int) []schemas.ReviewData{
	query := `SELECT id, username, rating, review_text, upload_date::text
	FROM reviews
	WHERE hotel_id=@id
	fetch first 10 rows only`

	args := pgx.NamedArgs{
		"id": id,
	}

	rows, err := repository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var reviews []schemas.ReviewData
	for rows.Next() {
		var review schemas.ReviewData
		err := rows.Scan(
			&review.Id,
			&review.Username,
			&review.Rating,
			&review.ReviewText,
			&review.UploadDate,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		reviews = append(reviews, review)
	}
	return reviews
}

func (repository *HotelRepository) getAddressFor(id int) schemas.AddressData{
	query := `SELECT d.city, d.street, d.house_number, c.name
	FROM hotels h 
		INNER JOIN addresses d on h.address_id=d.id
		INNER JOIN countries c on d.country_id=c.id
	WHERE h.id=@id`

	args := pgx.NamedArgs{
		"id": id,
	}

	rows, err := repository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var address schemas.AddressData
	for rows.Next() {
		err := rows.Scan(
			&address.City,
			&address.Street,
			&address.HouseNumber,
			&address.Country,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
	}
	return address
}

func (hotelRepository *HotelRepository) getRoomsForGuests(hotelId int, guests int) ([]schemas.RoomConfiguration){
	query := `SELECT single_bed_num, double_bed_num, base_price
		FROM rooms
		WHERE hotel_id=@id and single_bed_num + double_bed_num * 2 >= @guests;`
	
	args := pgx.NamedArgs{
		"id": hotelId,
		"guests": guests,
	}
	rows, err := hotelRepository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var rooms []schemas.RoomConfiguration
	for rows.Next() {
		var room schemas.RoomConfiguration
		err := rows.Scan(
			&room.SingleBeds,
			&room.DoubleBeds,
			&room.Price,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		rooms = append(rooms, room)
	}
	return rooms
}

func (repository *HotelRepository) getAverageRating(hotelId int) (float32, int, error) {
	query := `select current_rating, review_count from hotel_ratings where hotel_id=@id`

	args := pgx.NamedArgs{
		"id": hotelId,
	}
	var rating float32;
	var ratingCount int;
	err := repository.Db.Pool().QueryRow(context.Background(), query, args).Scan(
		&rating,
		&ratingCount,
		)

	if err != nil {
		return 0, 0, nil
	}
	return rating, ratingCount, nil
}

func (repository *HotelRepository) GetById(id int, guests int) (schemas.HotelSpecificData, error) {
	query := `select h.name, h.description, h.star_standard
	from hotels h inner join addresses d on h.address_id=d.id
	where h.id=@id`

	args := pgx.NamedArgs{
		"id": id,
	}
	hotel := schemas.HotelSpecificData{}
	err := repository.Db.Pool().QueryRow(context.Background(), query, args).Scan(
		&hotel.Name,
		&hotel.Description,
		&hotel.StarStandard,
		)
	hotel.PhotoUrl = "https://content.r9cdn.net/rimg/kimg/4a/83/41fc6b329baa5c28.jpg?width=1200&height=630&crop=true";
	hotel.Amenities = repository.getAmenitiesFor(id)
	hotel.Reviews = repository.getSomeReviewsFor(id)
	hotel.Address = repository.getAddressFor(id)
	hotel.Id = id
	hotel.AvgRating, hotel.RatingsCount, err = repository.getAverageRating(id)

	if err != nil {
		return schemas.HotelSpecificData{}, fmt.Errorf("unable to query hotels: %w", err)
	}
	return hotel, nil
}

func (hotelRepository *HotelRepository) GetHotelsSearchQuery(searchQuery *schemas.HotelSearchQueryDetails) ([]schemas.HotelInfo, error){
	query := `SELECT DISTINCT on (h.id) h.id, h.name, h.star_standard, r.single_bed_num, r.double_bed_num, r.base_price
		FROM hotels h 
			INNER JOIN addresses a on a.id=h.address_id
			INNER JOIN rooms r on r.hotel_id=h.id
		WHERE a.city like @city AND r.single_bed_num + r.double_bed_num * 2 >= @guests;`
	
	args := pgx.NamedArgs{
		"city": searchQuery.City,
		"guests": searchQuery.Guests,
	}
	rows, err := hotelRepository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var hotels []schemas.HotelInfo
	for rows.Next() {
		var hotel schemas.HotelInfo
		err := rows.Scan(
			&hotel.Id,
			&hotel.Name,
			&hotel.Star,
			&hotel.SingleBeds,
			&hotel.DoubleBeds,
			&hotel.Price,
		)
		hotel.PhotoUrl = "https://content.r9cdn.net/rimg/kimg/4a/83/41fc6b329baa5c28.jpg?width=1200&height=630&crop=true";
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (hotelRepository *HotelRepository) getReservedHotelsId(city string, startDate string, endDate string, ) ([]schemas.HotelInfo, error){
	//query := "SELECT id FROM hotels "
	return nil, nil
}

func (repository *HotelRepository) GetRoomConfigurations(hotelId int, guests int) ([]schemas.RoomConfiguration, error) {
	var roomConfiguration []schemas.RoomConfiguration;
	roomConfiguration = repository.getRoomsForGuests(hotelId, guests)

	return roomConfiguration, nil
}

func (repository *HotelRepository) GetAddons(hotelId int) ([]schemas.AddonData, error){
	query := `SELECT id, name, price
	FROM reservation_addons
	WHERE hotel_id = @hotelId`
	
	args := pgx.NamedArgs{
		"hotelId": hotelId,
	}
	rows, err := repository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var addons []schemas.AddonData
	for rows.Next() {
		var addon schemas.AddonData
		err := rows.Scan(
			&addon.Id,
			&addon.Name,
			&addon.Price,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		addons = append(addons, addon)
	}
	return addons, nil
}