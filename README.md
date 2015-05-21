# goriot
go tool for riot api

install :
go get "github.com/kindermoumoute/goriot"

Using :
import "github.com/kindermoumoute/goriot"

Example :
myapi := goriot.Get(APIkey, goriot.EUW, 10, 500)

# goriot
--
    import "github.com/kindermoumoute/goriot"


## Usage

```go
const (
	//BaseURL is the base of the url used by Riot's API service
	BaseURL = "api.pvp.net"

	//BR represents the string for the Brazilian League of Legends Servers,
	//only used as a helper to prevent typos
	BR = "br"
	//EUNE represents the string for the North Eastern European League of Legends Servers,
	//only used as a helper to prevent typos
	EUNE = "eune"
	//EUW represents the string for the West European League of Legends Servers,
	//only used as a helper to prevent typos
	EUW = "euw"
	//KR represents the string for the Korean League of Legends Servers,
	//only used as a helper to prevent typos
	KR = "kr"
	//LAN represents the string for the Latin America North League of Legends Servers,
	//only used as a helper to prevent typos
	LAN = "lan"
	//LAS represents the string for the Latin America South League of Legends Servers,
	//only used as a helper to prevent typos
	LAS = "las"
	//NA represents the string for the North American League of Legends Servers,
	//only used as a helper to prevent typos
	NA = "na"
	//OCE represents the string for the Oceanic League of Legends Servers,
	//only used as a helper to prevent typos
	OCE = "oce"
	//RU represents the string for the Russian League of Legends Servers,
	//only used as a helper to prevent typos
	RU = "ru"
	//TR represents the string for the Turkish League of Legends Servers,
	//only used as a helper to prevent typos
	TR = "tr"

	// same
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
```

#### type AggregatedStats

```go
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
```

AggregatedStats contain all values for a player's game stat values

#### type BannedChampion

```go
type BannedChampion struct {
	ChampionId int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}
```


#### type Champion

```go
type Champion struct {
	Active            bool `json:"active"`
	BotEnabled        bool `json:"botEnabled"`
	BotMmEnabled      bool `json:"botMmEnabled"`
	FreeToPlay        bool `json:"freeToPlay"`
	ID                int  `json:"id"`
	RankedPlayEnabled bool `json:"rankedPlayEnabled"`
}
```

Champion represents a League of Legends champion

#### type ChampionStats

```go
type ChampionStats struct {
	ID    int             `json:"id"`
	Stats AggregatedStats `json:"stats"`
}
```

ChampionStats are the stats for a League of Legends player's champion in ranked

#### type CurrentBannedChampion

```go
type CurrentBannedChampion struct {
	ChampionId int   `json:"championId"`
	PickTurn   int   `json:"pickTurn"`
	TeamId     int64 `json:"teamId"`
}
```


#### type CurrentGameInfo

```go
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
```


#### type CurrentGameParticipant

```go
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
```


#### type Event

```go
type Event struct {
	AscendedType            string   `json:"ascendedType"`
	AssistingParticipantIDs []int    `json:"assistingParticipantIds"`
	BuildingType            string   `json:"buildingType"`
	CreatorID               int      `json:"creatorId"`
	EventType               string   `json:"eventType"`
	ItemAfter               int      `json:"itemAfter"`
	ItemBefore              int      `json:"itemBefore"`
	ItemID                  int      `json:"itemId"`
	KillerID                int      `json:"killerId"`
	LaneType                string   `json:"laneType"`
	LevelUpType             string   `json:"levelUpType"`
	MonsterType             string   `json:"monsterType"`
	ParticipantID           int      `json:"participantId"`
	PointCaptured           string   `json:"pointCaptured"`
	Position                Position `json:"position"`
	SkillSlot               int      `json:"skillSlot"`
	TeamID                  int      `json:"teamId"`
	Timestamp               int64    `json:"timestamp"`
	TowerType               string   `json:"towerType"`
	VictimId                int      `json:"victimId"`
	WardType                string   `json:"wardType"`
}
```


#### type FeaturedBannedChampion

```go
type FeaturedBannedChampion struct {
	ChampionId int64 `json:"championId"`
	PickTurn   int   `json:"pickTurn"`
	TeamId     int64 `json:"teamId"`
}
```


#### type FeaturedGameInfo

```go
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
```


#### type FeaturedGames

