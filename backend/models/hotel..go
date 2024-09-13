package models

import (
	"database/sql"
	"encoding/json"

	"github.com/CrossStack-Q/Vacation-App/types"
	"github.com/lib/pq"
)

// GetHotelByID retrieves a hotel by its ID
func GetHotelByID(db *sql.DB, hotelID string) (types.Hotel, error) {
	var hotel types.Hotel
	var imageUrls []string
	var reviewJsons []string

	query := `
		SELECT hotel_id,name, price, rating, landmark, images, reviews
		FROM hotels
		WHERE hotel_id = $1
	`
	err := db.QueryRow(query, hotelID).Scan(&hotel.HotelID, &hotel.Name, &hotel.Price, &hotel.Rating, &hotel.Landmark, pq.Array(&imageUrls), pq.Array(&reviewJsons))
	if err != nil {
		return hotel, err
	}

	hotel.Images = imageUrls
	for _, reviewJson := range reviewJsons {
		var review types.Review
		if err := json.Unmarshal([]byte(reviewJson), &review); err != nil {
			return hotel, err
		}
		hotel.Reviews = append(hotel.Reviews, review)
	}

	return hotel, nil
}

// GetHotelsByLocation retrieves hotels located in the specified location
func GetHotelsByLocation(db *sql.DB, location string) ([]types.Hotel, error) {
	var hotels []types.Hotel

	query := `
        SELECT hotel_id, name, price, rating, landmark, images, reviews
        FROM hotels
        WHERE landmark ILIKE $1
    `
	rows, err := db.Query(query, "%"+location+"%")
	if err != nil {
		return hotels, err
	}
	defer rows.Close()

	for rows.Next() {
		var hotel types.Hotel
		var imageUrls []string
		var reviewJsons []string

		if err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Price, &hotel.Rating, &hotel.Landmark, pq.Array(&imageUrls), pq.Array(&reviewJsons)); err != nil {
			return hotels, err
		}

		hotel.Images = imageUrls
		for _, reviewJson := range reviewJsons {
			var review types.Review
			if err := json.Unmarshal([]byte(reviewJson), &review); err != nil {
				return hotels, err
			}
			hotel.Reviews = append(hotel.Reviews, review)
		}

		hotels = append(hotels, hotel)
	}

	if err := rows.Err(); err != nil {
		return hotels, err
	}

	return hotels, nil
}
