package goriot

import (
	"fmt"
)

type masteryBook struct {
	Pages      []MasteryPage `json:"pages"`
	SummonerID int64         `json:"summonerId"`
}

//MasteryPage represents a League of Legends mastery page
type MasteryPage struct {
	Current   bool         `json:"current"`
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Masteries []MasteryDto `json:"masteries"`
}

//Mastery located inside a page
type MasteryDto struct {
	ID   int `json:"id"`
	Rank int `json:"rank"`
}

type runeBook struct {
	Pages      []RunePage `json:"pages"`
	SummonerID int        `json:"summonerId"`
}

//RunePage is a League of Legends rune page
type RunePage struct {
	Current bool       `json:"current"`
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Slots   []RuneSlot `json:"slots"`
}

//RuneSlot is a slot for a Rune to go on a RunePage
type RuneSlot struct {
	RuneId     int `json:"runeId"`
	RuneSlotID int `json:"runeSlotId"`
}

//Summoner is a player of League of Legends
type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	SummonerLevel int    `json:"summonerLevel"`
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	RevisionDate  int64  `json:"revisionDate"`
}

func (api *RiotAPI) SummonerByName(names string) (summoners map[string]Summoner, err error) {
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v1.4/summoner/by-name/%v?api_key=%v", api.Region, BaseURL, api.Region, names, api.APIkey)
	err = api.requestAndUnmarshal(url, &summoners)
	return
}

func (api *RiotAPI) SummonerByID(summonerID string) (summoners map[string]Summoner, err error) {
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v1.4/summoner/%v?api_key=%v", api.Region, BaseURL, api.Region, summonerID, api.APIkey)
	err = api.requestAndUnmarshal(url, &summoners)
	return
}

func (api *RiotAPI) RunesByID(summonerID string) (runes map[string]runeBook, err error) {
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v1.4/summoner/%v/runes?api_key=%v", api.Region, BaseURL, api.Region, summonerID, api.APIkey)
	err = api.requestAndUnmarshal(url, &runes)
	return
}

func (api *RiotAPI) MasteriesByID(summonerID string) (masteries map[string]masteryBook, err error) {
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v1.4/summoner/%v/masteries?api_key=%v", api.Region, BaseURL, api.Region, summonerID, api.APIkey)
	err = api.requestAndUnmarshal(url, &masteries)
	return
}