```go
type FeaturedGames struct {
	ClientRefreshInterval int64              `json:"clientRefreshInterval"`
	GameList              []FeaturedGameInfo `json:"gameList"`
}
```


#### type FeaturedParticipant

```go
type FeaturedParticipant struct {
	Bot           bool   `json:"bot"`
	ChampionId    int64  `json:"championId"`
	ProfileIconId int64  `json:"profileIconId"`
	Spell1Id      int64  `json:"spell1Id"`
	Spell2Id      int64  `json:"spell2Id"`
	SummonerName  string `json:"summonerName"`
	TeamId        int64  `json:"teamId"`
}
```


#### type Frame

```go
type Frame struct {
	Events            []Event                     `json:"events"`
	ParticipantFrames map[string]ParticipantFrame `json:"participantsFrames"`
	Timestamp         int64                       `json:"timestamp"`
}
```


#### type Game

```go
type Game struct {
	ChampionID    int      `json:"championId"`
	CreateDate    int64    `json:"createDate"`
	FellowPlayers []Player `json:"fellowPlayers"`
	GameID        int64    `json:"gameId"`
	GameMode      string   `json:"gameMode"`
	GameType      string   `json:"gameType"`
	Invalid       bool     `json:"invalid"`
	IPEarned      int      `json:"ipEarned"`
	Level         int      `json:"level"`
	MapID         int      `json:"mapId"`
	Spell1        int      `json:"spell1"`
	Spell2        int      `json:"spell2"`
	Statistics    GameStat `json:"stats"`
	SubType       string   `json:"subType"`
	TeamID        int      `json:"teamId"`
}
```

Game represents a League of Legends game

#### type GameStat

```go
type GameStat struct {
	Assists                         int  `json:"assists"`
	BarracksKilled                  int  `json:"barracksKilled"`
	ChampionsKilled                 int  `json:"championsKilled"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	ConsumablesPurchased            int  `json:"consumablesPurchased"`
	DamageDealtPlayer               int  `json:"damageDealtPlayer"`
	DoubleKills                     int  `json:"doubleKills"`
	FirstBlood                      int  `json:"firstBlood"`
	Gold                            int  `json:"gold"`
	GoldEarned                      int  `json:"goldEarned"`
	GoldSpent                       int  `json:"goldSpent"`
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Item6                           int  `json:"item6"`
	ItemsPurchased                  int  `json:"itemsPurchased"`
	KillingSprees                   int  `json:"killingSprees"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	LegendaryItemsCreated           int  `json:"legendaryItemsCreated"`
	Level                           int  `json:"level"`
	MagicDamageDealtPlayer          int  `json:"magicDamageDealtPlayer"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	MagicDamageTaken                int  `json:"magicDamageTaken"`
	MinionsDenied                   int  `json:"minionsDenied"`
	MinionsKilled                   int  `json:"minionsKilled"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledYourJungle  int  `json:"neutralMinionsKilledYourJungle"`
	NexusKilled                     bool `json:"nexusKilled"`
	NodeCapture                     int  `json:"nodeCapture"`
	NodeCaptureAssist               int  `json:"nodeCaptureAssist"`
	NodeNeutralize                  int  `json:"nodeNeutralize"`
	NodeNeutralizeAssist            int  `json:"nodeNeutralizeAssist"`
	NumDeaths                       int  `json:"numDeaths"`
	NumItemsBought                  int  `json:"numItemsBought"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	PentaKills                      int  `json:"pentaKills"`
	PhysicalDamageDealtPlayer       int  `json:"physicalDamageDealtPlayer"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	PlayerPosition                  int  `json:"playerPosition"`
	PlayerRole                      int  `json:"playerRole"`
	QuadraKills                     int  `json:"quadraKills"`
	SightWardsBought                int  `json:"sightWardsBought"`
	Spell1Cast                      int  `json:"spell1Cast"`
	Spell2Cast                      int  `json:"spell2Cast"`
	Spell3Cast                      int  `json:"spell3Cast"`
	Spell4Cast                      int  `json:"spell4Cast"`
	SummonSpell1Cast                int  `json:"summonSpell1Cast"`
	SummonSpell2Cast                int  `json:"summonSpell2Cast"`
	SuperMonsterKilled              int  `json:"superMonsterKilled"`
	Team                            int  `json:"team"`
	TeamObjective                   int  `json:"teamObjective"`
	TimePlayed                      int  `json:"timePlayed"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	TripleKills                     int  `json:"tripleKills"`
	TrueDamageDealtPlayer           int  `json:"trueDamageDealtPlayer"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	TurretsKilled                   int  `json:"turretsKilled"`
	UnrealKills                     int  `json:"unrealKills"`
	VictoryPointTotal               int  `json:"victoryPointTotal"`
	VisionWardsBought               int  `json:"visionWardsBought"`
	WardKilled                      int  `json:"wardKilled"`
	WardPlaced                      int  `json:"wardPlaced"`
	Win                             bool `json:"win"`
}
```

