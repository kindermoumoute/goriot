package goriot

import (
	"fmt"
)

//Champion represents a League of Legends champion
type Champion struct {
	Active            bool `json:"active"`
	BotEnabled        bool `json:"botEnabled"`
	BotMmEnabled      bool `json:"botMmEnabled"`
	FreeToPlay        bool `json:"freeToPlay"`
	ID                int  `json:"id"`
	RankedPlayEnabled bool `json:"rankedPlayEnabled"`
}

type championList struct {
	Champions []Champion
}

func (api *RiotAPI) ChampionList(freetoplay string) (champions championList, err error) {
	args := "&freeToPlay=t" + freetoplay
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v1.2/champion?api_key=%v", api.Region, BaseURL, api.Region, api.APIkey+args)
	err = api.requestAndUnmarshal(url, &champions)
	return
}

func (api *RiotAPI) ChampionByID(championID string) (champion Champion, err error) {
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v1.2/champion/%v?api_key=%v", api.Region, BaseURL, api.Region, championID, api.APIkey)
	err = api.requestAndUnmarshal(url, &champion)
	return
}
