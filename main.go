package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jadugnap/golang-fpl-101/pkg/csv"
	"github.com/jadugnap/golang-fpl-101/pkg/element"
	"github.com/jadugnap/golang-fpl-101/pkg/fpl"
	"github.com/jadugnap/golang-fpl-101/pkg/team"
	"github.com/jadugnap/golang-fpl-101/proto/pb"
)

func main() {
	start := time.Now()
	defer func() {
		log.Printf("Took %v overall to execute main()\n", time.Since(start))
	}()

	// define general http client
	cli := http.Client{Timeout: time.Second * 10}

	// get bootstrap-static data
	fplResp := getFplResponseToCsv(cli, fpl.Endpoint)
	if len(fplResp.Players) == 0 {
		log.Println("error executing getFplResponse().")
		return
	}

	// get element-summary data
	playerIDlist := []int{}
	for _, p := range fplResp.Players {
		playerIDlist = append(playerIDlist, p.ID)
	}
	getElementSummaryToCsv(cli, element.RawEndpoint, playerIDlist)
}

// getFplResponseToCsv struct from api/bootstrap-static/
// input: http.Client, endpoint string
// output: *fpl.Response
// action: write csv
func getFplResponseToCsv(client http.Client, endpoint string) (fplResp *fpl.Response) {
	start := time.Now()
	defer func() {
		log.Printf("Took %v to GetResponse from %v\n", time.Since(start), endpoint)
	}()

	bodyBytes := getResponse(client, endpoint)
	if err := json.Unmarshal(bodyBytes, &fplResp); err != nil {
		log.Println("error json.Unmarshal():", err)
		return
	}

	team2players := fillPlayersPerTeam(fplResp.Players)
	for team, players := range team2players {
		teamPrefix := fmt.Sprintf("fpl-players/%+v", team)
		csv.StructSlice(players, teamPrefix)
	}
	csv.StructSlice(fplResp.Players, "fpl-players/allteam")
	csv.StructSlice(fplResp.PlayerRoles, "fpl-roles")
	csv.StructSlice(fplResp.Teams, "fpl-teams")
	return fplResp
}

// fillPlayersPerTeam with positions and teams related info
// input: []team.Player{}
// output: team2players (map[string][]team.Player{})
func fillPlayersPerTeam(players []team.Player) (team2players map[string][]team.Player) {
	// team2players = make(map[string][]int)
	team2players = make(map[string][]team.Player)
	for i, p := range players {
		playerName := pb.Player_Webname_name[int32(p.ID)]
		teamName := pb.Team_Shortname_name[int32(p.TeamID)]
		posName := pb.Player_Position_name[int32(p.RoleID)]
		players[i].WebName = playerName
		players[i].TeamName = teamName
		players[i].RoleName = posName

		if existingSlice, ok := team2players[teamName]; !ok {
			// for new team, init new slice & append playerID / struct
			team2players[teamName] = []team.Player{players[i]}
		} else {
			// for existing team, append playerID
			team2players[teamName] = append(existingSlice, players[i])
		}
	}
	return team2players
}

// fillFixturesHistoryPerPlayer from teams to element summary
// input: element.SummaryResponse{}
// output: (playerName, teamName)
func fillFixturesHistoryPerPlayer(sr element.SummaryResponse, playerID int) (playerName, teamName string) {
	playerName = pb.Player_Webname_name[int32(playerID)]
	teamName = "n/a"
	if len(sr.Fixtures) == 0 {
		fmt.Printf("No fixture found for playerName: %+v.\n", playerName)
	} else if sr.Fixtures[0].IsHome {
		teamName = pb.Team_Shortname_name[int32(sr.Fixtures[0].TeamH)]
	} else {
		teamName = pb.Team_Shortname_name[int32(sr.Fixtures[0].TeamA)]
	}

	for i, f := range sr.Fixtures {
		sr.Fixtures[i].PlayerName = playerName
		sr.Fixtures[i].Team = teamName
		if sr.Fixtures[i].IsHome {
			sr.Fixtures[i].Opponent = pb.Team_Shortname_name[int32(f.TeamA)]
		} else {
			sr.Fixtures[i].Opponent = pb.Team_Shortname_name[int32(f.TeamH)]
		}
	}

	for i, h := range sr.PastMatches {
		sr.PastMatches[i].PlayerName = playerName
		sr.PastMatches[i].Team = teamName
		sr.PastMatches[i].Opponent = pb.Team_Shortname_name[int32(h.OpponentID)]
	}

	for i := range sr.PastYears {
		sr.PastYears[i].PlayerName = playerName
		sr.PastYears[i].TeamNameNow = teamName
		sr.PastYears[i].TeamNameThen = "n/a"
	}

	return
}

// getElementSummaryToCsv from api/element-summary/
// input: http.Client, endpoint string, []int
// output: void
// action: write csv
func getElementSummaryToCsv(client http.Client, rawEndpoint string, playerIDlist []int) {
	start := time.Now()
	defer func() {
		log.Printf("Took %v to GetResponse from %v\n", time.Since(start), rawEndpoint[:len(rawEndpoint)-3])
	}()

	// use WaitGroup to getResponse concurrently
	var wg sync.WaitGroup
	for _, pID := range playerIDlist {
		wg.Add(1)
		go func(pID int, wg *sync.WaitGroup) {
			defer wg.Done()

			endpoint := fmt.Sprintf(rawEndpoint, pID)
			bodyBytes := getResponse(client, endpoint)

			// define and use summaryResp here, no need to return
			summaryResp := element.SummaryResponse{}
			if err := json.Unmarshal(bodyBytes, &summaryResp); err != nil {
				log.Printf("error json.Unmarshal(): %+v\n", err)
				log.Printf("error on playerID: %+v\n", pID)
				return // from go func()
			}
			// fillFixturesHistoryPerPlayer(summaryResp, pID)
			pName, team := fillFixturesHistoryPerPlayer(summaryResp, pID)

			// store element-summary into csv
			fixturePrefix := fmt.Sprintf("fpl-players/individual/fixtures/%+v-%+v-%+v", team, pName, pID)
			matchPrefix := fmt.Sprintf("fpl-players/individual/pastmatches/%+v-%+v-%+v", team, pName, pID)
			yearPrefix := fmt.Sprintf("fpl-players/individual/pastyears/%+v-%+v-%+v", team, pName, pID)
			csv.StructSlice(summaryResp.Fixtures, fixturePrefix)
			csv.StructSlice(summaryResp.PastMatches, matchPrefix)
			csv.StructSlice(summaryResp.PastYears, yearPrefix)
		}(pID, &wg)
	}
	wg.Wait()
}

// getResponse in general []byte
// input: http.Client, endpoint string
// output: []byte
func getResponse(client http.Client, endpoint string) []byte {
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	// any non-default "User-Agent", to resolve empty response bug
	req.Header.Set("User-Agent", "")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error HTTPClient.Do(req):", err)
		return nil
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return bodyBytes
}
