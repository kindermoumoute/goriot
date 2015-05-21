package goriot

import (
	"fmt"
)

type CurrentGameInfo struct {
	BannedChampions   []CurrentBannedChampion  `json:"bannedChampions"`
	GameId            int64                    `json:"gameId"`
	GameLength        int64                    `json:"gameLength"`
	GameMode          string                   `json:"gameMode"`
	GameQueueConfigId int64                    `json:"gameQueueConfigId"`
	GameStartTime     int64                    `json:"gameStartTime"`
	GameType          string                   `json:"gameType"`
	MapId             int64                    `json:"mapId"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
	PlatformId        string                   `json:"platformId"`
}

type CurrentBannedChampion struct {
	ChampionId int   `json:"championId"`
	PickTurn   int   `json:"pickTurn"`
	TeamId     int64 `json:"teamId"`
}

type CurrentGameParticipant struct {
	Bot           bool        `json:"bot"`
	ChampionId    int64       `json:"championId"`
	Masteries     []Mastery   `json:"masteries"`
	ProfileIconId int64       `json:"profileIconId"`
	Runes         []RuneCount `json:"runes"`
	Spell1Id      int64       `json:"spell1Id"`
	Spell2Id      int64       `json:"spell2Id"`
	SummonerId    int64       `json:"summonerId"`
	SummonerName  string      `json:"summonerName"`
	TeamId        int64       `json:"teamId"`
}

type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

// Mastery alreay exist
type RuneCount struct {
	Count  int   `json:"count"`
	RuneId int64 `json:"runeId"`
}

func (api *RiotAPI) CurrentGameBySummonerID(plateform string, summonerID string) (currentGame CurrentGameInfo, err error) {
	url := fmt.Sprintf("https://%v.%v/observer-mode/rest/consumer/getSpectatorGameInfo/%v/%v?api_key=%v", api.Region, BaseURL, plateform, summonerID, api.APIkey)
	err = api.requestAndUnmarshal(url, &currentGame)
	return
}
