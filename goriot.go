package goriot

// Thank to TrevorSStone
// All this code is an update of github.com/TrevorSStone/goriot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	//BaseURL is the base of the url used by Riot's API service
	BaseURL = "api.pvp.net"

	//BR represents the string for the Brazilian League of Legends Servers
	BR = "br"
	//EUNE represents the string for the North Eastern European League of Legends Servers
	EUNE = "eune"
	//EUW represents the string for the West European League of Legends Servers
	EUW = "euw"
	//KR represents the string for the Korean League of Legends Servers
	KR = "kr"
	//LAN represents the string for the Latin America North League of Legends Servers
	LAN = "lan"
	//LAS represents the string for the Latin America South League of Legends Servers
	LAS = "las"
	//NA represents the string for the North American League of Legends Servers
	NA = "na"
	//OCE represents the string for the Oceanic League of Legends Servers
	OCE = "oce"
	//RU represents the string for the Russian League of Legends Servers
	RU = "ru"
	//TR represents the string for the Turkish League of Legends Servers
	TR = "tr"
	// PBE represents the string for the Public Beta Environment League of Legends Servers
	PBE = "pbe"

	//SEASON3 is the string of "SEASON3".
	SEASON3 = "SEASON3"
	//SEASON4 is the string of "SEASON2014".
	SEASON4 = "SEASON2014"
	//SEASON5 is the string of "SEASON2015".
	SEASON5 = "SEASON2015"

	//Ranked Solo 5s
	RANKED_SOLO_5x5 = "RANKED_SOLO_5x5"
	//Ranked Team 3s
	RANKED_TEAM_3x3 = "RANKED_TEAM_3x3"
	//Ranked Team 5s
	RANKED_TEAM_5x5 = "RANKED_TEAM_5x5"
)

// RiotAPI is the main object, it contains the API key, the selected region,
// and a boolean to know if you want to ignore error 429 and force
// requestAndUnmarshal to send a result
type RiotAPI struct {
	APIkey string
	Region string
	Force  bool
}

//RiotError contains the http status code of the error
type RiotError struct {
	StatusCode int
}

//	Create the API
func Get(key string, region string, force bool) *RiotAPI {
	// TODO : test if api key is valid ?
	var myapi = RiotAPI{
		APIkey: key,
		Region: region,
		Force:  force,
	}
	return &myapi
}

// Every api methods use this function, it requests and unmarshal the gotten json
func (api *RiotAPI) requestAndUnmarshal(requestURL string, v interface{}) (err error) {

	resp, err := http.Get(requestURL)
	if err != nil {
		return
	}
	for resp.StatusCode == 429 && api.Force {
		resp, err = http.Get(requestURL)
		if err != nil {
			return
		}
	}

	if resp.StatusCode != http.StatusOK {
		return RiotError{StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}
	err = json.Unmarshal(body, v)
	return
}

//Error prints the error message for a RiotError
func (err RiotError) Error() string {
	return fmt.Sprintf("Error: HTTP Status %d", err.StatusCode)
}
