package cricket

type victoryType string

const (
	ByRuns     victoryType = "ByRuns"
	ByWkts     victoryType = "ByWkts"
	BySuperOvr victoryType = "BySuperOvr"
)

type team struct {
	Name string `json:"name" validate:"required" bson:"name"`
	Id   string `json:"id" validate:"required" bson:"id"` // ObjectId
}

type winner struct {
	Name   string      `json:"name" validate:"required" bson:"name"`
	TeamId string      `json:"teamId" validate:"required" bson:"teamId"`
	By     victoryType `json:"by" validate:"required" bson:"by"`
	Margin int         `json:"margin" validate:"required" bson:"margin"`
}

type Match struct {
	TeamA     team   `json:"teamA" validate:"required" bson:"teamA"`
	TeamB     team   `json:"teamB" validate:"required" bson:"teamB"`
	Winner    winner `json:"winner" bson:"winner"`
	StartDate string `json:"startDate" validate:"required" bson:"startDate"`
	EndDate   string `json:"endDate" validate:"required" bson:"endDate"`
	IsLive    bool   `json:"isLive" validate:"required" bson:"isLive"`
	IsTie     bool   `json:"isTie" validate:"required" bson:"isTie"`
	IsStarted bool   `json:"isStarted" validate:"required" bson:"isStarted"`
	IsEnded   bool   `json:"isEnded" validate:"required" bson:"isEnded"`
	IsPlayed  bool   `json:"isPlayed" validate:"required" bson:"isPlayed"`
	IsDeleted bool   `json:"isDeleted" validate:"required" bson:"isDeleted"`
	IsActive  bool   `json:"isActive" validate:"required" bson:"isActive"`
	Version   string `json:"version" validate:"required" bson:"version"`
}
