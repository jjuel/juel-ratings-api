package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"juel-ratings-api/internal/store"
)

type TeamsServer struct {
	Store *store.Store
}

func (ts *TeamsServer) GetAllTeamsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET /teams request received")
		teams, err := ts.Store.GetAllTeams()
		if err != nil {
			log.Printf("Error fetching teams: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(teams); err != nil {
			log.Printf("Error encoding teams to JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Successfully returned %d teams", len(teams))
	})
}

func (ts *TeamsServer) GetTeamByIDHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid team id", http.StatusBadRequest)
			return
		}

		team, err := ts.Store.GetTeamByID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Team not found", http.StatusNotFound)
				return
			}

			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(team); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Successfully returned team: %s", team.School)
	})
}
