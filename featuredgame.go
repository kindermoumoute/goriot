package goriot

import (
	"fmt"
)

type FeaturedGames struct {
	ClientRefreshInterval int64              `json:"clientRefreshInterval"`
	GameList              []FeaturedGameInfo `json:"gameList"`
}

type FeaturedGameInfo struct {
	BannedChampions   []FeaturedBannedChampion `json:"bannedChampions"`
	GameId            int64                    `json:"gameId"`
	GameLength        int64                    `json:"gameLength"`
	GameMode          string                   `json:"gameMode"`
	GameQueueConfigId int64                    `json:"gameQueueConfigId"`
	GameStartTime     int64                    `json:"gameStartTime"`
	GameType          string                   `json:"gameType"`
	MapId             int64                    `json:"mapId"`
	Observers         Observer                 `json:"observers"`
	Participants      []FeaturedParticipant    `json:"participants"`
	PlatformId        string                   `json:"platformId"`
}

type FeaturedBannedChampion struct {
	ChampionId int64 `json:"championId"`
	PickTurn   int   `json:"pickTurn"`
	TeamId     int64 `json:"teamId"`
}

type FeaturedParticipant struct {
	Bot           bool   `json:"bot"`
	ChampionId    int64  `json:"championId"`
	ProfileIconId int64  `json:"profileIconId"`
	Spell1Id      int64  `json:"spell1Id"`
	Spell2Id      int64  `json:"spell2Id"`
	SummonerName  string `json:"summonerName"`
	TeamId        int64  `json:"teamId"`
}

func (api *RiotAPI) FeaturedGames(region string) (featuredGames FeaturedGames, err error) {
	url := fmt.Sprintf("https://%v.%v/observer-mode/rest/featured?api_key=%v", api.Region, BaseURL, api.APIkey)
	fmt.Printf("https://%v.%v/observer-mode/rest/featured?api_key=%v", api.Region, BaseURL, api.APIkey)
	err = api.requestAndUnmarshal(url, &featuredGames)
	return
}
