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
	Current   bool      `json:"current"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Masteries []Mastery `json:"masteries"`
}

//Mastery located inside a page
type Mastery struct {
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
	args := "api_key=" + api.APIkey
	url := fmt.Sprintf("https://%v.%v/lol/%v/v1.4/summoner/by-name/%v?%v", api.Region, BaseURL, api.Region, names, args)
	err = api.requestAndUnmarshal(url, &summoners)
	return
}
