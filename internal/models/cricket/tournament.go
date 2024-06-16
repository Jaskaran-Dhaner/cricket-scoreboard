package cricket

type TournamentType string

const (
	Knockout   TournamentType = "knockout"
	RoundRobin TournamentType = "roundrobin"
)

type Tournament struct {
	Name           string         `json:"name" binding:"required" validate:"name" bson:"name"`
	Logo           string         `json:"logo" validate:"required" bson:"logo"`
	TournamentType TournamentType `json:"tournamentType" validate:"required" bson:"tournamentType"`
	StartDate      string         `json:"startDate" validate:"required" bson:"startDate"`
	EndDate        string         `json:"endDate" validate:"required" bson:"endDate"`
	MatchIds       []string       `json:"matchIds" bson:"matchIds"`
	TeamIds        []string       `json:"teamIds"  bson:"teamIds"`
	IsActive       bool           `json:"isActive" bson:"isActive"`
	IsDeleted      bool           `json:"isDeleted" bson:"isDeleted"`
	Version        string         `json:"version" bson:"version"`
}
