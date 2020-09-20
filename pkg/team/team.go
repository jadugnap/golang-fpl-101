// Package team provides structures to manage team-summary
package team

// Team ... to skip go-lint
type Team struct {
	ID        int    `json:"id"`
	ShortName string `json:"short_name"`
	LongName  string `json:"name"`
	// Code                int         `json:"code"`
	// Form                interface{} `json:"form"`
	// PulseID             int         `json:"pulse_id"`
	// TeamDivision        interface{} `json:"team_division"`
	// Unavailable         bool        `json:"unavailable"`
	// Played              int `json:"played"`
	// Position            int `json:"position"`
	// Strength            int `json:"strength"`
	// StrengthOverallHome int `json:"strength_overall_home"`
	// StrengthOverallAway int `json:"strength_overall_away"`
	StrengthAttackHome  int `json:"strength_attack_home"`
	StrengthAttackAway  int `json:"strength_attack_away"`
	StrengthDefenceHome int `json:"strength_defence_home"`
	StrengthDefenceAway int `json:"strength_defence_away"`
	Win                 int `json:"win"`
	Draw                int `json:"draw"`
	Loss                int `json:"loss"`
	TeamPoints          int `json:"points"`
	FplPoints           int
}

// PlayerRoles ... to skip go-lint
type PlayerRoles struct {
	ID int `json:"id"`
	// ElementCount       int    `json:"element_count"`
	// SquadMaxPlay       int    `json:"squad_max_play"`
	// SquadMinPlay       int    `json:"squad_min_play"`
	// SquadSelect        int    `json:"squad_select"`
	// SubPositionsLocked []int  `json:"sub_positions_locked"`
	// UIShirtSpecific    bool   `json:"ui_shirt_specific"`
	// PluralName         string `json:"plural_name"`
	// PluralNameShort    string `json:"plural_name_short"`
	LongName  string `json:"singular_name"`
	ShortName string `json:"singular_name_short"`
}

// Player ... to skip go-lint
type Player struct {
	ID       int    `json:"id"`
	WebName  string `json:"web_name"`
	TeamName string
	RoleName string
	TeamID   int `json:"team"`
	RoleID   int `json:"element_type"`
	// Assists                          int         `json:"assists"`
	// ChanceOfPlayingNextRound         interface{} `json:"chance_of_playing_next_round"`
	// ChanceOfPlayingThisRound         interface{} `json:"chance_of_playing_this_round"`
	// CleanSheets                      int         `json:"clean_sheets"`
	// Code                             int         `json:"code"`
	// CostChangeEvent                  int         `json:"cost_change_event"`
	// CostChangeEventFall              int         `json:"cost_change_event_fall"`
	// CostChangeStart                  int         `json:"cost_change_start"`
	// CostChangeStartFall              int         `json:"cost_change_start_fall"`
	// CornersAndIndirectFreekicksOrder int         `json:"corners_and_indirect_freekicks_order"`
	// CornersAndIndirectFreekicksText  string      `json:"corners_and_indirect_freekicks_text"`
	// Creativity                       string      `json:"creativity"`
	// CreativityRank                   int         `json:"creativity_rank"`
	// CreativityRankType               int         `json:"creativity_rank_type"`
	// DirectFreekicksOrder             interface{} `json:"direct_freekicks_order"`
	// DirectFreekicksText              string      `json:"direct_freekicks_text"`
	// DreamteamCount                   int         `json:"dreamteam_count"`
	// EpNext                           string      `json:"ep_next"`
	// EpThis                           string      `json:"ep_this"`
	// EventPoints                      int         `json:"event_points"`
	// FirstName                        string      `json:"first_name"`
	// GoalsConceded                    int         `json:"goals_conceded"`
	// GoalsScored                      int         `json:"goals_scored"`
	// IctIndexRank                     int         `json:"ict_index_rank"`
	// IctIndexRankType                 int         `json:"ict_index_rank_type"`
	// InDreamteam                      bool        `json:"in_dreamteam"`
	// Influence                        string      `json:"influence"`
	// InfluenceRank                    int         `json:"influence_rank"`
	// InfluenceRankType                int         `json:"influence_rank_type"`
	// News                             string      `json:"news"`
	// NewsAdded                        interface{} `json:"news_added"`
	// Photo                            string      `json:"photo"`
	// Special                          bool        `json:"special"`
	// SquadNumber                      interface{} `json:"squad_number"`
	// Status                           string      `json:"status"`
	// OwnGoals                         int         `json:"own_goals"`
	// PenaltiesMissed                  int         `json:"penalties_missed"`
	// PenaltiesOrder                   interface{} `json:"penalties_order"`
	// PenaltiesSaved                   int         `json:"penalties_saved"`
	// PenaltiesText                    string      `json:"penalties_text"`
	// RedCards                         int         `json:"red_cards"`
	// Saves                            int         `json:"saves"`
	// SecondName                       string      `json:"second_name"`
	// TeamCode                         int         `json:"team_code"`
	// Threat                           string      `json:"threat"`
	// ThreatRank                       int         `json:"threat_rank"`
	// ThreatRankType                   int         `json:"threat_rank_type"`
	// TransfersInEvent                 int         `json:"transfers_in_event"`
	// TransfersOutEvent                int         `json:"transfers_out_event"`
	// YellowCards                      int         `json:"yellow_cards"`
	// SelectedByPercent                string      `json:"selected_by_percent"`
	// TransfersIn                      int         `json:"transfers_in"`
	// TransfersOut                     int         `json:"transfers_out"`
	// Bonus                            int         `json:"bonus"`
	// Bps                              int         `json:"bps"`
	PlayerCount        int
	RegularPlayerCount int
	PointsPerGame      string `json:"points_per_game"`
	OppPointsPerGame   string
	Form               string `json:"form"`
	TotalPoints        int    `json:"total_points"`
	ValueForm          string `json:"value_form"`
	ValueSeason        string `json:"value_season"`
	IctIndex           string `json:"ict_index"`
	NowCost            int    `json:"now_cost"`
	Minutes            int    `json:"minutes"`
}
