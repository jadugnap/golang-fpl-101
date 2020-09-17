// Package element provides structures to manage element-summary
package element

// RawEndpoint to get element-summary response
var (
	RawEndpoint = "https://fantasy.premierleague.com/api/element-summary/%d/"
)

// SummaryResponse ... to skip go-lint
type SummaryResponse struct {
	PlayerID    int
	Fixtures    []Fixture  `json:"fixtures"`
	PastMatches []History  `json:"history"`
	PastYears   []PastYear `json:"history_past"`
}

// Fixture ... to skip go-lint
type Fixture struct {
	FixtureID  int `json:"id"`
	PlayerName string
	Team       string
	Opponent   string
	Difficulty int  `json:"difficulty"`
	TeamH      int  `json:"team_h"`
	TeamA      int  `json:"team_a"`
	IsHome     bool `json:"is_home"`
	// TeamHScore           interface{} `json:"team_h_score"`
	// TeamAScore           interface{} `json:"team_a_score"`
	// Code                 int       `json:"code"`
	// Event                int       `json:"event"`
	// Finished             bool      `json:"finished"`
	// KickoffTime          time.Time `json:"kickoff_time"`
	// Minutes              int       `json:"minutes"`
	// ProvisionalStartTime bool      `json:"provisional_start_time"`
	Gameweek string `json:"event_name"`
}

// History ... to skip go-lint
type History struct {
	PlayerID   int `json:"element"`
	PlayerName string

	Value       int `json:"value"`
	TotalPoints int `json:"total_points"`
	Minutes     int `json:"minutes"`

	Team       string
	Opponent   string
	OpponentID int  `json:"opponent_team"`
	WasHome    bool `json:"was_home"`
	TeamHScore int  `json:"team_h_score"`
	TeamAScore int  `json:"team_a_score"`

	// Bps              int       `json:"bps"`
	// Creativity       string    `json:"creativity"`
	// Fixture          int       `json:"fixture"`
	// IctIndex         string    `json:"ict_index"`
	// Influence        string    `json:"influence"`
	// KickoffTime      time.Time `json:"kickoff_time"`
	// Round            int       `json:"round"`
	// Selected         int       `json:"selected"`
	// Threat           string    `json:"threat"`
	// TransfersBalance int       `json:"transfers_balance"`
	// TransfersIn      int       `json:"transfers_in"`
	// TransfersOut     int       `json:"transfers_out"`

	Assists        int `json:"assists"`
	Bonus          int `json:"bonus"`
	CleanSheets    int `json:"clean_sheets"`
	GoalsScored    int `json:"goals_scored"`
	PenaltiesSaved int `json:"penalties_saved"`
	Saves          int `json:"saves"`

	GoalsConceded   int `json:"goals_conceded"`
	OwnGoals        int `json:"own_goals"`
	PenaltiesMissed int `json:"penalties_missed"`
	RedCards        int `json:"red_cards"`
	YellowCards     int `json:"yellow_cards"`
}

// PastYear ... to skip go-lint
type PastYear struct {
	SeasonName   string `json:"season_name"`
	ElementCode  int    `json:"element_code"`
	PlayerName   string
	TeamNameNow  string
	TeamNameThen string
	// GoalsScored     int    `json:"goals_scored"`
	// Assists         int    `json:"assists"`
	// CleanSheets     int    `json:"clean_sheets"`
	// GoalsConceded   int    `json:"goals_conceded"`
	// OwnGoals        int    `json:"own_goals"`
	// PenaltiesSaved  int    `json:"penalties_saved"`
	// PenaltiesMissed int    `json:"penalties_missed"`
	// YellowCards     int    `json:"yellow_cards"`
	// RedCards        int    `json:"red_cards"`
	// Saves           int    `json:"saves"`
	// Bonus           int    `json:"bonus"`
	// Bps             int    `json:"bps"`
	// Influence       string `json:"influence"`
	// Creativity      string `json:"creativity"`
	// Threat          string `json:"threat"`
	// IctIndex        string `json:"ict_index"`
	StartCost   int `json:"start_cost"`
	EndCost     int `json:"end_cost"`
	Minutes     int `json:"minutes"`
	TotalPoints int `json:"total_points"`
}
