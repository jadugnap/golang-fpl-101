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
	Client            client.GenericClient
	PlayerIDlist      []int
	Player2Proportion map[string]float64
	Team2TotalForm    map[string]float64
	Res               SummaryResponse
	CombinedFixtures  []Fixture
	FixtureList       []Fixture
	HistoryList       []History
	Team2Gw2Points    map[string]map[int]int
	PlayerID          int
	PlayerName        string
	Team              string
	Opp               string
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
	// Code                 int         `json:"code"`
	// Event                int         `json:"event"`
	// Finished             bool        `json:"finished"`
	// KickoffTime          time.Time   `json:"kickoff_time"`
	// Minutes              int         `json:"minutes"`
	// ProvisionalStartTime bool        `json:"provisional_start_time"`
	Gameweek        string `json:"event_name"`
	PredictedPoints string
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
	Round      int  `json:"round"`
	WasHome    bool `json:"was_home"`
	TeamHScore int  `json:"team_h_score"`
	TeamAScore int  `json:"team_a_score"`

	// Bps              int       `json:"bps"`
	// Creativity       string    `json:"creativity"`
	// Fixture          int       `json:"fixture"`
	// IctIndex         string    `json:"ict_index"`
	// Influence        string    `json:"influence"`
	// KickoffTime      time.Time `json:"kickoff_time"`
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

	fixtureQueue := make(chan []Fixture, len(e.PlayerIDlist))
	historyQueue := make(chan []History, len(e.PlayerIDlist))
	// use WaitGroup to getResponse concurrently
	var wg sync.WaitGroup
	for _, pID := range e.PlayerIDlist {
		// filter out entries from (f *FPL) addSummaryRow()
		if pID > 1000 {
			continue
		}

		// start WaitGroup counter & the goroutine
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
			fixtureQueue <- localE.Res.Fixtures
			historyQueue <- localE.Res.PastMatches

			// store element-summary into csv
			// fixturePrefix := fmt.Sprintf("fpl-players/individual/fixtures/%+v-%+v-%+v", localE.Team, localE.PlayerName, pID)
			matchPrefix := fmt.Sprintf("fpl-players/individual/pastmatches/%+v-%+v-%+v", localE.Team, localE.PlayerName, pID)
			yearPrefix := fmt.Sprintf("fpl-players/individual/pastyears/%+v-%+v-%+v", localE.Team, localE.PlayerName, pID)
			// csv.StructSlice(localE.Res.Fixtures, fixturePrefix)
			csv.StructSlice(localE.Res.PastMatches, matchPrefix)
			csv.StructSlice(localE.Res.PastYears, yearPrefix)
		}(pID, &wg)
	}
	wg.Wait()
	close(historyQueue)
	close(fixtureQueue)

	e.Team2Gw2Points = make(map[string]map[int]int)
	for historyList := range historyQueue {
		e.HistoryList = historyList
		e.fillOpponentPoints()
	}
	for fixtureList := range fixtureQueue {
		e.FixtureList = fixtureList
		e.PlayerName = e.FixtureList[0].PlayerName
		e.PlayerID = int(pb.Player_Webname_value[e.PlayerName])
		e.fixtureToCsv()
	}
	// fmt.Printf("Debug e.Team2Gw2Points: %+v\n", e.Team2Gw2Points)
	e.CombinedFixtures = append(e.CombinedFixtures, e.FixtureList...)
	csv.StructSlice(e.CombinedFixtures, "fpl-players/allfixtures")
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

	for i := range e.Res.Fixtures {
		e.Res.Fixtures[i].PlayerName = e.PlayerName
	}

	for i, h := range e.Res.PastMatches {
		e.Res.PastMatches[i].PlayerName = e.PlayerName
		e.Res.PastMatches[i].Team = e.Team
		opp := pb.Team_Shortname_name[int32(h.OpponentID)]
		e.Res.PastMatches[i].Opponent = opp
	}

	for i := range e.Res.PastYears {
		e.Res.PastYears[i].PlayerName = e.PlayerName
		e.Res.PastYears[i].TeamNameNow = e.Team
		e.Res.PastYears[i].TeamNameThen = "na"
	}

	return
}

// fillOpponentPoints from element summary to a map
func (e *Element) fillOpponentPoints() {
	for i, h := range e.HistoryList {
		opp := pb.Team_Shortname_name[int32(h.OpponentID)]

		// check whether inner map for opposing team "opp" exists
		if _, ok := e.Team2Gw2Points[opp]; !ok {
			// for new team, init new inner map & append GW1 totalPoints
			e.Team2Gw2Points[opp] = make(map[int]int)
			e.Team2Gw2Points[opp][e.HistoryList[i].Round] = e.HistoryList[i].TotalPoints
		} else {
			// for existing inner map, check whether gameweek key exists
			if _, ok := e.Team2Gw2Points[opp][e.HistoryList[i].Round]; !ok {
				// for new GW, create new entry
				e.Team2Gw2Points[opp][e.HistoryList[i].Round] = e.HistoryList[i].TotalPoints
			} else {
				e.Team2Gw2Points[opp][e.HistoryList[i].Round] += e.HistoryList[i].TotalPoints
			}
		}
	}
}

// calcOpponentPoints from a map obtained from element-summary info
func (e *Element) fixtureToCsv() {
	for i, f := range e.FixtureList {
		e.FixtureList[i].PlayerName = e.PlayerName
		e.Team = "na"
		e.Opp = "na"
		if e.FixtureList[i].IsHome {
			e.Team = pb.Team_Shortname_name[int32(f.TeamH)]
			e.Opp = pb.Team_Shortname_name[int32(f.TeamA)]
		} else {
			e.Team = pb.Team_Shortname_name[int32(f.TeamA)]
			e.Opp = pb.Team_Shortname_name[int32(f.TeamH)]
		}
		e.FixtureList[i].Team = e.Team
		e.FixtureList[i].Opponent = e.Opp

		teamPPG := 0.0
		oppPPG := 0.0
		oppTotalPoints := 0
		for t, innerMap := range e.Team2Gw2Points {
			if t == e.Opp {
				for _, gameweekPoint := range innerMap {
					oppTotalPoints += gameweekPoint
				}
				oppPPG = float64(oppTotalPoints) / float64(len(innerMap))
			}
			if t == e.Team {
				teamPPG = e.Team2TotalForm[e.Team] / float64(len(innerMap))
			}
		}
		e.FixtureList[i].PredictedPoints = fmt.Sprintf("%.2f", e.Player2Proportion[e.PlayerName]/100*0.5*(teamPPG+oppPPG))
		// fmt.Printf("Debug e.PlayerName: %+v | e.Player2Proportion: %+v\n", e.PlayerName, e.Player2Proportion[e.PlayerName])
		// fmt.Printf("Debug teamPPG: %+v | oppPPG: %+v\n", teamPPG, oppPPG)
	}
	e.CombinedFixtures = append(e.CombinedFixtures, e.FixtureList...)

	fixturePrefix := fmt.Sprintf("fpl-players/individual/fixtures/%+v-%+v-%+v", e.Team, e.PlayerName, e.PlayerID)
	csv.StructSlice(e.FixtureList, fixturePrefix)
}
