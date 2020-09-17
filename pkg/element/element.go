// Package element provides structures and methods to manage element-summary
package element

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jadugnap/golang-fpl-101/pkg/client"
	"github.com/jadugnap/golang-fpl-101/pkg/csv"
	"github.com/jadugnap/golang-fpl-101/proto/pb"
)

// RawEndpoint to get element-summary response
var (
	RawEndpoint = "https://fantasy.premierleague.com/api/element-summary/%d/"
)

// Element for element-summary information
type Element struct {
	Client       client.GenericClient
	PlayerIDlist []int
	Res          SummaryResponse
	PlayerID     int
	PlayerName   string
	Team         string
}

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

// GetElementSummaryToCsv from api/element-summary/
func (e *Element) GetElementSummaryToCsv() {
	start := time.Now()
	defer func() {
		log.Printf("Took %v to GetResponse from %v\n", time.Since(start), e.Client.Endpoint[:len(e.Client.Endpoint)-3])
	}()

	// use WaitGroup to getResponse concurrently
	var wg sync.WaitGroup
	for _, pID := range e.PlayerIDlist {
		wg.Add(1)
		go func(pID int, wg *sync.WaitGroup) {
			defer wg.Done()

			// new instance for each local Element in each goroutine
			localE := Element{
				Client:   e.Client,
				PlayerID: pID,
			}
			localE.Client.Endpoint = fmt.Sprintf(e.Client.Endpoint, pID)
			bodyBytes := localE.Client.GetResponse()

			// define and use SummaryResponse here, no need to return
			localE.Res = SummaryResponse{}
			if err := json.Unmarshal(bodyBytes, &localE.Res); err != nil {
				log.Printf("error json.Unmarshal(): %+v\n", err)
				log.Printf("error on playerID: %+v\n", pID)
				return // from go func()
			}
			localE.fillFixturesHistoryPerPlayer()

			// store element-summary into csv
			fixturePrefix := fmt.Sprintf("fpl-players/individual/fixtures/%+v-%+v-%+v", localE.Team, localE.PlayerName, pID)
			matchPrefix := fmt.Sprintf("fpl-players/individual/pastmatches/%+v-%+v-%+v", localE.Team, localE.PlayerName, pID)
			yearPrefix := fmt.Sprintf("fpl-players/individual/pastyears/%+v-%+v-%+v", localE.Team, localE.PlayerName, pID)
			csv.StructSlice(localE.Res.Fixtures, fixturePrefix)
			csv.StructSlice(localE.Res.PastMatches, matchPrefix)
			csv.StructSlice(localE.Res.PastYears, yearPrefix)
		}(pID, &wg)
	}
	wg.Wait()
}

// fillFixturesHistoryPerPlayer from teams to element summary
func (e *Element) fillFixturesHistoryPerPlayer() {
	e.PlayerName = pb.Player_Webname_name[int32(e.PlayerID)]
	e.Team = "na"
	if len(e.Res.Fixtures) == 0 {
		fmt.Printf("No fixture found for playerName: %+v.\n", e.PlayerName)
	} else if e.Res.Fixtures[0].IsHome {
		e.Team = pb.Team_Shortname_name[int32(e.Res.Fixtures[0].TeamH)]
	} else {
		e.Team = pb.Team_Shortname_name[int32(e.Res.Fixtures[0].TeamA)]
	}

	for i, f := range e.Res.Fixtures {
		e.Res.Fixtures[i].PlayerName = e.PlayerName
		e.Res.Fixtures[i].Team = e.Team
		if e.Res.Fixtures[i].IsHome {
			e.Res.Fixtures[i].Opponent = pb.Team_Shortname_name[int32(f.TeamA)]
		} else {
			e.Res.Fixtures[i].Opponent = pb.Team_Shortname_name[int32(f.TeamH)]
		}
	}

	for i, h := range e.Res.PastMatches {
		e.Res.PastMatches[i].PlayerName = e.PlayerName
		e.Res.PastMatches[i].Team = e.Team
		e.Res.PastMatches[i].Opponent = pb.Team_Shortname_name[int32(h.OpponentID)]
	}

	for i := range e.Res.PastYears {
		e.Res.PastYears[i].PlayerName = e.PlayerName
		e.Res.PastYears[i].TeamNameNow = e.Team
		e.Res.PastYears[i].TeamNameThen = "na"
	}

	return
}
