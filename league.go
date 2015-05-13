package goriot

import (
	"fmt"
)

//League represents a League of Legends league
type League struct {
	Entries       []LeagueItem `json:"entries"`
	Name          string       `json:"name"`
	ParticipantId string       `json:"participantId"`
	Queue         string       `json:"queue"`
	Tier          string       `json:"tier"`
}

//LeagueItem is an entry in a League. It represents a player or team
type LeagueItem struct {
	Division         string     `json:"division"`
	IsFreshBlood     bool       `json:"isFreshBlood"`
	IsHotStreak      bool       `json:"isHotStreak"`
	IsInactive       bool       `json:"isInactive"`
	IsVeteran        bool       `json:"isVeteran"`
	LeaguePoints     int        `json:"leaguePoints"`
	Losses           int        `json:"losses"`
	MiniSeries       MiniSeries `json:"miniSeries"`
	PlayerOrTeamID   string     `json:"playerOrTeamId"`
	PlayerOrTeamName string     `json:"playerOrTeamName"`
	Wins             int        `json:"wins"`
}

//MiniSeries shows if a player is in their Series and how far they are within it
type MiniSeries struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

func (api *RiotAPI) LeagueBySummoner(summonerID string) (leagues map[string][]League, err error) {
	url := fmt.Sprintf("https://%v.%v/lol/%v/v2.5/league/by-summoner/%v?api_key=%v", api.Region, BaseURL, api.Region, summonerID, api.APIkey)
	err = api.requestAndUnmarshal(url, &leagues)
	return
}
