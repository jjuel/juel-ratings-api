package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"juel-ratings-api/internal/store"
)

type RatingsServer struct {
	Store *store.Store
}

func (rs RatingsServer) GetRatingsByYearAndWeekHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET /ratings request received")
		yearStr := r.PathValue("year")
		weekStr := r.PathValue("week")

		year, err := strconv.Atoi(yearStr)
		if err != nil {
			http.Error(w, "Invalid year", http.StatusBadRequest)
			return
		}

		week, err := strconv.Atoi(weekStr)
		if err != nil {
			http.Error(w, "Invalid week", http.StatusBadRequest)
			return
		}

		ratings, err := rs.Store.GetRatingsByYearAndWeek(year, week)
		if err != nil {
			log.Printf("Error fetching ratings: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ratings); err != nil {
			log.Printf("Error encoding ratings to JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Successfully returned %d teams", len(ratings))
	})
}
