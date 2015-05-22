package goriot

import (
	"errors"
	"fmt"
)

type PlayerHistory struct {
	Matches []MatchSummary `json:"matches"`
}

type MatchSummary struct {
	MapID                 int                   `json:"mapId"`
	MatchCreation         int64                 `json:"matchCreation"`
	MatchDuration         int64                 `json:"matchDuration"`
	MatchId               int64                 `json:"matchId"`
	MatchMode             string                `json:"matchMode"`
	MatchType             string                `json:"matchType"`
	MatchVersion          string                `json:"matchVersion"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json:"participants"`
	PlatformId            string                `json:"platformId"`
	QueueType             string                `json:"queueType"`
	Region                string                `json:"region"`
	Season                string                `json:"season"`
}

// Retrieve match history by summoner ID
// Syntaxe : myapi.MatchHistoryBySummonerID(summonerID string [, championIds string [,rankedQueues string [, beginIndex string [, endIndex string]]]])
//
// The maximum range for begin and end index is 15. If the range is more than 15, the end index will be modified to enforce the 15 game limit. If only one of the index parameters is specified, the other will be computed accordingly.
func (api *RiotAPI) MatchHistoryBySummonerID(options ...string) (playerHistory PlayerHistory, err error) {
	nbargs := len(options)
	if 1 > nbargs || nbargs > 5 {
		err = errors.New("Incorrect number of parameters.")
		return
	}

	// default values
	summonerID := "22642975"
	championids := ""
	queue := "RANKED_SOLO_5x5,RANKED_TEAM_3x3,RANKED_TEAM_5x5"
	begin := ""
	end := ""
	for i, p := range options {
		switch i {
		case 0:
			summonerID = p
		case 1:
			championids = p
		case 2:
			queue = p
		case 3:
			begin = p
		case 4:
			end = p
		}
	}
	args := "&rankedQueues=" + queue + "&championIds=" + championids + "&beginIndex=" + begin + "&endIndex=" + end
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v2.2/matchhistory/%v?api_key=%v", api.Region, BaseURL, api.Region, summonerID, api.APIkey+args)
	err = api.requestAndUnmarshal(url, &playerHistory)
	return
}
