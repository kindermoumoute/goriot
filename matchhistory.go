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
	MatchVersion          string                `json:"matchVersion"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json:"participants"`
	QueueType             string                `json:"queueType"`
	Region                string                `json:"region"`
	Season                string                `json:"season"`
}

func (api *RiotAPI) MatchHistoryBySummonerID(options ...interface{}) (playerHistory PlayerHistory, err error) {
	nbargs := len(options)
	if 1 > nbargs || nbargs > 5 {
		err = errors.New("Incorrect number of parameters.")
		return
	}
	summonerID := "23990649"
	championids := ""
	queue := "RANKED_SOLO_5x5,RANKED_TEAM_3x3,RANKED_TEAM_5x5"
	begin := ""
	end := ""
	for i, p := range options {
		switch i {
		case 0:
			param, ok := p.(string)
			if !ok {
				err = errors.New("1st parameter not type string.")
				return
			}
			summonerID = param
		case 1:
			param, ok := p.(string)
			if !ok {
				err = errors.New("2nd parameter not type string.")
				return
			}
			championids = param
		case 2:
			param, ok := p.(string)
			if !ok {
				err = errors.New("3rd parameter not type string.")
				return
			}
			queue = param
		case 3:
			param, ok := p.(string)
			if !ok {
				err = errors.New("4th parameter not type string.")
				return
			}
			begin = param
		case 4:
			param, ok := p.(string)
			if !ok {
				err = errors.New("5th parameter not type string.")
				return
			}
			end = param
		}
	}
	args := "&rankedQueues=" + queue + "&championIds=" + championids + "&beginIndex=" + begin + "&endIndex=" + end
	url := fmt.Sprintf("https://%v.%v/lol/%v/v2.2/matchhistory/%v?api_key=%v", api.Region, BaseURL, api.Region, summonerID, api.APIkey+args)
	err = api.requestAndUnmarshal(url, &playerHistory)
	return
}
