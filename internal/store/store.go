package store

import (
	"database/sql"
	"log"

	"juel-ratings-api/internal/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllTeams() ([]models.Team, error) {
	log.Printf("Fetching all teams from database")
	query := `
		SELECT id, cfbd_id, school, mascot, abbreviation, conference, division, classification, city, state 
		FROM teams
	`
	rows, err := s.db.Query(query)
	if err != nil {
		log.Printf("Error querying teams: %v", err)
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var t models.Team
		err := rows.Scan(
			&t.ID,
			&t.CfbdID,
			&t.School,
			&t.Mascot,
			&t.Abbreviation,
			&t.Conference,
			&t.Division,
			&t.Classification,
			&t.City,
			&t.State)
		if err != nil {
			log.Printf("Error scanning team row: %v", err)
			return nil, err
		}
		teams = append(teams, t)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating team rows: %v", err)
		return nil, err
	}

	log.Printf("Successfully fetched %d teams from database", len(teams))
	return teams, nil
}

func (s *Store) GetTeamByID(id int) (models.Team, error) {
	log.Printf("Fetching team by id: %d", id)
	query := `
		SELECT id, cfbd_id, school, mascot, abbreviation, conference, division, classification, city, state 
		FROM teams 
		WHERE id = ?
	`
	var team models.Team
	err := s.db.QueryRow(query, id).Scan(
		&team.ID,
		&team.CfbdID,
		&team.School,
		&team.Mascot,
		&team.Abbreviation,
		&team.Conference,
		&team.Division,
		&team.Classification,
		&team.City,
		&team.State,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No team with id %d found in db", id)
			return models.Team{}, err
		}
		log.Printf("Error fetching team from db %v", err)
		return models.Team{}, err
	}

	log.Printf("Successfully fetched team with id: %d from db", id)
	return team, nil
}

func (s *Store) GetRatingsByYearAndWeek(year int, week int) ([]models.TeamRating, error) {
	query := `
		SELECT
			tr.snapshot_id,
			tr.team,
			tr.classification,
			tr.conference,
			tr.games,
			tr.drives_per_game,
			tr.off_rating,
			tr.def_rating,
			tr.sp_efficiency,
			tr.massey_results,
			tr.blended_rating,
			tr.rank,
			tr.success_rate,
			tr.explosiveness,
			tr.sp_component,
			tr.massey_component
		FROM team_ratings tr
		JOIN rating_snapshots rs ON rs.id = tr.snapshot_id
		WHERE rs.season = ?
		  AND rs.through_week = ?
		ORDER BY tr.rank ASC, tr.team ASC
	`
	rows, err := s.db.Query(query, year, week)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ratings []models.TeamRating
	for rows.Next() {
		var r models.TeamRating
		err := rows.Scan(
			&r.SnapshotID,
			&r.Team,
			&r.Classification,
			&r.Conference,
			&r.Games,
			&r.DrivesPerGame,
			&r.OffRating,
			&r.DefRating,
			&r.SPEfficiency,
			&r.MasseyResults,
			&r.BlendedRating,
			&r.Rank,
			&r.SuccessRate,
			&r.Explosiveness,
			&r.SPComponent,
			&r.MasseyComponent,
		)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ratings, nil
}
