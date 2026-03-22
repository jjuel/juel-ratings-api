package store

import (
	"database/sql"
	"juel-ratings-api/internal/models"
	"log"
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

func (s *Store) GetTeamById(id int) (models.Team, error) {
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
