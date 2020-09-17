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
	Team       string
	Opponent   string
	OpponentID int  `json:"opponent_team"`
	Round      int  `json:"round"`
	WasHome    bool `json:"was_home"`
	TeamHScore int  `json:"team_h_score"`
	TeamAScore int  `json:"team_a_score"`
	// Fixture          int       `json:"fixture"`
	// KickoffTime      time.Time `json:"kickoff_time"`
	// GoalsScored      int       `json:"goals_scored"`
	// Assists          int       `json:"assists"`
	// CleanSheets      int       `json:"clean_sheets"`
	// GoalsConceded    int       `json:"goals_conceded"`
	// OwnGoals         int       `json:"own_goals"`
	// PenaltiesSaved   int       `json:"penalties_saved"`
	// PenaltiesMissed  int       `json:"penalties_missed"`
	// YellowCards      int       `json:"yellow_cards"`
	// RedCards         int       `json:"red_cards"`
	// Saves            int       `json:"saves"`
	// Bps              int       `json:"bps"`
	// Influence        string    `json:"influence"`
	// Creativity       string    `json:"creativity"`
	// Threat           string    `json:"threat"`
	// TransfersBalance int       `json:"transfers_balance"`
	// Selected         int       `json:"selected"`
	// TransfersIn      int       `json:"transfers_in"`
	// TransfersOut     int       `json:"transfers_out"`
	// Bonus            int       `json:"bonus"`
	// IctIndex         string    `json:"ict_index"`
	Minutes     int `json:"minutes"`
	TotalPoints int `json:"total_points"`
	Value       int `json:"value"`
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
