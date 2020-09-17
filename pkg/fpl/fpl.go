// Package fpl provides structures to manage main fpl bootstrap-static
package fpl

import (
	"time"
)

// Endpoint to get bootstrap-static fpl response
var (
	Endpoint = "https://fantasy.premierleague.com/api/bootstrap-static/"
)

// Response from api/bootstrap-static/
type Response struct {
	// not used
	ElementStats []ElementStat `json:"element_stats"`
	Events       []Event       `json:"events"`
	GameSettings GameSetting   `json:"game_settings"`
	Phases       []Phase       `json:"phases"`
	TotalPlayers int           `json:"total_players"`
	// frequently used
	PlayerPositions []PlayerPosition `json:"element_types"`
	Players         []Player         `json:"elements"`
	Teams           []Team           `json:"teams"`
}

// Team ... to skip go-lint
type Team struct {
	ID int `json:"id"`
	// Code                int         `json:"code"`
	// Form                interface{} `json:"form"`
	// Strength            int         `json:"strength"`
	// TeamDivision        interface{} `json:"team_division"`
	// Unavailable         bool        `json:"unavailable"`
	// StrengthOverallHome int         `json:"strength_overall_home"`
	// StrengthOverallAway int         `json:"strength_overall_away"`
	// StrengthAttackHome  int         `json:"strength_attack_home"`
	// StrengthAttackAway  int         `json:"strength_attack_away"`
	// StrengthDefenceHome int         `json:"strength_defence_home"`
	// StrengthDefenceAway int         `json:"strength_defence_away"`
	// PulseID             int         `json:"pulse_id"`
	ShortName string `json:"short_name"`
	// Played    int    `json:"played"`
	// Win       int    `json:"win"`
	// Draw      int    `json:"draw"`
	// Loss      int    `json:"loss"`
	// Points    int    `json:"points"`
	// Position  int    `json:"position"`
	Name string `json:"name"`
}

// PlayerPosition ... to skip go-lint
type PlayerPosition struct {
	ID int `json:"id"`
	// ElementCount       int    `json:"element_count"`
	// SquadMaxPlay       int    `json:"squad_max_play"`
	// SquadMinPlay       int    `json:"squad_min_play"`
	// SquadSelect        int    `json:"squad_select"`
	// SubPositionsLocked []int  `json:"sub_positions_locked"`
	// UIShirtSpecific    bool   `json:"ui_shirt_specific"`
	// PluralName        string `json:"plural_name"`
	// PluralNameShort   string `json:"plural_name_short"`
	SingularName      string `json:"singular_name"`
	SingularNameShort string `json:"singular_name_short"`
}

// Player ... to skip go-lint
type Player struct {
	ID           int    `json:"id"`
	WebName      string `json:"web_name"`
	TeamName     string
	PositionName string
	Team         int `json:"team"`
	Position     int `json:"element_type"`
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
	// Form                             string      `json:"form"`
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
	// TotalPoints       int    `json:"total_points"`
	// SelectedByPercent string `json:"selected_by_percent"`
	// TransfersIn       int    `json:"transfers_in"`
	// TransfersOut      int    `json:"transfers_out"`
	// Bonus             int    `json:"bonus"`
	// Bps               int    `json:"bps"`
	// IctIndex          string `json:"ict_index"`
	// PointsPerGame     string `json:"points_per_game"`
	NowCost     int    `json:"now_cost"`
	Minutes     int    `json:"minutes"`
	ValueForm   string `json:"value_form"`
	ValueSeason string `json:"value_season"`
}

// ElementStat ... to skip go-lint
type ElementStat struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

// GameSetting ... to skip go-lint
type GameSetting struct {
	LeagueJoinPrivateMax         int           `json:"league_join_private_max"`
	LeagueJoinPublicMax          int           `json:"league_join_public_max"`
	LeagueMaxSizePublicClassic   int           `json:"league_max_size_public_classic"`
	LeagueMaxSizePublicH2H       int           `json:"league_max_size_public_h2h"`
	LeagueMaxSizePrivateH2H      int           `json:"league_max_size_private_h2h"`
	LeagueMaxKoRoundsPrivateH2H  int           `json:"league_max_ko_rounds_private_h2h"`
	LeaguePrefixPublic           string        `json:"league_prefix_public"`
	LeaguePointsH2HWin           int           `json:"league_points_h2h_win"`
	LeaguePointsH2HLose          int           `json:"league_points_h2h_lose"`
	LeaguePointsH2HDraw          int           `json:"league_points_h2h_draw"`
	LeagueKoFirstInsteadOfRandom bool          `json:"league_ko_first_instead_of_random"`
	CupStartEventID              int           `json:"cup_start_event_id"`
	CupStopEventID               int           `json:"cup_stop_event_id"`
	CupQualifyingMethod          string        `json:"cup_qualifying_method"`
	CupType                      string        `json:"cup_type"`
	SquadSquadplay               int           `json:"squad_squadplay"`
	SquadSquadsize               int           `json:"squad_squadsize"`
	SquadTeamLimit               int           `json:"squad_team_limit"`
	SquadTotalSpend              int           `json:"squad_total_spend"`
	UICurrencyMultiplier         int           `json:"ui_currency_multiplier"`
	UIUseSpecialShirts           bool          `json:"ui_use_special_shirts"`
	UISpecialShirtExclusions     []interface{} `json:"ui_special_shirt_exclusions"`
	StatsFormDays                int           `json:"stats_form_days"`
	SysViceCaptainEnabled        bool          `json:"sys_vice_captain_enabled"`
	TransfersSellOnFee           float64       `json:"transfers_sell_on_fee"`
	LeagueH2HTiebreakStats       []string      `json:"league_h2h_tiebreak_stats"`
	Timezone                     string        `json:"timezone"`
}

// Phase ... to skip go-lint
type Phase struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StartEvent int    `json:"start_event"`
	StopEvent  int    `json:"stop_event"`
}

// Event ... to skip go-lint
type Event struct {
	ID                     int            `json:"id"`
	Name                   string         `json:"name"`
	DeadlineTime           time.Time      `json:"deadline_time"`
	AverageEntryScore      int            `json:"average_entry_score"`
	Finished               bool           `json:"finished"`
	DataChecked            bool           `json:"data_checked"`
	HighestScoringEntry    interface{}    `json:"highest_scoring_entry"`
	DeadlineTimeEpoch      int            `json:"deadline_time_epoch"`
	DeadlineTimeGameOffset int            `json:"deadline_time_game_offset"`
	HighestScore           interface{}    `json:"highest_score"`
	IsPrevious             bool           `json:"is_previous"`
	IsCurrent              bool           `json:"is_current"`
	IsNext                 bool           `json:"is_next"`
	ChipPlays              []ChipPlay     `json:"chip_plays"`
	MostSelected           int            `json:"most_selected"`
	MostTransferredIn      int            `json:"most_transferred_in"`
	TopElement             int            `json:"top_element"`
	TopElementInfo         TopElementInfo `json:"top_element_info"`
	TransfersMade          int            `json:"transfers_made"`
	MostCaptained          int            `json:"most_captained"`
	MostViceCaptained      int            `json:"most_vice_captained"`
}

// ChipPlay ... to skip go-lint
type ChipPlay struct {
	ChipName  string `json:"chip_name"`
	NumPlayed int    `json:"num_played"`
}

// TopElementInfo ... to skip go-lint
type TopElementInfo struct {
	ID     int `json:"id"`
	Points int `json:"points"`
}
