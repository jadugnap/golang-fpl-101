// Package fpl provides structures and methods to manage main fpl bootstrap-static
package fpl

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jadugnap/golang-fpl-101/pkg/client"
	"github.com/jadugnap/golang-fpl-101/pkg/csv"
	"github.com/jadugnap/golang-fpl-101/pkg/team"
	"github.com/jadugnap/golang-fpl-101/proto/pb"
)

// Endpoint to get main fpl response
var (
	Endpoint = "https://fantasy.premierleague.com/api/bootstrap-static/"
)

// FPL for main fpl information
type FPL struct {
	Client      client.GenericClient
	Res         Response
	Team2Player map[string][]team.Player
	Team        string
	Players     []team.Player
}

// Response from api/bootstrap-static/
type Response struct {
	// not used
	ElementStats []ElementStat `json:"element_stats"`
	Events       []Event       `json:"events"`
	GameSettings GameSetting   `json:"game_settings"`
	Phases       []Phase       `json:"phases"`
	TotalPlayers int           `json:"total_players"`
	// frequently used
	PlayerRoles []team.PlayerRoles `json:"element_types"`
	Players     []team.Player      `json:"elements"`
	Teams       []team.Team        `json:"teams"`
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

// GetFplResponseToCsv from api/bootstrap-static/
func (f *FPL) GetFplResponseToCsv() {
	start := time.Now()
	defer func() {
		log.Printf("Took %v to GetResponse from %v\n", time.Since(start), f.Client.Endpoint)
	}()

	bodyBytes := f.Client.GetResponse()
	if err := json.Unmarshal(bodyBytes, &f.Res); err != nil {
		log.Println("error json.Unmarshal():", err)
		return
	}

	f.fillPlayersPerTeam()
	for team, players := range f.Team2Player {
		f.Team = team
		f.Players = players
		f.addSummaryRow()
		teamPrefix := fmt.Sprintf("fpl-players/%+v", team)
		csv.StructSlice(f.Players, teamPrefix)
	}
	csv.StructSlice(f.Res.Players, "fpl-players/allteam")
	csv.StructSlice(f.Res.PlayerRoles, "fpl-roles")
	csv.StructSlice(f.Res.Teams, "fpl-teams")
	return
}

// fillPlayersPerTeam with positions and teams related info
// input: *FPL
func (f *FPL) fillPlayersPerTeam() {
	f.Team2Player = make(map[string][]team.Player)
	for i, p := range f.Res.Players {
		playerName := pb.Player_Webname_name[int32(p.ID)]
		teamName := pb.Team_Shortname_name[int32(p.TeamID)]
		posName := pb.Player_Position_name[int32(p.RoleID)]
		f.Res.Players[i].WebName = playerName
		f.Res.Players[i].TeamName = teamName
		f.Res.Players[i].RoleName = posName

		if existingSlice, ok := f.Team2Player[teamName]; !ok {
			// for new team, init new slice & append playerID / struct
			f.Team2Player[teamName] = []team.Player{f.Res.Players[i]}
		} else {
			// for existing team, append playerID
			f.Team2Player[teamName] = append(existingSlice, f.Res.Players[i])
		}
	}
}

// addSummaryRow with team cumulative information per team
func (f *FPL) addSummaryRow() {
	if len(f.Players) == 0 {
		log.Println("unable to addSummaryRow() due to empty f.Players")
		return
	}
	p0 := f.Players[0]
	summary := team.Player{
		ID:       p0.TeamID + 1000,
		WebName:  fmt.Sprintf("AllPlayed_%s_Players", p0.TeamName),
		TeamName: p0.TeamName,
		RoleName: "all",
		TeamID:   p0.TeamID,
		RoleID:   -1,
		// IctIndex:    "0",
		// Form:        "-1",
		// ValueForm:   "-1",
		// ValueSeason: "-1",
		Minutes:     0,
		TotalPoints: 0,
		NowCost:     0,
	}

	// for int, sum up all the values
	// for float, convert strings => sum up all floats => string
	tempIctIndex := 0.0
	tempValueForm := 0.0
	tempValueSeason := 0.0
	tempForm := 0.0
	pCount := 0.0
	for _, player := range f.Players {
		// filter out players with 0 minute
		if player.Minutes == 0 {
			continue
		}

		pCount++
		// convert strings => sum up all floats
		pIctIndex, _ := strconv.ParseFloat(player.IctIndex, 64)
		tempIctIndex += pIctIndex

		pValueForm, _ := strconv.ParseFloat(player.ValueForm, 64)
		tempValueForm += pValueForm

		pValueSeason, _ := strconv.ParseFloat(player.ValueSeason, 64)
		tempValueSeason += pValueSeason

		pForm, _ := strconv.ParseFloat(player.Form, 64)
		tempForm += pForm

		// sum up all ints
		summary.TotalPoints += player.TotalPoints
		summary.NowCost += player.NowCost
		summary.Minutes += player.Minutes
	}
	// convert floats => string
	summary.IctIndex = fmt.Sprintf("%f", tempIctIndex/pCount)
	// summary.ValueForm = fmt.Sprintf("%f", tempValueForm)
	summary.ValueForm = fmt.Sprintf("%f", tempForm*10.0/float64(summary.NowCost))
	summary.ValueSeason = fmt.Sprintf("%f", float64(summary.TotalPoints)*10.0/float64(summary.NowCost))
	// summary.ValueSeason = fmt.Sprintf("%f", tempValueSeason)
	summary.Form = fmt.Sprintf("%f", tempForm)

	f.Players = append([]team.Player{summary}, f.Players...)
	f.Res.Players = append([]team.Player{summary}, f.Res.Players...)
	return
}
