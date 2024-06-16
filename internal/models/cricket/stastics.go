package cricket

type Over struct {
	Overs int `json:"overs" validate:"required" bson:"overs"`
	Balls int `json:"balls" validate:"required" bson:"balls"`
}

type Batting struct {
	Runs  int `json:"runs" validate:"required" bson:"runs"`
	Balls int `json:"balls" validate:"required" bson:"balls"`
	Fours int `json:"fours" validate:"required" bson:"fours"`
	Sixes int `json:"sixes" validate:"required" bson:"sixes"`
}

type Bowling struct {
	Overs   Over `json:"overs" validate:"required" bson:"overs"`
	Maidens int  `json:"maidens" validate:"required" bson:"maidens"`
	Wickets int  `json:"wickets" validate:"required" bson:"wickets"`
}

type PlayerStats struct {
	PlayerId  string  `json:"playerId" validate:"required" bson:"playerId"`
	TeamId    string  `json:"teamId" validate:"required" bson:"teamId"`
	MatchId   string  `json:"matchId" validate:"required" bson:"matchId"`
	Batting   Batting `json:"batting" validate:"required" bson:"batting"`
	Bowling   Bowling `json:"bowling" validate:"required" bson:"bowling"`
	IsDeleted bool    `json:"isDeleted" validate:"required" bson:"isDeleted"`
	IsActive  bool    `json:"isActive" validate:"required" bson:"isActive"`
}
