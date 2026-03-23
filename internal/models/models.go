package models

type Team struct {
	ID             int     `db:"id" json:"id"`
	CfbdID         int     `db:"cfbd_id" json:"cfbd_id"`
	School         string  `db:"school" json:"school"`
	Mascot         *string `db:"mascot" json:"mascot"`
	Abbreviation   *string `db:"abbreviation" json:"abbreviation"`
	Conference     *string `db:"conference" json:"conference"`
	Division       *string `db:"division" json:"division"`
	Classification *string `db:"classification" json:"classification"`
	City           *string `db:"city" json:"city"`
	State          *string `db:"state" json:"state"`
}

type RatingSnapshot struct {
	ID           int     `db:"id" json:"id,omitempty"`
	Season       int     `db:"season" json:"season,omitempty"`
	ThroughWeek  int     `db:"through_week" json:"through_week,omitempty"`
	RatingsScope string  `db:"ratings_scope" json:"ratings_scope,omitempty"`
	ModelName    string  `db:"model_name" json:"model_name,omitempty"`
	ModelVersion string  `db:"model_version" json:"model_version,omitempty"`
	BlendAlpha   float64 `db:"blend_alpha" json:"blend_alpha,omitempty"`
	CreatedAt    string  `db:"created_at" json:"created_at,omitempty"`
}

type TeamRating struct {
	SnapshotID      int      `db:"snapshot_id" json:"snapshot_id,omitempty"`
	Team            string   `db:"team" json:"team,omitempty"`
	Classification  *string  `db:"classification" json:"classification,omitempty"`
	Conference      *string  `db:"conference" json:"conference,omitempty"`
	Games           *int     `db:"games" json:"games,omitempty"`
	DrivesPerGame   *float64 `db:"drives_per_game" json:"drives_per_game,omitempty"`
	OffRating       *float64 `db:"off_rating" json:"off_rating,omitempty"`
	DefRating       *float64 `db:"def_rating" json:"def_rating,omitempty"`
	SPEfficiency    *float64 `db:"sp_efficiency" json:"sp_efficiency,omitempty"`
	MasseyResults   *float64 `db:"massey_results" json:"massey_results,omitempty"`
	BlendedRating   *float64 `db:"blended_rating" json:"blended_rating,omitempty"`
	Rank            *int     `db:"rank" json:"rank,omitempty"`
	SuccessRate     *float64 `db:"success_rate" json:"success_rate,omitempty"`
	Explosiveness   *float64 `db:"explosiveness" json:"explosiveness,omitempty"`
	SPComponent     *float64 `db:"sp_component" json:"sp_component,omitempty"`
	MasseyComponent *float64 `db:"massey_component" json:"massey_component,omitempty"`
}
