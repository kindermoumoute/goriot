package goriot

import (
	"fmt"
)

//Team is a League of Legends Ranked Team
type Team struct {
	CreateDate                    int64                 `json:"createDate"`
	FullID                        string                `json:"fullId"`
	LastGameDate                  int64                 `json:"lastGameDate"`
	LastJoinDate                  int64                 `json:"lastJoinDate"`
	LastJoinedRankedTeamQueueDate int64                 `json:"lastJoinedRankedTeamQueueDate"`
	MatchHistory                  []MatchHistorySummary `json:"matchHistory"`
	ModifyDate                    int64                 `json:"modifyDate"`
	Name                          string                `json:"name"`
	Roster                        Roster                `json:"roster"`
	SecondLastJoinDate            int64                 `json:"secondLastJoinDate"`
	Status                        string                `json:"status"`
	Tag                           string                `json:"tag"`
	TeamStatDetails               []TeamStatDetail      `json:"teamStatDetails"`
	ThirdJoinDate                 int64                 `json:"thirdLastJoinDate"`
}

//MatchHistorySummary is a summary of a matches played by a Team
type MatchHistorySummary struct {
	Assists           int    `json:"assists"`
	Date              int64  `json:"date"`
	Deaths            int    `json:"deaths"`
	GameID            int64  `json:"gameId"`
	GameMode          string `json:"gameMode"`
	Invalid           bool   `json:"invalid"`
	Kills             int    `json:"kills"`
	MapID             int    `json:"mapId"`
	OpposingTeamKills int    `json:"opposingTeamKills"`
	OpposingTeamName  string `json:"opposingTeamName"`
	Win               bool   `json:"win"`
}

//Roster represents the roster of a League of Legends ranked team
type Roster struct {
	MemberList []TeamMemberInfo `json:"memberList"`
	OwnerID    int64            `json:"ownerId"`
}

//TeamMemberInfo is the individual information for a player on a ranked team
type TeamMemberInfo struct {
	InviteDate int64  `json:"inviteDate"`
	JoinDate   int64  `json:"joinDate"`
	PlayerID   int64  `json:"playerId"`
	Status     string `json:"status"`
}

//TeamStatDetail is the statistics for a ranked team
type TeamStatDetail struct {
	AverageGamesPlayed int    `json:"averageGamesPlayed"`
	Losses             int    `json:"losses"`
	TeamStatType       string `json:"teamStatType"`
	Wins               int    `json:"wins"`
}

func (api *RiotAPI) TeamBySummonerID(summonerID string) (teams map[string][]Team, err error) {
	url := fmt.Sprintf("https://%v.%v/lol/%v/v2.4/team/by-summoner/%v?api_key=%v", api.Region, BaseURL, api.Region, summonerID, api.APIkey)
	err = api.requestAndUnmarshal(url, &teams)
	return
}

func (api *RiotAPI) TeamByTeamID(teamID string) (teams map[string]Team, err error) {
	url := fmt.Sprintf("https://%v.%v/lol/%v/v2.4/team/%v?api_key=%v", api.Region, BaseURL, api.Region, teamID, api.APIkey)
	err = api.requestAndUnmarshal(url, &teams)
	return
}
