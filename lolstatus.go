package goriot

import (
	"fmt"
)

type Shard struct {
	Hostname   string   `json:"hostname"`
	Locales    []string `json:"locales"`
	Name       string   `json:"name"`
	Region_tag string   `json:"region_tag"`
	Slug       string   `json:"slug"`
}

type ShardStatus struct {
	Hostname   string    `json:"hostname"`
	Locales    []string  `json:"locales"`
	Name       string    `json:"name"`
	Region_tag string    `json:"region_tag"`
	Services   []Service `json:"services"`
	Slug       string    `json:"slug"`
}

type Service struct {
	Incidents []Incident `json:"incidents"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	Status    string     `json:"status"`
}

type Incident struct {
	Active     bool      `json:"active"`
	Created_at string    `json:"created_at"`
	Id         int64     `json:"id"`
	Updates    []Message `json:"updates"`
}

type Message struct {
	Author       string        `json:"author"`
	Content      string        `json:"content"`
	Created_at   string        `json:"created_at"`
	Severity     string        `json:"severity"`
	Id           int64         `json:"id"`
	Translations []Translation `json:"translations"`
	Updated_at   string        `json:"updated_at"`
}

type Translation struct {
	Content    string `json:"content"`
	Locale     string `json:"locale"`
	Updated_at string `json:"updated_at"`
}

func (api *RiotAPI) Shards() (shards []Shard, err error) {
	url := fmt.Sprintf("http://status.leagueoflegends.com/shards")
	err = api.requestAndUnmarshal(url, &shards)
	return
}

func (api *RiotAPI) ShardsByRegion() (shardStatus ShardStatus, err error) {
	url := fmt.Sprintf("http://status.leagueoflegends.com/shards/%v", api.Region)
	err = api.requestAndUnmarshal(url, &shardStatus)
	return
}
