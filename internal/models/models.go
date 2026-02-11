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
