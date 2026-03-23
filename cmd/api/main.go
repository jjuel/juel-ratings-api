package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"juel-ratings-api/internal/handlers"
	"juel-ratings-api/internal/store"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Printf("Starting Juel Ratings API server...")
	mux := http.NewServeMux()
	db := dbConnect(getDatabaseURL())
	log.Printf("Connected to database")
	store := store.NewStore(db)

	teamServer := &handlers.TeamsServer{Store: store}
	handleTeams(mux, teamServer)

	ratingsServer := &handlers.RatingsServer{Store: store}
	handleRatings(mux, ratingsServer)

	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleTeams(mux *http.ServeMux, ts *handlers.TeamsServer) {
	mux.Handle("GET /teams", ts.GetAllTeamsHandler())
	mux.Handle("GET /teams/{id}", ts.GetTeamByIDHandler())
}

func handleRatings(mux *http.ServeMux, rs *handlers.RatingsServer) {
	mux.Handle("GET /ratings/{year}/{week}", rs.GetRatingsByYearAndWeekHandler())
}

func dbConnect(url string) *sql.DB {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		log.Fatal("Error opening Sqlite connection.")
	}

	return db
}

func getDatabaseURL() string {
	loadDotEnv()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	return databaseURL
}

func loadDotEnv() {
	if os.Getenv("DATABASE_URL") != "" {
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Could not determine working directory: %v", err)
		return
	}

	for dir := wd; ; dir = filepath.Dir(dir) {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				log.Fatalf("Error loading .env file from %s: %v", envPath, err)
			}
			return
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return
		}
	}
}
