package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/CrossStack-Q/Vacation-App/db"
	"github.com/CrossStack-Q/Vacation-App/models"
	"github.com/CrossStack-Q/Vacation-App/utils"
)

// DataHandler handles requests to get hotel details by ID
func DataHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	w.Header().Set("Content-Type", "application/json")

	hotelID := r.URL.Query().Get("hotelID")
	if hotelID == "" {
		http.Error(w, "Missing hotelID parameter", http.StatusBadRequest)
		return
	}

	db := db.Connect()

	defer db.Close()

	hotel, err := models.GetHotelByID(db, hotelID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Hotel not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to query database", http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(hotel); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// HotelsLocateHandler handles requests to get hotels by location
func HotelsLocateHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	w.Header().Set("Content-Type", "application/json")

	location := r.URL.Query().Get("location")
	if location != "gaza" {
		http.Error(w, "Invalid location", http.StatusBadRequest)
		return
	}

	db := db.Connect()
	defer db.Close()

	hotels, err := models.GetHotelsByLocation(db, location)
	if err != nil {
		http.Error(w, "Failed to query database", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(hotels); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