GameStat represents a stat for the assosiated Game

#### type League

```go
type League struct {
	Entries       []LeagueItem `json:"entries"`
	Name          string       `json:"name"`
	ParticipantId string       `json:"participantId"`
	Queue         string       `json:"queue"`
	Tier          string       `json:"tier"`
}
```

League represents a League of Legends league

#### type LeagueItem

```go
type LeagueItem struct {
	Division         string     `json:"division"`
	IsFreshBlood     bool       `json:"isFreshBlood"`
	IsHotStreak      bool       `json:"isHotStreak"`
	IsInactive       bool       `json:"isInactive"`
	IsVeteran        bool       `json:"isVeteran"`
	LeaguePoints     int        `json:"leaguePoints"`
	Losses           int        `json:"losses"`
	MiniSeries       MiniSeries `json:"miniSeries"`
	PlayerOrTeamID   string     `json:"playerOrTeamId"`
	PlayerOrTeamName string     `json:"playerOrTeamName"`
	Wins             int        `json:"wins"`
}
```

LeagueItem is an entry in a League. It represents a player or team

#### type Mastery

```go
type Mastery struct {
	MasteryId int64 `json:"masteryId"`
	Rank      int64 `json:"rank"`
}
```

Mastery being used

#### type MasteryDto

```go
type MasteryDto struct {
	ID   int `json:"id"`
	Rank int `json:"rank"`
}
```

Mastery located inside a page

#### type MasteryPage

```go
type MasteryPage struct {
	Current   bool         `json:"current"`
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Masteries []MasteryDto `json:"masteries"`
}
```

MasteryPage represents a League of Legends mastery page

#### type MatchDetail

```go
type MatchDetail struct {
	MapID                 int                   `json:"mapId"`
	MatchCreation         int64                 `json:"matchCreation"`
	MatchDuration         int64                 `json:"matchDuration"`
	MatchID               int64                 `json:"matchId"`
	MatchMode             string                `json:"matchMode"`
	MatchType             string                `json:"matchType"`
	MatchVersion          string                `json:"matchVersion"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json:"participants"`
	PlatformId            string                `json:"platformId"`
	QueueType             string                `json:"queueType"`
	Region                string                `json:"region"`
	Season                string                `json:"season"`
	Teams                 []TeamGameDetails     `json:"teams"`
	Timeline              Timeline              `json:"timeline"`
}
```


#### type MatchHistorySummary

```go
type MatchHistorySummary struct {
	Assists           int    `json:"assists"`
	Date              int64  `json:"date"`
	Deaths            int    `json:"deaths"`
	GameID            int64  `json:"gameId"`
	GameMode          string `json:"gameMode"`
	Invalid           bool   `json:"invalid"`
	Kills             int    `json:"kills"`
	MapID             int    `json:"mapId"`
	OpposingTeamKills int    `json:"opposingTeamKills"`
	OpposingTeamName  string `json:"opposingTeamName"`
	Win               bool   `json:"win"`
}
```

MatchHistorySummary is a summary of a matches played by a Team

#### type MatchPlayer

```go
type MatchPlayer struct {
	MatchHistoryURI string `json:"matchHistoryUri"`
	ProfileIcon     int    `json:"profileIcon"`
	SummonerId      int64  `json:"summonerId"`
	SummonerName    string `json:"summonerName"`
}
```


#### type MatchSummary

```go
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
```


#### type MiniSeries

```go
type MiniSeries struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}
```

MiniSeries shows if a player is in their Series and how far they are within it

#### type Observer

```go
type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}
```


#### type Participant

