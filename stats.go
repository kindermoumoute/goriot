package goriot

import (
	"fmt"
)

type playerStatsSummaryList struct {
	PlayerStatSummaries []PlayerStatsSummary `json:"playerStatSummaries"`
	SummonerID          int64                `json:"summonerId"`
}

//PlayerStatsSummary represents a summary of a League of Legends player's game stats
type PlayerStatsSummary struct {
	AggregatedStats       AggregatedStats `json:"aggregatedStats"`
	Losses                int             `json:"losses"`
	ModifyDate            int64           `json:"modifyDate"`
	PlayerStatSummaryType string          `json:"playerStatSummaryType"`
	Wins                  int             `json:"wins"`
}

//AggregatedStats contain all values for a player's game stat values
type AggregatedStats struct {
	AverageAssists              int `json:"averageAssists"`
	AverageChampionsKilled      int `json:"averageChampionsKilled"`
	AverageCombatPlayerScore    int `json:"averageCombatPlayerScore"`
	AverageNodeCapture          int `json:"averageNodeCapture"`
	AverageNodeCaptureAssist    int `json:"averageNodeCaptureAssist"`
	AverageNodeNeutralize       int `json:"averageNodeNeutralize"`
	AverageNodeNeutralizeAssist int `json:"averageNodeNeutralizeAssist"`
	AverageNumDeaths            int `json:"averageNumDeaths"`
	AverageObjectivePlayerScore int `json:"averageObjectivePlayerScore"`
	AverageTeamObjective        int `json:"averageTeamObjective"`
	AverageTotalPlayerScore     int `json:"averageTotalPlayerScore"`
	BotGamesPlayed              int `json:"botGamesPlayed"`
	KillingSpree                int `json:"killingSpree"`
	MaxAssists                  int `json:"maxAssists"`
	MaxChampionsKilled          int `json:"maxChampionsKilled"`
	MaxCombatPlayerScore        int `json:"maxCombatPlayerScore"`
	MaxLargestCriticalStrike    int `json:"maxLargestCriticalStrike"`
	MaxLargestKillingSpree      int `json:"maxLargestKillingSpree"`
	MaxNodeCapture              int `json:"maxNodeCapture"`
	MaxNodeCaptureAssist        int `json:"maxNodeCaptureAssist"`
	MaxNodeNeutralize           int `json:"maxNodeNeutralize"`
	MaxNodeNeutralizeAssist     int `json:"maxNodeNeutralizeAssist"`
	MaxNumDeaths                int `json:"maxNumDeaths"`
	MaxObjectivePlayerScore     int `json:"maxObjectivePlayerScore"`
	MaxTeamObjective            int `json:"maxTeamObjective"`
	MaxTimePlayed               int `json:"maxTimePlayed"`
	MaxTimeSpentLiving          int `json:"maxTimeSpentLiving"`
	MaxTotalPlayerScore         int `json:"maxTotalPlayerScore"`
	MostChampionKillsPerSession int `json:"mostChampionKillsPerSession"`
	MostSpellsCast              int `json:"mostSpellsCast"`
	NormalGamesPlayed           int `json:"normalGamesPlayed"`
	RankedPremadeGamesPlayed    int `json:"rankedPremadeGamesPlayed"`
	RankedSoloGamesPlayed       int `json:"rankedSoloGamesPlayed"`
	TotalAssists                int `json:"totalAssists"`
	TotalChampionKills          int `json:"totalChampionKills"`
	TotalDamageDealt            int `json:"totalDamageDealt"`
	TotalDamageTaken            int `json:"totalDamageTaken"`
	TotalDeathsPerSession       int `json:"totalDeathsPerSession"`
	TotalDoubleKills            int `json:"totalDoubleKills"`
	TotalFirstBlood             int `json:"totalFirstBlood"`
	TotalGoldEarned             int `json:"totalGoldEarned"`
	TotalHeal                   int `json:"totalHeal"`
	TotalMagicDamageDealt       int `json:"totalMagicDamageDealt"`
	TotalMinionKills            int `json:"totalMinionKills"`
	TotalNeutralMinionsKilled   int `json:"totalNeutralMinionsKilled"`
	TotalNodeCapture            int `json:"totalNodeCapture"`
	TotalNodeNeutralize         int `json:"totalNodeNeutralize"`
	TotalPentaKills             int `json:"totalPentaKills"`
	TotalPhysicalDamageDealt    int `json:"totalPhysicalDamageDealt"`
	TotalQuadraKills            int `json:"totalQuadraKills"`
	TotalSessionsLost           int `json:"totalSessionsLost"`
	TotalSessionsPlayed         int `json:"totalSessionsPlayed"`
	TotalSessionsWon            int `json:"totalSessionsWon"`
	TotalTripleKills            int `json:"totalTripleKills"`
	TotalTurretsKilled          int `json:"totalTurretsKilled"`
	TotalUnrealKills            int `json:"totalUnrealKills"`
}

//RankedStats represents a League of Legends player's statistics for ranked play
type RankedStats struct {
	Champions  []ChampionStats `json:"champions"`
	ModifyDate int64           `json:"modifyDate"`
	SummonerID int64           `json:"summonerId"`
}

//ChampionStats are the stats for a League of Legends player's champion in ranked
type ChampionStats struct {
	ID    int             `json:"id"`
	Stats AggregatedStats `json:"stats"`
}

func (api *RiotAPI) StatSummariesBySummoner(summonerID string, season string) (stats playerStatsSummaryList, err error) {
	args := "&season=" + season
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v1.3/stats/by-summoner/%v/summary?api_key=%v", api.Region, BaseURL, api.Region, summonerID, api.APIkey+args)
	err = api.requestAndUnmarshal(url, &stats)
	return
}

func (api *RiotAPI) RankedStatsBySummoner(summonerID string, season string) (stats RankedStats, err error) {
	args := "&season=" + season
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v1.3/stats/by-summoner/%v/ranked?api_key=%v", api.Region, BaseURL, api.Region, summonerID, api.APIkey+args)
	err = api.requestAndUnmarshal(url, &stats)
	return
}
