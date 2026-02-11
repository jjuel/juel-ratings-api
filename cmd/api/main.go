package main

import (
	"database/sql"
	"fmt"
	"juel-ratings-api/internal/handlers"
	"juel-ratings-api/internal/store"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Starting Juel Ratings API server...")
	mux := http.NewServeMux()
	db := dbConnect(getDatabaseURL())
	log.Printf("Connected to database")
	store := store.NewStore(db)
	teamServer := &handlers.TeamsServer{Store: store}
	handleTeams(mux, teamServer)
	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleTeams(mux *http.ServeMux, ts *handlers.TeamsServer) {
	mux.Handle("GET /teams", ts.GetAllTeamsHandler())
	mux.Handle("GET /teams/{id}", somethingHandler())
}

func somethingHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is something.")
	})
}

func dbConnect(url string) *sql.DB {
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatal("Error opening Postgres connection.")
	}

	return db
}

func getDatabaseURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DATABASE_URL")
}