```go
type Participant struct {
	ChampionID                int                 `json:"championId"`
	HighestAchievedSeasonTier string              `json:"highestAchievedSeasonTier"`
	Masteries                 []Mastery           `json:"masteries"`
	ParticipantID             int                 `json:"participantId"`
	Runes                     []Rune              `json:"runes"`
	Spell1ID                  int                 `json:"spell1Id"`
	Spell2ID                  int                 `json:"spell2Id"`
	Stats                     ParticipantStats    `json:"stats"`
	TeamID                    int                 `json:"teamId"`
	Timeline                  ParticipantTimeline `json:"timeline"`
}
```


#### type ParticipantFrame

```go
type ParticipantFrame struct {
	CurrentGold         int      `json:"currentGold"`
	DominionScore       int      `json:"dominionScore"`
	JungleMinionsKilled int      `json:"jungleMinionsKilled"`
	Level               int      `json:"level"`
	MinionsKilled       int      `json:"minionsKilled"`
	ParticipantID       int      `json:"participantId"`
	Position            Position `json:"position"`
	TeamScore           int      `json:"teamScore"`
	TotalGold           int      `json:"totalGold"`
	Xp                  int      `json:"xp"`
}
```


#### type ParticipantIdentity

```go
type ParticipantIdentity struct {
	ParticipantId int         `json:"participantId"`
	Player        MatchPlayer `json:"player"`
}
```


#### type ParticipantStats

