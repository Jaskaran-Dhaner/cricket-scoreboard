package cricket

type Team struct {
	Name          string   `json:"name" binding:"required" validate:"name" bson:"name"`
	Logo          string   `json:"logo" validate:"required" bson:"logo"`
	PlayerIds     []string `json:"playerIds" validate:"required" bson:"playerIds"`         // ObjectId
	MatchIds      []string `json:"matchIds" validate:"required" bson:"matchIds"`           // ObjectId
	TournamentIds []string `json:"tournamentIds" validate:"required" bson:"tournamentIds"` // ObjectId
	IsActive      bool     `json:"isActive" bson:"isActive"`
	IsDeleted     bool     `json:"isDeleted" bson:"isDeleted"`
	Version       string   `json:"version" bson:"version"`
}
