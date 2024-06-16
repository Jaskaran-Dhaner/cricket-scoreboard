package cricket

type IndividualRole string

const (
	Batsmen             IndividualRole = "batsmen"
	Bowlers             IndividualRole = "bowler"
	Allrounder          IndividualRole = "allrounder"
	WicketKeeperBatsmen IndividualRole = "wicketKeeperBatsmen"
)

type Player struct {
	Name           string         `json:"name" binding:"required" validate:"name" bson:"name"`
	Phone          string         `json:"phone" validate:"phone" bson:"phone"`
	IsCaptain      bool           `json:"isCaptain" validate:"required" bson:"isCaptain"`
	IsViceCap      bool           `json:"isViceCap" validate:"required" bson:"isViceCap"`
	IndividualRole IndividualRole `json:"role" validate:"required" bson:"role"`
	Email          string         `json:"email" validate:"email" bson:"email"`
	TeamId         string         `json:"teamId" validate:"required" bson:"teamId"`     // ObjectId
	MatchIds       []string       `json:"matchIds" validate:"required" bson:"matchIds"` // ObjectId
	IsActive       bool           `json:"isActive" bson:"isActive"`
	IsDeleted      bool           `json:"isDeleted" bson:"isDeleted"`
	Version        string         `json:"version" bson:"version"`
}