```go
type ParticipantStats struct {
	Assists                         int64 `json:"assists"`
	ChampLevel                      int64 `json:"champLevel"`
	CombatPlayerScore               int64 `json:"combatPlayerScore"`
	Deaths                          int64 `json:"deaths"`
	DoubleKills                     int64 `json:"doubleKills"`
	FirstBloodAssist                bool  `json:"firstBloodAssist"`
	FirstBloodKill                  bool  `json:"firstBloodKill"`
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`
	FirstInhibitorKill              bool  `json:"firstInhibitorKill"`
	FirstTowerAssist                bool  `json:"firstTowerAssist"`
	FirstTowerKill                  bool  `json:"firstTowerKill"`
	GoldEarned                      int64 `json:"goldEarned"`
	GoldSpent                       int64 `json:"goldSpent"`
	InhibitorKills                  int64 `json:"inhibitorKills"`
	Item0                           int64 `json:"item0"`
	Item1                           int64 `json:"item1"`
	Item2                           int64 `json:"item2"`
	Item3                           int64 `json:"item3"`
	Item4                           int64 `json:"item4"`
	Item5                           int64 `json:"item5"`
	Item6                           int64 `json:"item6"`
	KillingSprees                   int64 `json:"killingSprees"`
	Kills                           int64 `json:"kills"`
	LargestCriticalStrike           int64 `json:"largestCriticalStrike"`
	LargestKillingSpree             int64 `json:"largestKillingSpree"`
	LargestMultiKill                int64 `json:"largestMultiKill"`
	MagicDamageDealt                int64 `json:"magicDamageDealt"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`
	MagicDamageTaken                int64 `json:"magicDamageTaken"`
	MinionsKilled                   int64 `json:"minionsKilled"`
	NeutralMinionsKilled            int64 `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int64 `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledTeamJungle  int64 `json:"neutralMinionsKilledTeamJungle"`
	NodeCapture                     int64 `json:"nodeCapture"`
	NodeCaptureAssist               int64 `json:"nodeCaptureAssist"`
	NodeNeutralize                  int64 `json:"nodeNeutralize"`
	NodeNeutralizeAssist            int64 `json:"nodeNeutralizeAssist"`
	ObjectivePlayerScore            int64 `json:"objectivePlayerScore"`
	PentaKills                      int64 `json:"pentaKills"`
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
	QuadraKills                     int64 `json:"quadraKills"`
	SightWardsBoughtInGame          int64 `json:"sightWardsBoughtInGame"`
	TeamObjective                   int64 `json:"teamObjective"`
	TotalDamageDealt                int64 `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                int64 `json:"totalDamageTaken"`
	TotalHeal                       int64 `json:"totalHeal"`
	TotalPlayerScore                int64 `json:"totalPlayerScore"`
	TotalScoreRank                  int64 `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int64 `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int64 `json:"totalUnitsHealed"`
	TowerKills                      int64 `json:"towerKills"`
	TripleKills                     int64 `json:"tripleKills"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken"`
	UnrealKills                     int64 `json:"unrealKills"`
	VisionWardsBoughtInGame         int64 `json:"visionWardsBoughtInGame"`
	WardsKilled                     int64 `json:"wardsKilled"`
	WardsPlaced                     int64 `json:"wardsPlaced"`
	Winner                          bool  `json:"winner"`
}
```


#### type ParticipantTimeline

```go
type ParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts ParticipantTimelineData `json:"ancientGolemAssistsPerMinCounts"`
	AncientGolemKillsPerMinCounts   ParticipantTimelineData `json:"ancientGolemKillsPerMinCounts"`
	AssistedLaneDeathsPerMinDeltas  ParticipantTimelineData `json:"assistedLaneDeathsPerMinDeltas"`
	AssistedLaneKillsPerMinDeltas   ParticipantTimelineData `json:"assistedLaneKillsPerMinDeltas"`
	BaronAssistsPerMinCounts        ParticipantTimelineData `json:"baronAssistsPerMinCounts"`
	BaronKillsPerMinCounts          ParticipantTimelineData `json:"baronKillsPerMinCounts"`
	CreepsPerMinDeltas              ParticipantTimelineData `json:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas              ParticipantTimelineData `json:"csDiffPerMinDeltas"`
	DamageTakenDiffPerMinDeltas     ParticipantTimelineData `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas         ParticipantTimelineData `json:"damageTakenPerMinDeltas"`
	DragonAssistsPerMinCounts       ParticipantTimelineData `json:"dragonAssistsPerMinCounts"`
	DragonKillsPerMinCounts         ParticipantTimelineData `json:"dragonKillsPerMinCounts"`
	ElderLizardAssistsPerMinCounts  ParticipantTimelineData `json:"elderLizardAssistsPerMinCounts"`
	ElderLizardKillsPerMinCounts    ParticipantTimelineData `json:"elderLizardKillsPerMinCounts"`
	GoldPerMinDeltas                ParticipantTimelineData `json:"goldPerMinDeltas"`
	InhibitorAssistsPerMinCounts    ParticipantTimelineData `json:"inhibitorAssistsPerMinCounts"`
	InhibitorKillsPerMinCounts      ParticipantTimelineData `json:"inhibitorKillsPerMinCounts"`
	Lane                            string                  `json:"lane"`
	Role                            string                  `json:"role"`
	TowerAssistsPerMinCounts        ParticipantTimelineData `json:"towerAssistsPerMinCounts"`
	TowerKillsPerMinCounts          ParticipantTimelineData `json:"towerKillsPerMinCounts"`
	TowerKillsPerMinDeltas          ParticipantTimelineData `json:"towerKillsPerMinDeltas"`
	VilemawAssistsPerMinCounts      ParticipantTimelineData `json:"vilemawAssistsPerMinCounts"`
	VilemawKillsPerMinCounts        ParticipantTimelineData `json:"vilemawKillsPerMinCounts"`
	WardsPerMinDeltas               ParticipantTimelineData `json:"wardsPerMinDeltas"`
	XpDiffPerMinDeltas              ParticipantTimelineData `json:"xpDiffPerMinDeltas"`
	XpPerMinDeltas                  ParticipantTimelineData `json:"xpPerMinDeltas"`
}
```


#### type ParticipantTimelineData

```go
type ParticipantTimelineData struct {
	TenToTwenty    float64 `json:"tenToTwenty"`
	ThirtyToEnd    float64 `json:"thirtyToEnd"`
	TwentyToThirty float64 `json:"twentyToThirty"`
	ZeroToTen      float64 `json:"zeroToTen"`
}
```


#### type Player

```go
type Player struct {
	ChampionID int   `json:"championId"`
	SummonerID int64 `json:"summonerId"`
	TeamID     int   `json:"teamId"`
}
```

Player represents a League of Legends player that was in the requested game

#### type PlayerHistory

```go
type PlayerHistory struct {
	Matches []MatchSummary `json:"matches"`
}
```


#### type PlayerStatsSummary

```go
type PlayerStatsSummary struct {
	AggregatedStats       AggregatedStats `json:"aggregatedStats"`
	Losses                int             `json:"losses"`
	ModifyDate            int64           `json:"modifyDate"`
	PlayerStatSummaryType string          `json:"playerStatSummaryType"`
	Wins                  int             `json:"wins"`
}
```

PlayerStatsSummary represents a summary of a League of Legends player's game
stats

#### type Position

```go
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
```


#### type RankedStats

```go
type RankedStats struct {
	Champions  []ChampionStats `json:"champions"`
	ModifyDate int64           `json:"modifyDate"`
	SummonerID int64           `json:"summonerId"`
}
```

RankedStats represents a League of Legends player's statistics for ranked play

#### type RiotAPI

```go
type RiotAPI struct {
	APIkey string
	Region string
}
```


#### func  Get

```go
func Get(key string, region string, smallLimit int, longLimit int) *RiotAPI
```

#### func (*RiotAPI) ChampionByID

```go
func (api *RiotAPI) ChampionByID(championID string) (champion Champion, err error)
```

#### func (*RiotAPI) ChampionList

```go
func (api *RiotAPI) ChampionList(freetoplay string) (champions championList, err error)
```

#### func (*RiotAPI) CurrentGameBySummonerID

```go
func (api *RiotAPI) CurrentGameBySummonerID(plateform string, summonerID string) (currentGame CurrentGameInfo, err error)
```

#### func (*RiotAPI) FeaturedGames

```go
func (api *RiotAPI) FeaturedGames(region string) (featuredGames FeaturedGames, err error)
```

#### func (*RiotAPI) LeagueByChallenger

```go
func (api *RiotAPI) LeagueByChallenger(queue string) (league League, err error)
```

#### func (*RiotAPI) LeagueByMaster

```go
func (api *RiotAPI) LeagueByMaster(queue string) (league League, err error)
```

#### func (*RiotAPI) LeagueBySummoner

```go
func (api *RiotAPI) LeagueBySummoner(summonerID string) (leagues map[string][]League, err error)
```

#### func (*RiotAPI) LeagueByTeam

```go
func (api *RiotAPI) LeagueByTeam(teamID string) (leagues map[string][]League, err error)
```

#### func (*RiotAPI) LeagueEntryBySummoner

```go
func (api *RiotAPI) LeagueEntryBySummoner(summonerID string) (leagues map[string][]League, err error)
```

#### func (*RiotAPI) LeagueEntryByTeam

```go
func (api *RiotAPI) LeagueEntryByTeam(teamID string) (leagues map[string][]League, err error)
```

#### func (*RiotAPI) MasteriesByID

```go
func (api *RiotAPI) MasteriesByID(summonerID string) (masteries map[string]masteryBook, err error)
```

#### func (*RiotAPI) MatchByMatchID

```go
func (api *RiotAPI) MatchByMatchID(matchID string, includeTimeline string) (match MatchDetail, err error)
```

#### func (*RiotAPI) MatchHistoryBySummonerID

```go
func (api *RiotAPI) MatchHistoryBySummonerID(options ...interface{}) (playerHistory PlayerHistory, err error)
```

#### func (*RiotAPI) RankedStatsBySummoner

```go
func (api *RiotAPI) RankedStatsBySummoner(summonerID string, season string) (stats RankedStats, err error)
```

#### func (*RiotAPI) RecentGameBySummoner

```go
func (api *RiotAPI) RecentGameBySummoner(summonerID string) (games gamesList, err error)
```

#### func (*RiotAPI) RunesByID

```go
func (api *RiotAPI) RunesByID(summonerID string) (runes map[string]runeBook, err error)
```

#### func (*RiotAPI) SetLongRateLimit

```go
func (api *RiotAPI) SetLongRateLimit(numrequests int, pertime time.Duration)
```
SetLongRateLimit allows a custom rate limit to be set. For at the time of this
writing the default for a development API key is 500 requests every 10 minutes

#### func (*RiotAPI) SetSmallRateLimit

```go
func (api *RiotAPI) SetSmallRateLimit(numrequests int, pertime time.Duration)
```
SetSmallRateLimit allows a custom rate limit to be set. For at the time of this
writing the default for a development API key is 10 requests every 10 seconds

#### func (*RiotAPI) StatSummariesBySummoner

```go
func (api *RiotAPI) StatSummariesBySummoner(summonerID string, season string) (stats playerStatsSummaryList, err error)
```

#### func (*RiotAPI) SummonerByID

```go
func (api *RiotAPI) SummonerByID(summonerID string) (summoners map[string]Summoner, err error)
```

#### func (*RiotAPI) SummonerByName

```go
func (api *RiotAPI) SummonerByName(names string) (summoners map[string]Summoner, err error)
```

#### func (*RiotAPI) TeamBySummonerID

```go
func (api *RiotAPI) TeamBySummonerID(summonerID string) (teams map[string][]Team, err error)
```

#### func (*RiotAPI) TeamByTeamID

```go
func (api *RiotAPI) TeamByTeamID(teamID string) (teams map[string]Team, err error)
```

#### type RiotError

```go
type RiotError struct {
	StatusCode int
}
```

RiotError contains the http status code of the erro

#### func (RiotError) Error

```go
func (err RiotError) Error() string
```
Error prints the error message for a RiotError

#### type Roster

```go
type Roster struct {
	MemberList []TeamMemberInfo `json:"memberList"`
	OwnerID    int64            `json:"ownerId"`
}
```

Roster represents the roster of a League of Legends ranked team

#### type Rune

```go
type Rune struct {
	Rank   int64 `json:"rank"`
	RuneId int64 `json:"runeId"`
}
```


#### type RuneCount

```go
type RuneCount struct {
	Count  int   `json:"count"`
	RuneId int64 `json:"runeId"`
}
```

Mastery alreay exist

#### type RunePage

```go
type RunePage struct {
	Current bool       `json:"current"`
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Slots   []RuneSlot `json:"slots"`
}
```

RunePage is a League of Legends rune page

#### type RuneSlot

```go
type RuneSlot struct {
	RuneId     int `json:"runeId"`
	RuneSlotID int `json:"runeSlotId"`
}
```

RuneSlot is a slot for a Rune to go on a RunePage

#### type Summoner

```go
type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	SummonerLevel int    `json:"summonerLevel"`
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	RevisionDate  int64  `json:"revisionDate"`
}
```

Summoner is a player of League of Legends

#### type Team

```go
type Team struct {
	CreateDate                    int64                 `json:"createDate"`
	FullID                        string                `json:"fullId"`
	LastGameDate                  int64                 `json:"lastGameDate"`
	LastJoinDate                  int64                 `json:"lastJoinDate"`
	LastJoinedRankedTeamQueueDate int64                 `json:"lastJoinedRankedTeamQueueDate"`
	MatchHistory                  []MatchHistorySummary `json:"matchHistory"`
	ModifyDate                    int64                 `json:"modifyDate"`
	Name                          string                `json:"name"`
	Roster                        Roster                `json:"roster"`
	SecondLastJoinDate            int64                 `json:"secondLastJoinDate"`
	Status                        string                `json:"status"`
	Tag                           string                `json:"tag"`
	TeamStatDetails               []TeamStatDetail      `json:"teamStatDetails"`
	ThirdJoinDate                 int64                 `json:"thirdLastJoinDate"`
}
```

Team is a League of Legends Ranked Team

#### type TeamGameDetails

```go
type TeamGameDetails struct {
	Bans                 []BannedChampion `json:"bans"`
	BaronKills           int              `json:"baronKills"`
	DominionVictoryScore int64            `json:"dominionVictoryScore"`
	DragonKills          int              `json:"dragonKills"`
	FirstBaron           bool             `json:"firstBaron"`
	FirstBlood           bool             `json:"firstBlood"`
	FirstDragon          bool             `json:"firstDragon"`
	FirstInhibitor       bool             `json:"firstInhibitor"`
	FirstTower           bool             `json:"firstTower"`
	InhibitorKills       int              `json:"inhibitorKills"`
	TeamID               int              `json:"teamId"`
	TowerKills           int              `json:"towerKills"`
	VilemawKills         int              `json:"vilemawKills"`
	Winner               bool             `json:"winner"`
}
```


#### type TeamMemberInfo

```go
type TeamMemberInfo struct {
	InviteDate int64  `json:"inviteDate"`
	JoinDate   int64  `json:"joinDate"`
	PlayerID   int64  `json:"playerId"`
	Status     string `json:"status"`
}
```

TeamMemberInfo is the individual information for a player on a ranked team

#### type TeamStatDetail

```go
type TeamStatDetail struct {
	AverageGamesPlayed int    `json:"averageGamesPlayed"`
	Losses             int    `json:"losses"`
	TeamStatType       string `json:"teamStatType"`
	Wins               int    `json:"wins"`
}
```

TeamStatDetail is the statistics for a ranked team

#### type Timeline

```go
type Timeline struct {
	FrameInterval int64   `json:"frameInterval"`
	Frames        []Frame `json:"frames"`
}
```
