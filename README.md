# goriot
go tool for riot api

install :
go get "github.com/kindermoumoute/goriot"

Using :
import "github.com/kindermoumoute/goriot"

Example :
myapi := goriot.Get(APIkey, goriot.EUW, 10, 500)


<!--
	Copyright 2009 The Go Authors. All rights reserved.
	Use of this source code is governed by a BSD-style
	license that can be found in the LICENSE file.
-->
<!--
	Note: Static (i.e., not template-generated) href and id
	attributes start with "pkg-" to make it impossible for
	them to conflict with generated attributes (some of which
	correspond to Go identifiers).
-->

	<script type='text/javascript'>
	document.ANALYSIS_DATA = ;
	document.CALLGRAPH = ;
	</script>

	
		
		<div id="short-nav">
			<dl>
			<dd><code>import "github.com/kindermoumoute/goriot"</code></dd>
			</dl>
			<dl>
			<dd><a href="#pkg-overview" class="overviewLink">Overview</a></dd>
			<dd><a href="#pkg-index" class="indexLink">Index</a></dd>
			
			
			</dl>
		</div>
		<!-- The package's Name is printed as title by the top-level template -->
		<div id="pkg-overview" class="toggleVisible">
			<div class="collapsed">
				<h2 class="toggleButton" title="Click to show Overview section">Overview ▹</h2>
			</div>
			<div class="expanded">
				<h2 class="toggleButton" title="Click to hide Overview section">Overview ▾</h2>
				
			</div>
		</div>
		

		<div id="pkg-index" class="toggleVisible">
		<div class="collapsed">
			<h2 class="toggleButton" title="Click to show Index section">Index ▹</h2>
		</div>
		<div class="expanded">
			<h2 class="toggleButton" title="Click to hide Index section">Index ▾</h2>

		<!-- Table of contents for API; must be named manual-nav to turn off auto nav. -->
			<div id="manual-nav">
			<dl>
			
				<dd><a href="#pkg-constants">Constants</a></dd>
			
			
			
			
				
				<dd><a href="#AggregatedStats">type AggregatedStats</a></dd>
				
				
			
				
				<dd><a href="#BannedChampion">type BannedChampion</a></dd>
				
				
			
				
				<dd><a href="#Champion">type Champion</a></dd>
				
				
			
				
				<dd><a href="#ChampionStats">type ChampionStats</a></dd>
				
				
			
				
				<dd><a href="#CurrentBannedChampion">type CurrentBannedChampion</a></dd>
				
				
			
				
				<dd><a href="#CurrentGameInfo">type CurrentGameInfo</a></dd>
				
				
			
				
				<dd><a href="#CurrentGameParticipant">type CurrentGameParticipant</a></dd>
				
				
			
				
				<dd><a href="#Event">type Event</a></dd>
				
				
			
				
				<dd><a href="#FeaturedBannedChampion">type FeaturedBannedChampion</a></dd>
				
				
			
				
				<dd><a href="#FeaturedGameInfo">type FeaturedGameInfo</a></dd>
				
				
			
				
				<dd><a href="#FeaturedGames">type FeaturedGames</a></dd>
				
				
			
				
				<dd><a href="#FeaturedParticipant">type FeaturedParticipant</a></dd>
				
				
			
				
				<dd><a href="#Frame">type Frame</a></dd>
				
				
			
				
				<dd><a href="#Game">type Game</a></dd>
				
				
			
				
				<dd><a href="#GameStat">type GameStat</a></dd>
				
				
			
				
				<dd><a href="#League">type League</a></dd>
				
				
			
				
				<dd><a href="#LeagueItem">type LeagueItem</a></dd>
				
				
			
				
				<dd><a href="#Mastery">type Mastery</a></dd>
				
				
			
				
				<dd><a href="#MasteryDto">type MasteryDto</a></dd>
				
				
			
				
				<dd><a href="#MasteryPage">type MasteryPage</a></dd>
				
				
			
				
				<dd><a href="#MatchDetail">type MatchDetail</a></dd>
				
				
			
				
				<dd><a href="#MatchHistorySummary">type MatchHistorySummary</a></dd>
				
				
			
				
				<dd><a href="#MatchPlayer">type MatchPlayer</a></dd>
				
				
			
				
				<dd><a href="#MatchSummary">type MatchSummary</a></dd>
				
				
			
				
				<dd><a href="#MiniSeries">type MiniSeries</a></dd>
				
				
			
				
				<dd><a href="#Observer">type Observer</a></dd>
				
				
			
				
				<dd><a href="#Participant">type Participant</a></dd>
				
				
			
				
				<dd><a href="#ParticipantFrame">type ParticipantFrame</a></dd>
				
				
			
				
				<dd><a href="#ParticipantIdentity">type ParticipantIdentity</a></dd>
				
				
			
				
				<dd><a href="#ParticipantStats">type ParticipantStats</a></dd>
				
				
			
				
				<dd><a href="#ParticipantTimeline">type ParticipantTimeline</a></dd>
				
				
			
				
				<dd><a href="#ParticipantTimelineData">type ParticipantTimelineData</a></dd>
				
				
			
				
				<dd><a href="#Player">type Player</a></dd>
				
				
			
				
				<dd><a href="#PlayerHistory">type PlayerHistory</a></dd>
				
				
			
				
				<dd><a href="#PlayerStatsSummary">type PlayerStatsSummary</a></dd>
				
				
			
				
				<dd><a href="#Position">type Position</a></dd>
				
				
			
				
				<dd><a href="#RankedStats">type RankedStats</a></dd>
				
				
			
				
				<dd><a href="#RiotAPI">type RiotAPI</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Get">func Get(key string, region string, smallLimit int, longLimit int) *RiotAPI</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.ChampionByID">func (api *RiotAPI) ChampionByID(championID string) (champion Champion, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.ChampionList">func (api *RiotAPI) ChampionList(freetoplay string) (champions championList, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.CurrentGameBySummonerID">func (api *RiotAPI) CurrentGameBySummonerID(plateform string, summonerID string) (currentGame CurrentGameInfo, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.FeaturedGames">func (api *RiotAPI) FeaturedGames(region string) (featuredGames FeaturedGames, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.LeagueByChallenger">func (api *RiotAPI) LeagueByChallenger(queue string) (league League, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.LeagueByMaster">func (api *RiotAPI) LeagueByMaster(queue string) (league League, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.LeagueBySummoner">func (api *RiotAPI) LeagueBySummoner(summonerID string) (leagues map[string][]League, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.LeagueByTeam">func (api *RiotAPI) LeagueByTeam(teamID string) (leagues map[string][]League, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.LeagueEntryBySummoner">func (api *RiotAPI) LeagueEntryBySummoner(summonerID string) (leagues map[string][]League, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.LeagueEntryByTeam">func (api *RiotAPI) LeagueEntryByTeam(teamID string) (leagues map[string][]League, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.MasteriesByID">func (api *RiotAPI) MasteriesByID(summonerID string) (masteries map[string]masteryBook, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.MatchByMatchID">func (api *RiotAPI) MatchByMatchID(matchID string, includeTimeline string) (match MatchDetail, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.MatchHistoryBySummonerID">func (api *RiotAPI) MatchHistoryBySummonerID(options ...interface{}) (playerHistory PlayerHistory, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.RankedStatsBySummoner">func (api *RiotAPI) RankedStatsBySummoner(summonerID string, season string) (stats RankedStats, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.RecentGameBySummoner">func (api *RiotAPI) RecentGameBySummoner(summonerID string) (games gamesList, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.RunesByID">func (api *RiotAPI) RunesByID(summonerID string) (runes map[string]runeBook, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.SetLongRateLimit">func (api *RiotAPI) SetLongRateLimit(numrequests int, pertime time.Duration)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.SetSmallRateLimit">func (api *RiotAPI) SetSmallRateLimit(numrequests int, pertime time.Duration)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.StatSummariesBySummoner">func (api *RiotAPI) StatSummariesBySummoner(summonerID string, season string) (stats playerStatsSummaryList, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.SummonerByID">func (api *RiotAPI) SummonerByID(summonerID string) (summoners map[string]Summoner, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.SummonerByName">func (api *RiotAPI) SummonerByName(names string) (summoners map[string]Summoner, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.TeamBySummonerID">func (api *RiotAPI) TeamBySummonerID(summonerID string) (teams map[string][]Team, err error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotAPI.TeamByTeamID">func (api *RiotAPI) TeamByTeamID(teamID string) (teams map[string]Team, err error)</a></dd>
				
			
				
				<dd><a href="#RiotError">type RiotError</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#RiotError.Error">func (err RiotError) Error() string</a></dd>
				
			
				
				<dd><a href="#Roster">type Roster</a></dd>
				
				
			
				
				<dd><a href="#Rune">type Rune</a></dd>
				
				
			
				
				<dd><a href="#RuneCount">type RuneCount</a></dd>
				
				
			
				
				<dd><a href="#RunePage">type RunePage</a></dd>
				
				
			
				
				<dd><a href="#RuneSlot">type RuneSlot</a></dd>
				
				
			
				
				<dd><a href="#Summoner">type Summoner</a></dd>
				
				
			
				
				<dd><a href="#Team">type Team</a></dd>
				
				
			
				
				<dd><a href="#TeamGameDetails">type TeamGameDetails</a></dd>
				
				
			
				
				<dd><a href="#TeamMemberInfo">type TeamMemberInfo</a></dd>
				
				
			
				
				<dd><a href="#TeamStatDetail">type TeamStatDetail</a></dd>
				
				
			
				
				<dd><a href="#Timeline">type Timeline</a></dd>
				
				
			
			
			</dl>
			</div><!-- #manual-nav -->

		

		
			<h4>Package files</h4>
			<p>
			<span style="font-size:90%">
			
				<a href="/src/github.com/kindermoumoute/goriot/champion.go">champion.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/curentgame.go">curentgame.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/featuredgame.go">featuredgame.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/game.go">game.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/goriot.go">goriot.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/league.go">league.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/match.go">match.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/matchhistory.go">matchhistory.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/stats.go">stats.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/summoner.go">summoner.go</a>
			
				<a href="/src/github.com/kindermoumoute/goriot/team.go">team.go</a>
			
			</span>
			</p>
		
		</div><!-- .expanded -->
		</div><!-- #pkg-index -->

		<div id="pkg-callgraph" class="toggle" style="display: none">
		<div class="collapsed">
			<h2 class="toggleButton" title="Click to show Internal Call Graph section">Internal call graph ▹</h2>
		</div> <!-- .expanded -->
		<div class="expanded">
			<h2 class="toggleButton" title="Click to hide Internal Call Graph section">Internal call graph ▾</h2>
			<p>
			  In the call graph viewer below, each node
			  is a function belonging to this package
			  and its children are the functions it
			  calls&mdash;perhaps dynamically.
			</p>
			<p>
			  The root nodes are the entry points of the
			  package: functions that may be called from
			  outside the package.
			  There may be non-exported or anonymous
			  functions among them if they are called
			  dynamically from another package.
			</p>
			<p>
			  Click a node to visit that function's source code.
			  From there you can visit its callers by
			  clicking its declaring <code>func</code>
			  token.
			</p>
			<p>
			  Functions may be omitted if they were
			  determined to be unreachable in the
			  particular programs or tests that were
			  analyzed.
			</p>
			<!-- Zero means show all package entry points. -->
			<ul style="margin-left: 0.5in" id="callgraph-0" class="treeview"></ul>
		</div>
		</div> <!-- #pkg-callgraph -->

		
			<h2 id="pkg-constants">Constants</h2>
			
				<pre>const (
    <span class="comment">//BaseURL is the base of the url used by Riot&#39;s API service</span>
    <span id="BaseURL">BaseURL</span> = &#34;api.pvp.net&#34;

    <span class="comment">//BR represents the string for the Brazilian League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="BR">BR</span> = &#34;br&#34;
    <span class="comment">//EUNE represents the string for the North Eastern European League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="EUNE">EUNE</span> = &#34;eune&#34;
    <span class="comment">//EUW represents the string for the West European League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="EUW">EUW</span> = &#34;euw&#34;
    <span class="comment">//KR represents the string for the Korean League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="KR">KR</span> = &#34;kr&#34;
    <span class="comment">//LAN represents the string for the Latin America North League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="LAN">LAN</span> = &#34;lan&#34;
    <span class="comment">//LAS represents the string for the Latin America South League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="LAS">LAS</span> = &#34;las&#34;
    <span class="comment">//NA represents the string for the North American League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="NA">NA</span> = &#34;na&#34;
    <span class="comment">//OCE represents the string for the Oceanic League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="OCE">OCE</span> = &#34;oce&#34;
    <span class="comment">//RU represents the string for the Russian League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="RU">RU</span> = &#34;ru&#34;
    <span class="comment">//TR represents the string for the Turkish League of Legends Servers,</span>
    <span class="comment">//only used as a helper to prevent typos</span>
    <span id="TR">TR</span> = &#34;tr&#34;

    <span class="comment">// same</span>
    <span id="PBE">PBE</span> = &#34;pbe&#34;

    <span class="comment">//SEASON3 is the string of &#34;SEASON3&#34;.</span>
    <span id="SEASON3">SEASON3</span> = &#34;SEASON3&#34;
    <span class="comment">//SEASON4 is the string of &#34;SEASON2014&#34;.</span>
    <span id="SEASON4">SEASON4</span> = &#34;SEASON2014&#34;
    <span class="comment">//SEASON5 is the string of &#34;SEASON2015&#34;.</span>
    <span id="SEASON5">SEASON5</span> = &#34;SEASON2015&#34;
    <span class="comment">//Ranked Solo 5s</span>
    <span id="RANKED_SOLO_5x5">RANKED_SOLO_5x5</span> = &#34;RANKED_SOLO_5x5&#34;
    <span class="comment">//Ranked Team 3s</span>
    <span id="RANKED_TEAM_3x3">RANKED_TEAM_3x3</span> = &#34;RANKED_TEAM_3x3&#34;
    <span class="comment">//Ranked Team 5s</span>
    <span id="RANKED_TEAM_5x5">RANKED_TEAM_5x5</span> = &#34;RANKED_TEAM_5x5&#34;
)</pre>
				
			
		
		
		
		
			
			
			<h2 id="AggregatedStats">type <a href="/src/target/stats.go?s=720:4242#L12">AggregatedStats</a></h2>
			<pre>type AggregatedStats struct {
    AverageAssists              <a href="/pkg/builtin/#int">int</a> `json:&#34;averageAssists&#34;`
    AverageChampionsKilled      <a href="/pkg/builtin/#int">int</a> `json:&#34;averageChampionsKilled&#34;`
    AverageCombatPlayerScore    <a href="/pkg/builtin/#int">int</a> `json:&#34;averageCombatPlayerScore&#34;`
    AverageNodeCapture          <a href="/pkg/builtin/#int">int</a> `json:&#34;averageNodeCapture&#34;`
    AverageNodeCaptureAssist    <a href="/pkg/builtin/#int">int</a> `json:&#34;averageNodeCaptureAssist&#34;`
    AverageNodeNeutralize       <a href="/pkg/builtin/#int">int</a> `json:&#34;averageNodeNeutralize&#34;`
    AverageNodeNeutralizeAssist <a href="/pkg/builtin/#int">int</a> `json:&#34;averageNodeNeutralizeAssist&#34;`
    AverageNumDeaths            <a href="/pkg/builtin/#int">int</a> `json:&#34;averageNumDeaths&#34;`
    AverageObjectivePlayerScore <a href="/pkg/builtin/#int">int</a> `json:&#34;averageObjectivePlayerScore&#34;`
    AverageTeamObjective        <a href="/pkg/builtin/#int">int</a> `json:&#34;averageTeamObjective&#34;`
    AverageTotalPlayerScore     <a href="/pkg/builtin/#int">int</a> `json:&#34;averageTotalPlayerScore&#34;`
    BotGamesPlayed              <a href="/pkg/builtin/#int">int</a> `json:&#34;botGamesPlayed&#34;`
    KillingSpree                <a href="/pkg/builtin/#int">int</a> `json:&#34;killingSpree&#34;`
    MaxAssists                  <a href="/pkg/builtin/#int">int</a> `json:&#34;maxAssists&#34;`
    MaxChampionsKilled          <a href="/pkg/builtin/#int">int</a> `json:&#34;maxChampionsKilled&#34;`
    MaxCombatPlayerScore        <a href="/pkg/builtin/#int">int</a> `json:&#34;maxCombatPlayerScore&#34;`
    MaxLargestCriticalStrike    <a href="/pkg/builtin/#int">int</a> `json:&#34;maxLargestCriticalStrike&#34;`
    MaxLargestKillingSpree      <a href="/pkg/builtin/#int">int</a> `json:&#34;maxLargestKillingSpree&#34;`
    MaxNodeCapture              <a href="/pkg/builtin/#int">int</a> `json:&#34;maxNodeCapture&#34;`
    MaxNodeCaptureAssist        <a href="/pkg/builtin/#int">int</a> `json:&#34;maxNodeCaptureAssist&#34;`
    MaxNodeNeutralize           <a href="/pkg/builtin/#int">int</a> `json:&#34;maxNodeNeutralize&#34;`
    MaxNodeNeutralizeAssist     <a href="/pkg/builtin/#int">int</a> `json:&#34;maxNodeNeutralizeAssist&#34;`
    MaxNumDeaths                <a href="/pkg/builtin/#int">int</a> `json:&#34;maxNumDeaths&#34;`
    MaxObjectivePlayerScore     <a href="/pkg/builtin/#int">int</a> `json:&#34;maxObjectivePlayerScore&#34;`
    MaxTeamObjective            <a href="/pkg/builtin/#int">int</a> `json:&#34;maxTeamObjective&#34;`
    MaxTimePlayed               <a href="/pkg/builtin/#int">int</a> `json:&#34;maxTimePlayed&#34;`
    MaxTimeSpentLiving          <a href="/pkg/builtin/#int">int</a> `json:&#34;maxTimeSpentLiving&#34;`
    MaxTotalPlayerScore         <a href="/pkg/builtin/#int">int</a> `json:&#34;maxTotalPlayerScore&#34;`
    MostChampionKillsPerSession <a href="/pkg/builtin/#int">int</a> `json:&#34;mostChampionKillsPerSession&#34;`
    MostSpellsCast              <a href="/pkg/builtin/#int">int</a> `json:&#34;mostSpellsCast&#34;`
    NormalGamesPlayed           <a href="/pkg/builtin/#int">int</a> `json:&#34;normalGamesPlayed&#34;`
    RankedPremadeGamesPlayed    <a href="/pkg/builtin/#int">int</a> `json:&#34;rankedPremadeGamesPlayed&#34;`
    RankedSoloGamesPlayed       <a href="/pkg/builtin/#int">int</a> `json:&#34;rankedSoloGamesPlayed&#34;`
    TotalAssists                <a href="/pkg/builtin/#int">int</a> `json:&#34;totalAssists&#34;`
    TotalChampionKills          <a href="/pkg/builtin/#int">int</a> `json:&#34;totalChampionKills&#34;`
    TotalDamageDealt            <a href="/pkg/builtin/#int">int</a> `json:&#34;totalDamageDealt&#34;`
    TotalDamageTaken            <a href="/pkg/builtin/#int">int</a> `json:&#34;totalDamageTaken&#34;`
    TotalDeathsPerSession       <a href="/pkg/builtin/#int">int</a> `json:&#34;totalDeathsPerSession&#34;`
    TotalDoubleKills            <a href="/pkg/builtin/#int">int</a> `json:&#34;totalDoubleKills&#34;`
    TotalFirstBlood             <a href="/pkg/builtin/#int">int</a> `json:&#34;totalFirstBlood&#34;`
    TotalGoldEarned             <a href="/pkg/builtin/#int">int</a> `json:&#34;totalGoldEarned&#34;`
    TotalHeal                   <a href="/pkg/builtin/#int">int</a> `json:&#34;totalHeal&#34;`
    TotalMagicDamageDealt       <a href="/pkg/builtin/#int">int</a> `json:&#34;totalMagicDamageDealt&#34;`
    TotalMinionKills            <a href="/pkg/builtin/#int">int</a> `json:&#34;totalMinionKills&#34;`
    TotalNeutralMinionsKilled   <a href="/pkg/builtin/#int">int</a> `json:&#34;totalNeutralMinionsKilled&#34;`
    TotalNodeCapture            <a href="/pkg/builtin/#int">int</a> `json:&#34;totalNodeCapture&#34;`
    TotalNodeNeutralize         <a href="/pkg/builtin/#int">int</a> `json:&#34;totalNodeNeutralize&#34;`
    TotalPentaKills             <a href="/pkg/builtin/#int">int</a> `json:&#34;totalPentaKills&#34;`
    TotalPhysicalDamageDealt    <a href="/pkg/builtin/#int">int</a> `json:&#34;totalPhysicalDamageDealt&#34;`
    TotalQuadraKills            <a href="/pkg/builtin/#int">int</a> `json:&#34;totalQuadraKills&#34;`
    TotalSessionsLost           <a href="/pkg/builtin/#int">int</a> `json:&#34;totalSessionsLost&#34;`
    TotalSessionsPlayed         <a href="/pkg/builtin/#int">int</a> `json:&#34;totalSessionsPlayed&#34;`
    TotalSessionsWon            <a href="/pkg/builtin/#int">int</a> `json:&#34;totalSessionsWon&#34;`
    TotalTripleKills            <a href="/pkg/builtin/#int">int</a> `json:&#34;totalTripleKills&#34;`
    TotalTurretsKilled          <a href="/pkg/builtin/#int">int</a> `json:&#34;totalTurretsKilled&#34;`
    TotalUnrealKills            <a href="/pkg/builtin/#int">int</a> `json:&#34;totalUnrealKills&#34;`
}</pre>
			<p>
AggregatedStats contain all values for a player&#39;s game stat values
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="BannedChampion">type <a href="/src/target/match.go?s=9920:10023#L169">BannedChampion</a></h2>
			<pre>type BannedChampion struct {
    ChampionId <a href="/pkg/builtin/#int">int</a> `json:&#34;championId&#34;`
    PickTurn   <a href="/pkg/builtin/#int">int</a> `json:&#34;pickTurn&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="Champion">type <a href="/src/target/champion.go?s=93:385#L1">Champion</a></h2>
			<pre>type Champion struct {
    Active            <a href="/pkg/builtin/#bool">bool</a> `json:&#34;active&#34;`
    BotEnabled        <a href="/pkg/builtin/#bool">bool</a> `json:&#34;botEnabled&#34;`
    BotMmEnabled      <a href="/pkg/builtin/#bool">bool</a> `json:&#34;botMmEnabled&#34;`
    FreeToPlay        <a href="/pkg/builtin/#bool">bool</a> `json:&#34;freeToPlay&#34;`
    ID                <a href="/pkg/builtin/#int">int</a>  `json:&#34;id&#34;`
    RankedPlayEnabled <a href="/pkg/builtin/#bool">bool</a> `json:&#34;rankedPlayEnabled&#34;`
}</pre>
			<p>
Champion represents a League of Legends champion
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="ChampionStats">type <a href="/src/target/stats.go?s=4589:4694#L79">ChampionStats</a></h2>
			<pre>type ChampionStats struct {
    ID    <a href="/pkg/builtin/#int">int</a>             `json:&#34;id&#34;`
    Stats <a href="#AggregatedStats">AggregatedStats</a> `json:&#34;stats&#34;`
}</pre>
			<p>
ChampionStats are the stats for a League of Legends player&#39;s champion in ranked
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="CurrentBannedChampion">type <a href="/src/target/curentgame.go?s=795:944#L11">CurrentBannedChampion</a></h2>
			<pre>type CurrentBannedChampion struct {
    ChampionId <a href="/pkg/builtin/#int">int</a>   `json:&#34;championId&#34;`
    PickTurn   <a href="/pkg/builtin/#int">int</a>   `json:&#34;pickTurn&#34;`
    TeamId     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;teamId&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="CurrentGameInfo">type <a href="/src/target/curentgame.go?s=41:791#L1">CurrentGameInfo</a></h2>
			<pre>type CurrentGameInfo struct {
    BannedChampions   []<a href="#CurrentBannedChampion">CurrentBannedChampion</a>  `json:&#34;bannedChampions&#34;`
    GameId            <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;gameId&#34;`
    GameLength        <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;gameLength&#34;`
    GameMode          <a href="/pkg/builtin/#string">string</a>                   `json:&#34;gameMode&#34;`
    GameQueueConfigId <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;gameQueueConfigId&#34;`
    GameStartTime     <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;gameStartTime&#34;`
    GameType          <a href="/pkg/builtin/#string">string</a>                   `json:&#34;gameType&#34;`
    MapId             <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;mapId&#34;`
    Observers         <a href="#Observer">Observer</a>                 `json:&#34;observers&#34;`
    Participants      []<a href="#CurrentGameParticipant">CurrentGameParticipant</a> `json:&#34;participants&#34;`
    PlatformId        <a href="/pkg/builtin/#string">string</a>                   `json:&#34;platformId&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="CurrentGameParticipant">type <a href="/src/target/curentgame.go?s=948:1451#L17">CurrentGameParticipant</a></h2>
			<pre>type CurrentGameParticipant struct {
    Bot           <a href="/pkg/builtin/#bool">bool</a>        `json:&#34;bot&#34;`
    ChampionId    <a href="/pkg/builtin/#int64">int64</a>       `json:&#34;championId&#34;`
    Masteries     []<a href="#Mastery">Mastery</a>   `json:&#34;masteries&#34;`
    ProfileIconId <a href="/pkg/builtin/#int64">int64</a>       `json:&#34;profileIconId&#34;`
    Runes         []<a href="#RuneCount">RuneCount</a> `json:&#34;runes&#34;`
    Spell1Id      <a href="/pkg/builtin/#int64">int64</a>       `json:&#34;spell1Id&#34;`
    Spell2Id      <a href="/pkg/builtin/#int64">int64</a>       `json:&#34;spell2Id&#34;`
    SummonerId    <a href="/pkg/builtin/#int64">int64</a>       `json:&#34;summonerId&#34;`
    SummonerName  <a href="/pkg/builtin/#string">string</a>      `json:&#34;summonerName&#34;`
    TeamId        <a href="/pkg/builtin/#int64">int64</a>       `json:&#34;teamId&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="Event">type <a href="/src/target/match.go?s=10489:11667#L187">Event</a></h2>
			<pre>type Event struct {
    AscendedType            <a href="/pkg/builtin/#string">string</a>   `json:&#34;ascendedType&#34;`
    AssistingParticipantIDs []<a href="/pkg/builtin/#int">int</a>    `json:&#34;assistingParticipantIds&#34;`
    BuildingType            <a href="/pkg/builtin/#string">string</a>   `json:&#34;buildingType&#34;`
    CreatorID               <a href="/pkg/builtin/#int">int</a>      `json:&#34;creatorId&#34;`
    EventType               <a href="/pkg/builtin/#string">string</a>   `json:&#34;eventType&#34;`
    ItemAfter               <a href="/pkg/builtin/#int">int</a>      `json:&#34;itemAfter&#34;`
    ItemBefore              <a href="/pkg/builtin/#int">int</a>      `json:&#34;itemBefore&#34;`
    ItemID                  <a href="/pkg/builtin/#int">int</a>      `json:&#34;itemId&#34;`
    KillerID                <a href="/pkg/builtin/#int">int</a>      `json:&#34;killerId&#34;`
    LaneType                <a href="/pkg/builtin/#string">string</a>   `json:&#34;laneType&#34;`
    LevelUpType             <a href="/pkg/builtin/#string">string</a>   `json:&#34;levelUpType&#34;`
    MonsterType             <a href="/pkg/builtin/#string">string</a>   `json:&#34;monsterType&#34;`
    ParticipantID           <a href="/pkg/builtin/#int">int</a>      `json:&#34;participantId&#34;`
    PointCaptured           <a href="/pkg/builtin/#string">string</a>   `json:&#34;pointCaptured&#34;`
    Position                <a href="#Position">Position</a> `json:&#34;position&#34;`
    SkillSlot               <a href="/pkg/builtin/#int">int</a>      `json:&#34;skillSlot&#34;`
    TeamID                  <a href="/pkg/builtin/#int">int</a>      `json:&#34;teamId&#34;`
    Timestamp               <a href="/pkg/builtin/#int64">int64</a>    `json:&#34;timestamp&#34;`
    TowerType               <a href="/pkg/builtin/#string">string</a>   `json:&#34;towerType&#34;`
    VictimId                <a href="/pkg/builtin/#int">int</a>      `json:&#34;victimId&#34;`
    WardType                <a href="/pkg/builtin/#string">string</a>   `json:&#34;wardType&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="FeaturedBannedChampion">type <a href="/src/target/featuredgame.go?s=965:1115#L16">FeaturedBannedChampion</a></h2>
			<pre>type FeaturedBannedChampion struct {
    ChampionId <a href="/pkg/builtin/#int64">int64</a> `json:&#34;championId&#34;`
    PickTurn   <a href="/pkg/builtin/#int">int</a>   `json:&#34;pickTurn&#34;`
    TeamId     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;teamId&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="FeaturedGameInfo">type <a href="/src/target/featuredgame.go?s=210:961#L2">FeaturedGameInfo</a></h2>
			<pre>type FeaturedGameInfo struct {
    BannedChampions   []<a href="#FeaturedBannedChampion">FeaturedBannedChampion</a> `json:&#34;bannedChampions&#34;`
    GameId            <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;gameId&#34;`
    GameLength        <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;gameLength&#34;`
    GameMode          <a href="/pkg/builtin/#string">string</a>                   `json:&#34;gameMode&#34;`
    GameQueueConfigId <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;gameQueueConfigId&#34;`
    GameStartTime     <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;gameStartTime&#34;`
    GameType          <a href="/pkg/builtin/#string">string</a>                   `json:&#34;gameType&#34;`
    MapId             <a href="/pkg/builtin/#int64">int64</a>                    `json:&#34;mapId&#34;`
    Observers         <a href="#Observer">Observer</a>                 `json:&#34;observers&#34;`
    Participants      []<a href="#FeaturedParticipant">FeaturedParticipant</a>    `json:&#34;participants&#34;`
    PlatformId        <a href="/pkg/builtin/#string">string</a>                   `json:&#34;platformId&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="FeaturedGames">type <a href="/src/target/featuredgame.go?s=41:206#L1">FeaturedGames</a></h2>
			<pre>type FeaturedGames struct {
    ClientRefreshInterval <a href="/pkg/builtin/#int64">int64</a>              `json:&#34;clientRefreshInterval&#34;`
    GameList              []<a href="#FeaturedGameInfo">FeaturedGameInfo</a> `json:&#34;gameList&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="FeaturedParticipant">type <a href="/src/target/featuredgame.go?s=1119:1446#L22">FeaturedParticipant</a></h2>
			<pre>type FeaturedParticipant struct {
    Bot           <a href="/pkg/builtin/#bool">bool</a>   `json:&#34;bot&#34;`
    ChampionId    <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;championId&#34;`
    ProfileIconId <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;profileIconId&#34;`
    Spell1Id      <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;spell1Id&#34;`
    Spell2Id      <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;spell2Id&#34;`
    SummonerName  <a href="/pkg/builtin/#string">string</a> `json:&#34;summonerName&#34;`
    TeamId        <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;teamId&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="Frame">type <a href="/src/target/match.go?s=10027:10256#L174">Frame</a></h2>
			<pre>type Frame struct {
    Events            []<a href="#Event">Event</a>                     `json:&#34;events&#34;`
    ParticipantFrames map[<a href="/pkg/builtin/#string">string</a>]<a href="#ParticipantFrame">ParticipantFrame</a> `json:&#34;participantsFrames&#34;`
    Timestamp         <a href="/pkg/builtin/#int64">int64</a>                       `json:&#34;timestamp&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="Game">type <a href="/src/target/game.go?s=85:741#L1">Game</a></h2>
			<pre>type Game struct {
    ChampionID    <a href="/pkg/builtin/#int">int</a>      `json:&#34;championId&#34;`
    CreateDate    <a href="/pkg/builtin/#int64">int64</a>    `json:&#34;createDate&#34;`
    FellowPlayers []<a href="#Player">Player</a> `json:&#34;fellowPlayers&#34;`
    GameID        <a href="/pkg/builtin/#int64">int64</a>    `json:&#34;gameId&#34;`
    GameMode      <a href="/pkg/builtin/#string">string</a>   `json:&#34;gameMode&#34;`
    GameType      <a href="/pkg/builtin/#string">string</a>   `json:&#34;gameType&#34;`
    Invalid       <a href="/pkg/builtin/#bool">bool</a>     `json:&#34;invalid&#34;`
    IPEarned      <a href="/pkg/builtin/#int">int</a>      `json:&#34;ipEarned&#34;`
    Level         <a href="/pkg/builtin/#int">int</a>      `json:&#34;level&#34;`
    MapID         <a href="/pkg/builtin/#int">int</a>      `json:&#34;mapId&#34;`
    Spell1        <a href="/pkg/builtin/#int">int</a>      `json:&#34;spell1&#34;`
    Spell2        <a href="/pkg/builtin/#int">int</a>      `json:&#34;spell2&#34;`
    Statistics    <a href="#GameStat">GameStat</a> `json:&#34;stats&#34;`
    SubType       <a href="/pkg/builtin/#string">string</a>   `json:&#34;subType&#34;`
    TeamID        <a href="/pkg/builtin/#int">int</a>      `json:&#34;teamId&#34;`
}</pre>
			<p>
Game represents a League of Legends game
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="GameStat">type <a href="/src/target/game.go?s=1018:5977#L24">GameStat</a></h2>
			<pre>type GameStat struct {
    Assists                         <a href="/pkg/builtin/#int">int</a>  `json:&#34;assists&#34;`
    BarracksKilled                  <a href="/pkg/builtin/#int">int</a>  `json:&#34;barracksKilled&#34;`
    ChampionsKilled                 <a href="/pkg/builtin/#int">int</a>  `json:&#34;championsKilled&#34;`
    CombatPlayerScore               <a href="/pkg/builtin/#int">int</a>  `json:&#34;combatPlayerScore&#34;`
    ConsumablesPurchased            <a href="/pkg/builtin/#int">int</a>  `json:&#34;consumablesPurchased&#34;`
    DamageDealtPlayer               <a href="/pkg/builtin/#int">int</a>  `json:&#34;damageDealtPlayer&#34;`
    DoubleKills                     <a href="/pkg/builtin/#int">int</a>  `json:&#34;doubleKills&#34;`
    FirstBlood                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;firstBlood&#34;`
    Gold                            <a href="/pkg/builtin/#int">int</a>  `json:&#34;gold&#34;`
    GoldEarned                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;goldEarned&#34;`
    GoldSpent                       <a href="/pkg/builtin/#int">int</a>  `json:&#34;goldSpent&#34;`
    Item0                           <a href="/pkg/builtin/#int">int</a>  `json:&#34;item0&#34;`
    Item1                           <a href="/pkg/builtin/#int">int</a>  `json:&#34;item1&#34;`
    Item2                           <a href="/pkg/builtin/#int">int</a>  `json:&#34;item2&#34;`
    Item3                           <a href="/pkg/builtin/#int">int</a>  `json:&#34;item3&#34;`
    Item4                           <a href="/pkg/builtin/#int">int</a>  `json:&#34;item4&#34;`
    Item5                           <a href="/pkg/builtin/#int">int</a>  `json:&#34;item5&#34;`
    Item6                           <a href="/pkg/builtin/#int">int</a>  `json:&#34;item6&#34;`
    ItemsPurchased                  <a href="/pkg/builtin/#int">int</a>  `json:&#34;itemsPurchased&#34;`
    KillingSprees                   <a href="/pkg/builtin/#int">int</a>  `json:&#34;killingSprees&#34;`
    LargestCriticalStrike           <a href="/pkg/builtin/#int">int</a>  `json:&#34;largestCriticalStrike&#34;`
    LargestKillingSpree             <a href="/pkg/builtin/#int">int</a>  `json:&#34;largestKillingSpree&#34;`
    LargestMultiKill                <a href="/pkg/builtin/#int">int</a>  `json:&#34;largestMultiKill&#34;`
    LegendaryItemsCreated           <a href="/pkg/builtin/#int">int</a>  `json:&#34;legendaryItemsCreated&#34;`
    Level                           <a href="/pkg/builtin/#int">int</a>  `json:&#34;level&#34;`
    MagicDamageDealtPlayer          <a href="/pkg/builtin/#int">int</a>  `json:&#34;magicDamageDealtPlayer&#34;`
    MagicDamageDealtToChampions     <a href="/pkg/builtin/#int">int</a>  `json:&#34;magicDamageDealtToChampions&#34;`
    MagicDamageTaken                <a href="/pkg/builtin/#int">int</a>  `json:&#34;magicDamageTaken&#34;`
    MinionsDenied                   <a href="/pkg/builtin/#int">int</a>  `json:&#34;minionsDenied&#34;`
    MinionsKilled                   <a href="/pkg/builtin/#int">int</a>  `json:&#34;minionsKilled&#34;`
    NeutralMinionsKilled            <a href="/pkg/builtin/#int">int</a>  `json:&#34;neutralMinionsKilled&#34;`
    NeutralMinionsKilledEnemyJungle <a href="/pkg/builtin/#int">int</a>  `json:&#34;neutralMinionsKilledEnemyJungle&#34;`
    NeutralMinionsKilledYourJungle  <a href="/pkg/builtin/#int">int</a>  `json:&#34;neutralMinionsKilledYourJungle&#34;`
    NexusKilled                     <a href="/pkg/builtin/#bool">bool</a> `json:&#34;nexusKilled&#34;`
    NodeCapture                     <a href="/pkg/builtin/#int">int</a>  `json:&#34;nodeCapture&#34;`
    NodeCaptureAssist               <a href="/pkg/builtin/#int">int</a>  `json:&#34;nodeCaptureAssist&#34;`
    NodeNeutralize                  <a href="/pkg/builtin/#int">int</a>  `json:&#34;nodeNeutralize&#34;`
    NodeNeutralizeAssist            <a href="/pkg/builtin/#int">int</a>  `json:&#34;nodeNeutralizeAssist&#34;`
    NumDeaths                       <a href="/pkg/builtin/#int">int</a>  `json:&#34;numDeaths&#34;`
    NumItemsBought                  <a href="/pkg/builtin/#int">int</a>  `json:&#34;numItemsBought&#34;`
    ObjectivePlayerScore            <a href="/pkg/builtin/#int">int</a>  `json:&#34;objectivePlayerScore&#34;`
    PentaKills                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;pentaKills&#34;`
    PhysicalDamageDealtPlayer       <a href="/pkg/builtin/#int">int</a>  `json:&#34;physicalDamageDealtPlayer&#34;`
    PhysicalDamageDealtToChampions  <a href="/pkg/builtin/#int">int</a>  `json:&#34;physicalDamageDealtToChampions&#34;`
    PhysicalDamageTaken             <a href="/pkg/builtin/#int">int</a>  `json:&#34;physicalDamageTaken&#34;`
    PlayerPosition                  <a href="/pkg/builtin/#int">int</a>  `json:&#34;playerPosition&#34;`
    PlayerRole                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;playerRole&#34;`
    QuadraKills                     <a href="/pkg/builtin/#int">int</a>  `json:&#34;quadraKills&#34;`
    SightWardsBought                <a href="/pkg/builtin/#int">int</a>  `json:&#34;sightWardsBought&#34;`
    Spell1Cast                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;spell1Cast&#34;`
    Spell2Cast                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;spell2Cast&#34;`
    Spell3Cast                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;spell3Cast&#34;`
    Spell4Cast                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;spell4Cast&#34;`
    SummonSpell1Cast                <a href="/pkg/builtin/#int">int</a>  `json:&#34;summonSpell1Cast&#34;`
    SummonSpell2Cast                <a href="/pkg/builtin/#int">int</a>  `json:&#34;summonSpell2Cast&#34;`
    SuperMonsterKilled              <a href="/pkg/builtin/#int">int</a>  `json:&#34;superMonsterKilled&#34;`
    Team                            <a href="/pkg/builtin/#int">int</a>  `json:&#34;team&#34;`
    TeamObjective                   <a href="/pkg/builtin/#int">int</a>  `json:&#34;teamObjective&#34;`
    TimePlayed                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;timePlayed&#34;`
    TotalDamageDealt                <a href="/pkg/builtin/#int">int</a>  `json:&#34;totalDamageDealt&#34;`
    TotalDamageDealtToChampions     <a href="/pkg/builtin/#int">int</a>  `json:&#34;totalDamageDealtToChampions&#34;`
    TotalDamageTaken                <a href="/pkg/builtin/#int">int</a>  `json:&#34;totalDamageTaken&#34;`
    TotalHeal                       <a href="/pkg/builtin/#int">int</a>  `json:&#34;totalHeal&#34;`
    TotalPlayerScore                <a href="/pkg/builtin/#int">int</a>  `json:&#34;totalPlayerScore&#34;`
    TotalScoreRank                  <a href="/pkg/builtin/#int">int</a>  `json:&#34;totalScoreRank&#34;`
    TotalTimeCrowdControlDealt      <a href="/pkg/builtin/#int">int</a>  `json:&#34;totalTimeCrowdControlDealt&#34;`
    TotalUnitsHealed                <a href="/pkg/builtin/#int">int</a>  `json:&#34;totalUnitsHealed&#34;`
    TripleKills                     <a href="/pkg/builtin/#int">int</a>  `json:&#34;tripleKills&#34;`
    TrueDamageDealtPlayer           <a href="/pkg/builtin/#int">int</a>  `json:&#34;trueDamageDealtPlayer&#34;`
    TrueDamageDealtToChampions      <a href="/pkg/builtin/#int">int</a>  `json:&#34;trueDamageDealtToChampions&#34;`
    TrueDamageTaken                 <a href="/pkg/builtin/#int">int</a>  `json:&#34;trueDamageTaken&#34;`
    TurretsKilled                   <a href="/pkg/builtin/#int">int</a>  `json:&#34;turretsKilled&#34;`
    UnrealKills                     <a href="/pkg/builtin/#int">int</a>  `json:&#34;unrealKills&#34;`
    VictoryPointTotal               <a href="/pkg/builtin/#int">int</a>  `json:&#34;victoryPointTotal&#34;`
    VisionWardsBought               <a href="/pkg/builtin/#int">int</a>  `json:&#34;visionWardsBought&#34;`
    WardKilled                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;wardKilled&#34;`
    WardPlaced                      <a href="/pkg/builtin/#int">int</a>  `json:&#34;wardPlaced&#34;`
    Win                             <a href="/pkg/builtin/#bool">bool</a> `json:&#34;win&#34;`
}</pre>
			<p>
GameStat represents a stat for the assosiated Game
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="League">type <a href="/src/target/league.go?s=89:340#L1">League</a></h2>
			<pre>type League struct {
    Entries       []<a href="#LeagueItem">LeagueItem</a> `json:&#34;entries&#34;`
    Name          <a href="/pkg/builtin/#string">string</a>       `json:&#34;name&#34;`
    ParticipantId <a href="/pkg/builtin/#string">string</a>       `json:&#34;participantId&#34;`
    Queue         <a href="/pkg/builtin/#string">string</a>       `json:&#34;queue&#34;`
    Tier          <a href="/pkg/builtin/#string">string</a>       `json:&#34;tier&#34;`
}</pre>
			<p>
League represents a League of Legends league
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="LeagueItem">type <a href="/src/target/league.go?s=414:993#L7">LeagueItem</a></h2>
			<pre>type LeagueItem struct {
    Division         <a href="/pkg/builtin/#string">string</a>     `json:&#34;division&#34;`
    IsFreshBlood     <a href="/pkg/builtin/#bool">bool</a>       `json:&#34;isFreshBlood&#34;`
    IsHotStreak      <a href="/pkg/builtin/#bool">bool</a>       `json:&#34;isHotStreak&#34;`
    IsInactive       <a href="/pkg/builtin/#bool">bool</a>       `json:&#34;isInactive&#34;`
    IsVeteran        <a href="/pkg/builtin/#bool">bool</a>       `json:&#34;isVeteran&#34;`
    LeaguePoints     <a href="/pkg/builtin/#int">int</a>        `json:&#34;leaguePoints&#34;`
    Losses           <a href="/pkg/builtin/#int">int</a>        `json:&#34;losses&#34;`
    MiniSeries       <a href="#MiniSeries">MiniSeries</a> `json:&#34;miniSeries&#34;`
    PlayerOrTeamID   <a href="/pkg/builtin/#string">string</a>     `json:&#34;playerOrTeamId&#34;`
    PlayerOrTeamName <a href="/pkg/builtin/#string">string</a>     `json:&#34;playerOrTeamName&#34;`
    Wins             <a href="/pkg/builtin/#int">int</a>        `json:&#34;wins&#34;`
}</pre>
			<p>
LeagueItem is an entry in a League. It represents a player or team
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="Mastery">type <a href="/src/target/match.go?s=2924:3017#L56">Mastery</a></h2>
			<pre>type Mastery struct {
    MasteryId <a href="/pkg/builtin/#int64">int64</a> `json:&#34;masteryId&#34;`
    Rank      <a href="/pkg/builtin/#int64">int64</a> `json:&#34;rank&#34;`
}</pre>
			<p>
Mastery being used
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="MasteryDto">type <a href="/src/target/summoner.go?s=448:523#L11">MasteryDto</a></h2>
			<pre>type MasteryDto struct {
    ID   <a href="/pkg/builtin/#int">int</a> `json:&#34;id&#34;`
    Rank <a href="/pkg/builtin/#int">int</a> `json:&#34;rank&#34;`
}</pre>
			<p>
Mastery located inside a page
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="MasteryPage">type <a href="/src/target/summoner.go?s=221:411#L3">MasteryPage</a></h2>
			<pre>type MasteryPage struct {
    Current   <a href="/pkg/builtin/#bool">bool</a>         `json:&#34;current&#34;`
    ID        <a href="/pkg/builtin/#int64">int64</a>        `json:&#34;id&#34;`
    Name      <a href="/pkg/builtin/#string">string</a>       `json:&#34;name&#34;`
    Masteries []<a href="#MasteryDto">MasteryDto</a> `json:&#34;masteries&#34;`
}</pre>
			<p>
MasteryPage represents a League of Legends mastery page
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="MatchDetail">type <a href="/src/target/match.go?s=41:1054#L1">MatchDetail</a></h2>
			<pre>type MatchDetail struct {
    MapID                 <a href="/pkg/builtin/#int">int</a>                   `json:&#34;mapId&#34;`
    MatchCreation         <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;matchCreation&#34;`
    MatchDuration         <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;matchDuration&#34;`
    MatchID               <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;matchId&#34;`
    MatchMode             <a href="/pkg/builtin/#string">string</a>                `json:&#34;matchMode&#34;`
    MatchType             <a href="/pkg/builtin/#string">string</a>                `json:&#34;matchType&#34;`
    MatchVersion          <a href="/pkg/builtin/#string">string</a>                `json:&#34;matchVersion&#34;`
    ParticipantIdentities []<a href="#ParticipantIdentity">ParticipantIdentity</a> `json:&#34;participantIdentities&#34;`
    Participants          []<a href="#Participant">Participant</a>         `json:&#34;participants&#34;`
    PlatformId            <a href="/pkg/builtin/#string">string</a>                `json:&#34;platformId&#34;`
    QueueType             <a href="/pkg/builtin/#string">string</a>                `json:&#34;queueType&#34;`
    Region                <a href="/pkg/builtin/#string">string</a>                `json:&#34;region&#34;`
    Season                <a href="/pkg/builtin/#string">string</a>                `json:&#34;season&#34;`
    Teams                 []<a href="#TeamGameDetails">TeamGameDetails</a>     `json:&#34;teams&#34;`
    Timeline              <a href="#Timeline">Timeline</a>              `json:&#34;timeline&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="MatchHistorySummary">type <a href="/src/target/team.go?s=1231:1758#L16">MatchHistorySummary</a></h2>
			<pre>type MatchHistorySummary struct {
    Assists           <a href="/pkg/builtin/#int">int</a>    `json:&#34;assists&#34;`
    Date              <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;date&#34;`
    Deaths            <a href="/pkg/builtin/#int">int</a>    `json:&#34;deaths&#34;`
    GameID            <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;gameId&#34;`
    GameMode          <a href="/pkg/builtin/#string">string</a> `json:&#34;gameMode&#34;`
    Invalid           <a href="/pkg/builtin/#bool">bool</a>   `json:&#34;invalid&#34;`
    Kills             <a href="/pkg/builtin/#int">int</a>    `json:&#34;kills&#34;`
    MapID             <a href="/pkg/builtin/#int">int</a>    `json:&#34;mapId&#34;`
    OpposingTeamKills <a href="/pkg/builtin/#int">int</a>    `json:&#34;opposingTeamKills&#34;`
    OpposingTeamName  <a href="/pkg/builtin/#string">string</a> `json:&#34;opposingTeamName&#34;`
    Win               <a href="/pkg/builtin/#bool">bool</a>   `json:&#34;win&#34;`
}</pre>
			<p>
MatchHistorySummary is a summary of a matches played by a Team
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="MatchPlayer">type <a href="/src/target/match.go?s=9700:9916#L162">MatchPlayer</a></h2>
			<pre>type MatchPlayer struct {
    MatchHistoryURI <a href="/pkg/builtin/#string">string</a> `json:&#34;matchHistoryUri&#34;`
    ProfileIcon     <a href="/pkg/builtin/#int">int</a>    `json:&#34;profileIcon&#34;`
    SummonerId      <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;summonerId&#34;`
    SummonerName    <a href="/pkg/builtin/#string">string</a> `json:&#34;summonerName&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="MatchSummary">type <a href="/src/target/matchhistory.go?s=128:1017#L2">MatchSummary</a></h2>
			<pre>type MatchSummary struct {
    MapID                 <a href="/pkg/builtin/#int">int</a>                   `json:&#34;mapId&#34;`
    MatchCreation         <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;matchCreation&#34;`
    MatchDuration         <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;matchDuration&#34;`
    MatchId               <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;matchId&#34;`
    MatchMode             <a href="/pkg/builtin/#string">string</a>                `json:&#34;matchMode&#34;`
    MatchType             <a href="/pkg/builtin/#string">string</a>                `json:&#34;matchType&#34;`
    MatchVersion          <a href="/pkg/builtin/#string">string</a>                `json:&#34;matchVersion&#34;`
    ParticipantIdentities []<a href="#ParticipantIdentity">ParticipantIdentity</a> `json:&#34;participantIdentities&#34;`
    Participants          []<a href="#Participant">Participant</a>         `json:&#34;participants&#34;`
    PlatformId            <a href="/pkg/builtin/#string">string</a>                `json:&#34;platformId&#34;`
    QueueType             <a href="/pkg/builtin/#string">string</a>                `json:&#34;queueType&#34;`
    Region                <a href="/pkg/builtin/#string">string</a>                `json:&#34;region&#34;`
    Season                <a href="/pkg/builtin/#string">string</a>                `json:&#34;season&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="MiniSeries">type <a href="/src/target/league.go?s=1079:1242#L22">MiniSeries</a></h2>
			<pre>type MiniSeries struct {
    Losses   <a href="/pkg/builtin/#int">int</a>    `json:&#34;losses&#34;`
    Progress <a href="/pkg/builtin/#string">string</a> `json:&#34;progress&#34;`
    Target   <a href="/pkg/builtin/#int">int</a>    `json:&#34;target&#34;`
    Wins     <a href="/pkg/builtin/#int">int</a>    `json:&#34;wins&#34;`
}</pre>
			<p>
MiniSeries shows if a player is in their Series and how far they are within it
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="Observer">type <a href="/src/target/curentgame.go?s=1455:1526#L30">Observer</a></h2>
			<pre>type Observer struct {
    EncryptionKey <a href="/pkg/builtin/#string">string</a> `json:&#34;encryptionKey&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="Participant">type <a href="/src/target/match.go?s=1058:1763#L15">Participant</a></h2>
			<pre>type Participant struct {
    ChampionID                <a href="/pkg/builtin/#int">int</a>                 `json:&#34;championId&#34;`
    HighestAchievedSeasonTier <a href="/pkg/builtin/#string">string</a>              `json:&#34;highestAchievedSeasonTier&#34;`
    Masteries                 []<a href="#Mastery">Mastery</a>           `json:&#34;masteries&#34;`
    ParticipantID             <a href="/pkg/builtin/#int">int</a>                 `json:&#34;participantId&#34;`
    Runes                     []<a href="#Rune">Rune</a>              `json:&#34;runes&#34;`
    Spell1ID                  <a href="/pkg/builtin/#int">int</a>                 `json:&#34;spell1Id&#34;`
    Spell2ID                  <a href="/pkg/builtin/#int">int</a>                 `json:&#34;spell2Id&#34;`
    Stats                     <a href="#ParticipantStats">ParticipantStats</a>    `json:&#34;stats&#34;`
    TeamID                    <a href="/pkg/builtin/#int">int</a>                 `json:&#34;teamId&#34;`
    Timeline                  <a href="#ParticipantTimeline">ParticipantTimeline</a> `json:&#34;timeline&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="ParticipantFrame">type <a href="/src/target/match.go?s=11671:12216#L211">ParticipantFrame</a></h2>
			<pre>type ParticipantFrame struct {
    CurrentGold         <a href="/pkg/builtin/#int">int</a>      `json:&#34;currentGold&#34;`
    DominionScore       <a href="/pkg/builtin/#int">int</a>      `json:&#34;dominionScore&#34;`
    JungleMinionsKilled <a href="/pkg/builtin/#int">int</a>      `json:&#34;jungleMinionsKilled&#34;`
    Level               <a href="/pkg/builtin/#int">int</a>      `json:&#34;level&#34;`
    MinionsKilled       <a href="/pkg/builtin/#int">int</a>      `json:&#34;minionsKilled&#34;`
    ParticipantID       <a href="/pkg/builtin/#int">int</a>      `json:&#34;participantId&#34;`
    Position            <a href="#Position">Position</a> `json:&#34;position&#34;`
    TeamScore           <a href="/pkg/builtin/#int">int</a>      `json:&#34;teamScore&#34;`
    TotalGold           <a href="/pkg/builtin/#int">int</a>      `json:&#34;totalGold&#34;`
    Xp                  <a href="/pkg/builtin/#int">int</a>      `json:&#34;xp&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="ParticipantIdentity">type <a href="/src/target/match.go?s=1767:1898#L28">ParticipantIdentity</a></h2>
			<pre>type ParticipantIdentity struct {
    ParticipantId <a href="/pkg/builtin/#int">int</a>         `json:&#34;participantId&#34;`
    Player        <a href="#MatchPlayer">MatchPlayer</a> `json:&#34;player&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="ParticipantStats">type <a href="/src/target/match.go?s=3021:7133#L61">ParticipantStats</a></h2>
			<pre>type ParticipantStats struct {
    Assists                         <a href="/pkg/builtin/#int64">int64</a> `json:&#34;assists&#34;`
    ChampLevel                      <a href="/pkg/builtin/#int64">int64</a> `json:&#34;champLevel&#34;`
    CombatPlayerScore               <a href="/pkg/builtin/#int64">int64</a> `json:&#34;combatPlayerScore&#34;`
    Deaths                          <a href="/pkg/builtin/#int64">int64</a> `json:&#34;deaths&#34;`
    DoubleKills                     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;doubleKills&#34;`
    FirstBloodAssist                <a href="/pkg/builtin/#bool">bool</a>  `json:&#34;firstBloodAssist&#34;`
    FirstBloodKill                  <a href="/pkg/builtin/#bool">bool</a>  `json:&#34;firstBloodKill&#34;`
    FirstInhibitorAssist            <a href="/pkg/builtin/#bool">bool</a>  `json:&#34;firstInhibitorAssist&#34;`
    FirstInhibitorKill              <a href="/pkg/builtin/#bool">bool</a>  `json:&#34;firstInhibitorKill&#34;`
    FirstTowerAssist                <a href="/pkg/builtin/#bool">bool</a>  `json:&#34;firstTowerAssist&#34;`
    FirstTowerKill                  <a href="/pkg/builtin/#bool">bool</a>  `json:&#34;firstTowerKill&#34;`
    GoldEarned                      <a href="/pkg/builtin/#int64">int64</a> `json:&#34;goldEarned&#34;`
    GoldSpent                       <a href="/pkg/builtin/#int64">int64</a> `json:&#34;goldSpent&#34;`
    InhibitorKills                  <a href="/pkg/builtin/#int64">int64</a> `json:&#34;inhibitorKills&#34;`
    Item0                           <a href="/pkg/builtin/#int64">int64</a> `json:&#34;item0&#34;`
    Item1                           <a href="/pkg/builtin/#int64">int64</a> `json:&#34;item1&#34;`
    Item2                           <a href="/pkg/builtin/#int64">int64</a> `json:&#34;item2&#34;`
    Item3                           <a href="/pkg/builtin/#int64">int64</a> `json:&#34;item3&#34;`
    Item4                           <a href="/pkg/builtin/#int64">int64</a> `json:&#34;item4&#34;`
    Item5                           <a href="/pkg/builtin/#int64">int64</a> `json:&#34;item5&#34;`
    Item6                           <a href="/pkg/builtin/#int64">int64</a> `json:&#34;item6&#34;`
    KillingSprees                   <a href="/pkg/builtin/#int64">int64</a> `json:&#34;killingSprees&#34;`
    Kills                           <a href="/pkg/builtin/#int64">int64</a> `json:&#34;kills&#34;`
    LargestCriticalStrike           <a href="/pkg/builtin/#int64">int64</a> `json:&#34;largestCriticalStrike&#34;`
    LargestKillingSpree             <a href="/pkg/builtin/#int64">int64</a> `json:&#34;largestKillingSpree&#34;`
    LargestMultiKill                <a href="/pkg/builtin/#int64">int64</a> `json:&#34;largestMultiKill&#34;`
    MagicDamageDealt                <a href="/pkg/builtin/#int64">int64</a> `json:&#34;magicDamageDealt&#34;`
    MagicDamageDealtToChampions     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;magicDamageDealtToChampions&#34;`
    MagicDamageTaken                <a href="/pkg/builtin/#int64">int64</a> `json:&#34;magicDamageTaken&#34;`
    MinionsKilled                   <a href="/pkg/builtin/#int64">int64</a> `json:&#34;minionsKilled&#34;`
    NeutralMinionsKilled            <a href="/pkg/builtin/#int64">int64</a> `json:&#34;neutralMinionsKilled&#34;`
    NeutralMinionsKilledEnemyJungle <a href="/pkg/builtin/#int64">int64</a> `json:&#34;neutralMinionsKilledEnemyJungle&#34;`
    NeutralMinionsKilledTeamJungle  <a href="/pkg/builtin/#int64">int64</a> `json:&#34;neutralMinionsKilledTeamJungle&#34;`
    NodeCapture                     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;nodeCapture&#34;`
    NodeCaptureAssist               <a href="/pkg/builtin/#int64">int64</a> `json:&#34;nodeCaptureAssist&#34;`
    NodeNeutralize                  <a href="/pkg/builtin/#int64">int64</a> `json:&#34;nodeNeutralize&#34;`
    NodeNeutralizeAssist            <a href="/pkg/builtin/#int64">int64</a> `json:&#34;nodeNeutralizeAssist&#34;`
    ObjectivePlayerScore            <a href="/pkg/builtin/#int64">int64</a> `json:&#34;objectivePlayerScore&#34;`
    PentaKills                      <a href="/pkg/builtin/#int64">int64</a> `json:&#34;pentaKills&#34;`
    PhysicalDamageDealt             <a href="/pkg/builtin/#int64">int64</a> `json:&#34;physicalDamageDealt&#34;`
    PhysicalDamageDealtToChampions  <a href="/pkg/builtin/#int64">int64</a> `json:&#34;physicalDamageDealtToChampions&#34;`
    PhysicalDamageTaken             <a href="/pkg/builtin/#int64">int64</a> `json:&#34;physicalDamageTaken&#34;`
    QuadraKills                     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;quadraKills&#34;`
    SightWardsBoughtInGame          <a href="/pkg/builtin/#int64">int64</a> `json:&#34;sightWardsBoughtInGame&#34;`
    TeamObjective                   <a href="/pkg/builtin/#int64">int64</a> `json:&#34;teamObjective&#34;`
    TotalDamageDealt                <a href="/pkg/builtin/#int64">int64</a> `json:&#34;totalDamageDealt&#34;`
    TotalDamageDealtToChampions     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;totalDamageDealtToChampions&#34;`
    TotalDamageTaken                <a href="/pkg/builtin/#int64">int64</a> `json:&#34;totalDamageTaken&#34;`
    TotalHeal                       <a href="/pkg/builtin/#int64">int64</a> `json:&#34;totalHeal&#34;`
    TotalPlayerScore                <a href="/pkg/builtin/#int64">int64</a> `json:&#34;totalPlayerScore&#34;`
    TotalScoreRank                  <a href="/pkg/builtin/#int64">int64</a> `json:&#34;totalScoreRank&#34;`
    TotalTimeCrowdControlDealt      <a href="/pkg/builtin/#int64">int64</a> `json:&#34;totalTimeCrowdControlDealt&#34;`
    TotalUnitsHealed                <a href="/pkg/builtin/#int64">int64</a> `json:&#34;totalUnitsHealed&#34;`
    TowerKills                      <a href="/pkg/builtin/#int64">int64</a> `json:&#34;towerKills&#34;`
    TripleKills                     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;tripleKills&#34;`
    TrueDamageDealt                 <a href="/pkg/builtin/#int64">int64</a> `json:&#34;trueDamageDealt&#34;`
    TrueDamageDealtToChampions      <a href="/pkg/builtin/#int64">int64</a> `json:&#34;trueDamageDealtToChampions&#34;`
    TrueDamageTaken                 <a href="/pkg/builtin/#int64">int64</a> `json:&#34;trueDamageTaken&#34;`
    UnrealKills                     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;unrealKills&#34;`
    VisionWardsBoughtInGame         <a href="/pkg/builtin/#int64">int64</a> `json:&#34;visionWardsBoughtInGame&#34;`
    WardsKilled                     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;wardsKilled&#34;`
    WardsPlaced                     <a href="/pkg/builtin/#int64">int64</a> `json:&#34;wardsPlaced&#34;`
    Winner                          <a href="/pkg/builtin/#bool">bool</a>  `json:&#34;winner&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="ParticipantTimeline">type <a href="/src/target/match.go?s=7137:9611#L127">ParticipantTimeline</a></h2>
			<pre>type ParticipantTimeline struct {
    AncientGolemAssistsPerMinCounts <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;ancientGolemAssistsPerMinCounts&#34;`
    AncientGolemKillsPerMinCounts   <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;ancientGolemKillsPerMinCounts&#34;`
    AssistedLaneDeathsPerMinDeltas  <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;assistedLaneDeathsPerMinDeltas&#34;`
    AssistedLaneKillsPerMinDeltas   <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;assistedLaneKillsPerMinDeltas&#34;`
    BaronAssistsPerMinCounts        <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;baronAssistsPerMinCounts&#34;`
    BaronKillsPerMinCounts          <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;baronKillsPerMinCounts&#34;`
    CreepsPerMinDeltas              <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;creepsPerMinDeltas&#34;`
    CsDiffPerMinDeltas              <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;csDiffPerMinDeltas&#34;`
    DamageTakenDiffPerMinDeltas     <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;damageTakenDiffPerMinDeltas&#34;`
    DamageTakenPerMinDeltas         <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;damageTakenPerMinDeltas&#34;`
    DragonAssistsPerMinCounts       <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;dragonAssistsPerMinCounts&#34;`
    DragonKillsPerMinCounts         <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;dragonKillsPerMinCounts&#34;`
    ElderLizardAssistsPerMinCounts  <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;elderLizardAssistsPerMinCounts&#34;`
    ElderLizardKillsPerMinCounts    <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;elderLizardKillsPerMinCounts&#34;`
    GoldPerMinDeltas                <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;goldPerMinDeltas&#34;`
    InhibitorAssistsPerMinCounts    <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;inhibitorAssistsPerMinCounts&#34;`
    InhibitorKillsPerMinCounts      <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;inhibitorKillsPerMinCounts&#34;`
    Lane                            <a href="/pkg/builtin/#string">string</a>                  `json:&#34;lane&#34;`
    Role                            <a href="/pkg/builtin/#string">string</a>                  `json:&#34;role&#34;`
    TowerAssistsPerMinCounts        <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;towerAssistsPerMinCounts&#34;`
    TowerKillsPerMinCounts          <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;towerKillsPerMinCounts&#34;`
    TowerKillsPerMinDeltas          <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;towerKillsPerMinDeltas&#34;`
    VilemawAssistsPerMinCounts      <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;vilemawAssistsPerMinCounts&#34;`
    VilemawKillsPerMinCounts        <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;vilemawKillsPerMinCounts&#34;`
    WardsPerMinDeltas               <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;wardsPerMinDeltas&#34;`
    XpDiffPerMinDeltas              <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;xpDiffPerMinDeltas&#34;`
    XpPerMinDeltas                  <a href="#ParticipantTimelineData">ParticipantTimelineData</a> `json:&#34;xpPerMinDeltas&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="ParticipantTimelineData">type <a href="/src/target/match.go?s=10260:10485#L180">ParticipantTimelineData</a></h2>
			<pre>type ParticipantTimelineData struct {
    TenToTwenty    <a href="/pkg/builtin/#float64">float64</a> `json:&#34;tenToTwenty&#34;`
    ThirtyToEnd    <a href="/pkg/builtin/#float64">float64</a> `json:&#34;thirtyToEnd&#34;`
    TwentyToThirty <a href="/pkg/builtin/#float64">float64</a> `json:&#34;twentyToThirty&#34;`
    ZeroToTen      <a href="/pkg/builtin/#float64">float64</a> `json:&#34;zeroToTen&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="Player">type <a href="/src/target/game.go?s=824:960#L17">Player</a></h2>
			<pre>type Player struct {
    ChampionID <a href="/pkg/builtin/#int">int</a>   `json:&#34;championId&#34;`
    SummonerID <a href="/pkg/builtin/#int64">int64</a> `json:&#34;summonerId&#34;`
    TeamID     <a href="/pkg/builtin/#int">int</a>   `json:&#34;teamId&#34;`
}</pre>
			<p>
Player represents a League of Legends player that was in the requested game
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="PlayerHistory">type <a href="/src/target/matchhistory.go?s=52:124#L1">PlayerHistory</a></h2>
			<pre>type PlayerHistory struct {
    Matches []<a href="#MatchSummary">MatchSummary</a> `json:&#34;matches&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="PlayerStatsSummary">type <a href="/src/target/stats.go?s=305:646#L3">PlayerStatsSummary</a></h2>
			<pre>type PlayerStatsSummary struct {
    AggregatedStats       <a href="#AggregatedStats">AggregatedStats</a> `json:&#34;aggregatedStats&#34;`
    Losses                <a href="/pkg/builtin/#int">int</a>             `json:&#34;losses&#34;`
    ModifyDate            <a href="/pkg/builtin/#int64">int64</a>           `json:&#34;modifyDate&#34;`
    PlayerStatSummaryType <a href="/pkg/builtin/#string">string</a>          `json:&#34;playerStatSummaryType&#34;`
    Wins                  <a href="/pkg/builtin/#int">int</a>             `json:&#34;wins&#34;`
}</pre>
			<p>
PlayerStatsSummary represents a summary of a League of Legends player&#39;s game stats
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="Position">type <a href="/src/target/match.go?s=12220:12283#L224">Position</a></h2>
			<pre>type Position struct {
    X <a href="/pkg/builtin/#int">int</a> `json:&#34;x&#34;`
    Y <a href="/pkg/builtin/#int">int</a> `json:&#34;y&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="RankedStats">type <a href="/src/target/stats.go?s=4328:4502#L72">RankedStats</a></h2>
			<pre>type RankedStats struct {
    Champions  []<a href="#ChampionStats">ChampionStats</a> `json:&#34;champions&#34;`
    ModifyDate <a href="/pkg/builtin/#int64">int64</a>           `json:&#34;modifyDate&#34;`
    SummonerID <a href="/pkg/builtin/#int64">int64</a>           `json:&#34;summonerId&#34;`
}</pre>
			<p>
RankedStats represents a League of Legends player&#39;s statistics for ranked play
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="RiotAPI">type <a href="/src/target/goriot.go?s=2147:2267#L62">RiotAPI</a></h2>
			<pre>type RiotAPI struct {
    APIkey <a href="/pkg/builtin/#string">string</a>
    Region <a href="/pkg/builtin/#string">string</a>
    <span class="comment">// contains filtered or unexported fields</span>
}</pre>
			

			

			

			
			
			

			
				
				<h3 id="Get">func <a href="/src/target/goriot.go?s=2271:2346#L69">Get</a></h3>
				<pre>func Get(key <a href="/pkg/builtin/#string">string</a>, region <a href="/pkg/builtin/#string">string</a>, smallLimit <a href="/pkg/builtin/#int">int</a>, longLimit <a href="/pkg/builtin/#int">int</a>) *<a href="#RiotAPI">RiotAPI</a></pre>
				
				
				
			

			
				
				<h3 id="RiotAPI.ChampionByID">func (*RiotAPI) <a href="/src/target/champion.go?s=762:844#L18">ChampionByID</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) ChampionByID(championID <a href="/pkg/builtin/#string">string</a>) (champion <a href="#Champion">Champion</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.ChampionList">func (*RiotAPI) <a href="/src/target/champion.go?s=445:532#L11">ChampionList</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) ChampionList(freetoplay <a href="/pkg/builtin/#string">string</a>) (champions championList, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.CurrentGameBySummonerID">func (*RiotAPI) <a href="/src/target/curentgame.go?s=1646:1767#L40">CurrentGameBySummonerID</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) CurrentGameBySummonerID(plateform <a href="/pkg/builtin/#string">string</a>, summonerID <a href="/pkg/builtin/#string">string</a>) (currentGame <a href="#CurrentGameInfo">CurrentGameInfo</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.FeaturedGames">func (*RiotAPI) <a href="/src/target/featuredgame.go?s=1450:1539#L32">FeaturedGames</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) FeaturedGames(region <a href="/pkg/builtin/#string">string</a>) (featuredGames <a href="#FeaturedGames">FeaturedGames</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.LeagueByChallenger">func (*RiotAPI) <a href="/src/target/league.go?s=2456:2535#L53">LeagueByChallenger</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) LeagueByChallenger(queue <a href="/pkg/builtin/#string">string</a>) (league <a href="#League">League</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.LeagueByMaster">func (*RiotAPI) <a href="/src/target/league.go?s=2759:2834#L60">LeagueByMaster</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) LeagueByMaster(queue <a href="/pkg/builtin/#string">string</a>) (league <a href="#League">League</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.LeagueBySummoner">func (*RiotAPI) <a href="/src/target/league.go?s=1246:1342#L29">LeagueBySummoner</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) LeagueBySummoner(summonerID <a href="/pkg/builtin/#string">string</a>) (leagues map[<a href="/pkg/builtin/#string">string</a>][]<a href="#League">League</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.LeagueByTeam">func (*RiotAPI) <a href="/src/target/league.go?s=1867:1955#L41">LeagueByTeam</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) LeagueByTeam(teamID <a href="/pkg/builtin/#string">string</a>) (leagues map[<a href="/pkg/builtin/#string">string</a>][]<a href="#League">League</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.LeagueEntryBySummoner">func (*RiotAPI) <a href="/src/target/league.go?s=1551:1652#L35">LeagueEntryBySummoner</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) LeagueEntryBySummoner(summonerID <a href="/pkg/builtin/#string">string</a>) (leagues map[<a href="/pkg/builtin/#string">string</a>][]<a href="#League">League</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.LeagueEntryByTeam">func (*RiotAPI) <a href="/src/target/league.go?s=2156:2249#L47">LeagueEntryByTeam</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) LeagueEntryByTeam(teamID <a href="/pkg/builtin/#string">string</a>) (leagues map[<a href="/pkg/builtin/#string">string</a>][]<a href="#League">League</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.MasteriesByID">func (*RiotAPI) <a href="/src/target/summoner.go?s=2170:2268#L62">MasteriesByID</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) MasteriesByID(summonerID <a href="/pkg/builtin/#string">string</a>) (masteries map[<a href="/pkg/builtin/#string">string</a>]masteryBook, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.MatchByMatchID">func (*RiotAPI) <a href="/src/target/match.go?s=12287:12392#L229">MatchByMatchID</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) MatchByMatchID(matchID <a href="/pkg/builtin/#string">string</a>, includeTimeline <a href="/pkg/builtin/#string">string</a>) (match <a href="#MatchDetail">MatchDetail</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.MatchHistoryBySummonerID">func (*RiotAPI) <a href="/src/target/matchhistory.go?s=1021:1130#L18">MatchHistoryBySummonerID</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) MatchHistoryBySummonerID(options ...interface{}) (playerHistory <a href="#PlayerHistory">PlayerHistory</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.RankedStatsBySummoner">func (*RiotAPI) <a href="/src/target/stats.go?s=5066:5172#L91">RankedStatsBySummoner</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) RankedStatsBySummoner(summonerID <a href="/pkg/builtin/#string">string</a>, season <a href="/pkg/builtin/#string">string</a>) (stats <a href="#RankedStats">RankedStats</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.RecentGameBySummoner">func (*RiotAPI) <a href="/src/target/game.go?s=6086:6174#L110">RecentGameBySummoner</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) RecentGameBySummoner(summonerID <a href="/pkg/builtin/#string">string</a>) (games gamesList, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.RunesByID">func (*RiotAPI) <a href="/src/target/summoner.go?s=1880:1967#L56">RunesByID</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) RunesByID(summonerID <a href="/pkg/builtin/#string">string</a>) (runes map[<a href="/pkg/builtin/#string">string</a>]runeBook, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.SetLongRateLimit">func (*RiotAPI) <a href="/src/target/goriot.go?s=3803:3879#L122">SetLongRateLimit</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) SetLongRateLimit(numrequests <a href="/pkg/builtin/#int">int</a>, pertime <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>)</pre>
				<p>
SetLongRateLimit allows a custom rate limit to be set. For at the time of this writing the default
for a development API key is 500 requests every 10 minutes
</p>

				
				
				
			
				
				<h3 id="RiotAPI.SetSmallRateLimit">func (*RiotAPI) <a href="/src/target/goriot.go?s=3388:3465#L112">SetSmallRateLimit</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) SetSmallRateLimit(numrequests <a href="/pkg/builtin/#int">int</a>, pertime <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>)</pre>
				<p>
SetSmallRateLimit allows a custom rate limit to be set. For at the time of this writing the default
for a development API key is 10 requests every 10 seconds
</p>

				
				
				
			
				
				<h3 id="RiotAPI.StatSummariesBySummoner">func (*RiotAPI) <a href="/src/target/stats.go?s=4698:4817#L84">StatSummariesBySummoner</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) StatSummariesBySummoner(summonerID <a href="/pkg/builtin/#string">string</a>, season <a href="/pkg/builtin/#string">string</a>) (stats playerStatsSummaryList, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.SummonerByID">func (*RiotAPI) <a href="/src/target/summoner.go?s=1585:1679#L50">SummonerByID</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) SummonerByID(summonerID <a href="/pkg/builtin/#string">string</a>) (summoners map[<a href="/pkg/builtin/#string">string</a>]<a href="#Summoner">Summoner</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.SummonerByName">func (*RiotAPI) <a href="/src/target/summoner.go?s=1290:1381#L44">SummonerByName</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) SummonerByName(names <a href="/pkg/builtin/#string">string</a>) (summoners map[<a href="/pkg/builtin/#string">string</a>]<a href="#Summoner">Summoner</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.TeamBySummonerID">func (*RiotAPI) <a href="/src/target/team.go?s=2499:2591#L52">TeamBySummonerID</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) TeamBySummonerID(summonerID <a href="/pkg/builtin/#string">string</a>) (teams map[<a href="/pkg/builtin/#string">string</a>][]<a href="#Team">Team</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
				
				<h3 id="RiotAPI.TeamByTeamID">func (*RiotAPI) <a href="/src/target/team.go?s=2796:2878#L58">TeamByTeamID</a></h3>
				<pre>func (api *<a href="#RiotAPI">RiotAPI</a>) TeamByTeamID(teamID <a href="/pkg/builtin/#string">string</a>) (teams map[<a href="/pkg/builtin/#string">string</a>]<a href="#Team">Team</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
				
				
				
				
			
		
			
			
			<h2 id="RiotError">type <a href="/src/target/goriot.go?s=2703:2746#L86">RiotError</a></h2>
			<pre>type RiotError struct {
    StatusCode <a href="/pkg/builtin/#int">int</a>
}</pre>
			<p>
RiotError contains the http status code of the erro
</p>


			

			

			
			
			

			

			
				
				<h3 id="RiotError.Error">func (RiotError) <a href="/src/target/goriot.go?s=5465:5500#L186">Error</a></h3>
				<pre>func (err <a href="#RiotError">RiotError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>
				<p>
Error prints the error message for a RiotError
</p>

				
				
				
			
		
			
			
			<h2 id="Roster">type <a href="/src/target/team.go?s=1829:1949#L31">Roster</a></h2>
			<pre>type Roster struct {
    MemberList []<a href="#TeamMemberInfo">TeamMemberInfo</a> `json:&#34;memberList&#34;`
    OwnerID    <a href="/pkg/builtin/#int64">int64</a>            `json:&#34;ownerId&#34;`
}</pre>
			<p>
Roster represents the roster of a League of Legends ranked team
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="Rune">type <a href="/src/target/match.go?s=9615:9696#L157">Rune</a></h2>
			<pre>type Rune struct {
    Rank   <a href="/pkg/builtin/#int64">int64</a> `json:&#34;rank&#34;`
    RuneId <a href="/pkg/builtin/#int64">int64</a> `json:&#34;runeId&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="RuneCount">type <a href="/src/target/curentgame.go?s=1555:1642#L35">RuneCount</a></h2>
			<pre>type RuneCount struct {
    Count  <a href="/pkg/builtin/#int">int</a>   `json:&#34;count&#34;`
    RuneId <a href="/pkg/builtin/#int64">int64</a> `json:&#34;runeId&#34;`
}</pre>
			<p>
Mastery alreay exist
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="RunePage">type <a href="/src/target/summoner.go?s=684:851#L22">RunePage</a></h2>
			<pre>type RunePage struct {
    Current <a href="/pkg/builtin/#bool">bool</a>       `json:&#34;current&#34;`
    ID      <a href="/pkg/builtin/#int64">int64</a>      `json:&#34;id&#34;`
    Name    <a href="/pkg/builtin/#string">string</a>     `json:&#34;name&#34;`
    Slots   []<a href="#RuneSlot">RuneSlot</a> `json:&#34;slots&#34;`
}</pre>
			<p>
RunePage is a League of Legends rune page
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="RuneSlot">type <a href="/src/target/summoner.go?s=908:1003#L30">RuneSlot</a></h2>
			<pre>type RuneSlot struct {
    RuneId     <a href="/pkg/builtin/#int">int</a> `json:&#34;runeId&#34;`
    RuneSlotID <a href="/pkg/builtin/#int">int</a> `json:&#34;runeSlotId&#34;`
}</pre>
			<p>
RuneSlot is a slot for a Rune to go on a RunePage
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="Summoner">type <a href="/src/target/summoner.go?s=1052:1286#L36">Summoner</a></h2>
			<pre>type Summoner struct {
    ProfileIconID <a href="/pkg/builtin/#int">int</a>    `json:&#34;profileIconId&#34;`
    SummonerLevel <a href="/pkg/builtin/#int">int</a>    `json:&#34;summonerLevel&#34;`
    ID            <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;id&#34;`
    Name          <a href="/pkg/builtin/#string">string</a> `json:&#34;name&#34;`
    RevisionDate  <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;revisionDate&#34;`
}</pre>
			<p>
Summoner is a player of League of Legends
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="Team">type <a href="/src/target/team.go?s=84:1161#L1">Team</a></h2>
			<pre>type Team struct {
    CreateDate                    <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;createDate&#34;`
    FullID                        <a href="/pkg/builtin/#string">string</a>                `json:&#34;fullId&#34;`
    LastGameDate                  <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;lastGameDate&#34;`
    LastJoinDate                  <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;lastJoinDate&#34;`
    LastJoinedRankedTeamQueueDate <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;lastJoinedRankedTeamQueueDate&#34;`
    MatchHistory                  []<a href="#MatchHistorySummary">MatchHistorySummary</a> `json:&#34;matchHistory&#34;`
    ModifyDate                    <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;modifyDate&#34;`
    Name                          <a href="/pkg/builtin/#string">string</a>                `json:&#34;name&#34;`
    Roster                        <a href="#Roster">Roster</a>                `json:&#34;roster&#34;`
    SecondLastJoinDate            <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;secondLastJoinDate&#34;`
    Status                        <a href="/pkg/builtin/#string">string</a>                `json:&#34;status&#34;`
    Tag                           <a href="/pkg/builtin/#string">string</a>                `json:&#34;tag&#34;`
    TeamStatDetails               []<a href="#TeamStatDetail">TeamStatDetail</a>      `json:&#34;teamStatDetails&#34;`
    ThirdJoinDate                 <a href="/pkg/builtin/#int64">int64</a>                 `json:&#34;thirdLastJoinDate&#34;`
}</pre>
			<p>
Team is a League of Legends Ranked Team
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="TeamGameDetails">type <a href="/src/target/match.go?s=1902:2782#L33">TeamGameDetails</a></h2>
			<pre>type TeamGameDetails struct {
    Bans                 []<a href="#BannedChampion">BannedChampion</a> `json:&#34;bans&#34;`
    BaronKills           <a href="/pkg/builtin/#int">int</a>              `json:&#34;baronKills&#34;`
    DominionVictoryScore <a href="/pkg/builtin/#int64">int64</a>            `json:&#34;dominionVictoryScore&#34;`
    DragonKills          <a href="/pkg/builtin/#int">int</a>              `json:&#34;dragonKills&#34;`
    FirstBaron           <a href="/pkg/builtin/#bool">bool</a>             `json:&#34;firstBaron&#34;`
    FirstBlood           <a href="/pkg/builtin/#bool">bool</a>             `json:&#34;firstBlood&#34;`
    FirstDragon          <a href="/pkg/builtin/#bool">bool</a>             `json:&#34;firstDragon&#34;`
    FirstInhibitor       <a href="/pkg/builtin/#bool">bool</a>             `json:&#34;firstInhibitor&#34;`
    FirstTower           <a href="/pkg/builtin/#bool">bool</a>             `json:&#34;firstTower&#34;`
    InhibitorKills       <a href="/pkg/builtin/#int">int</a>              `json:&#34;inhibitorKills&#34;`
    TeamID               <a href="/pkg/builtin/#int">int</a>              `json:&#34;teamId&#34;`
    TowerKills           <a href="/pkg/builtin/#int">int</a>              `json:&#34;towerKills&#34;`
    VilemawKills         <a href="/pkg/builtin/#int">int</a>              `json:&#34;vilemawKills&#34;`
    Winner               <a href="/pkg/builtin/#bool">bool</a>             `json:&#34;winner&#34;`
}</pre>
			

			

			

			
			
			

			

			
		
			
			
			<h2 id="TeamMemberInfo">type <a href="/src/target/team.go?s=2031:2214#L37">TeamMemberInfo</a></h2>
			<pre>type TeamMemberInfo struct {
    InviteDate <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;inviteDate&#34;`
    JoinDate   <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;joinDate&#34;`
    PlayerID   <a href="/pkg/builtin/#int64">int64</a>  `json:&#34;playerId&#34;`
    Status     <a href="/pkg/builtin/#string">string</a> `json:&#34;status&#34;`
}</pre>
			<p>
TeamMemberInfo is the individual information for a player on a ranked team
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="TeamStatDetail">type <a href="/src/target/team.go?s=2272:2495#L45">TeamStatDetail</a></h2>
			<pre>type TeamStatDetail struct {
    AverageGamesPlayed <a href="/pkg/builtin/#int">int</a>    `json:&#34;averageGamesPlayed&#34;`
    Losses             <a href="/pkg/builtin/#int">int</a>    `json:&#34;losses&#34;`
    TeamStatType       <a href="/pkg/builtin/#string">string</a> `json:&#34;teamStatType&#34;`
    Wins               <a href="/pkg/builtin/#int">int</a>    `json:&#34;wins&#34;`
}</pre>
			<p>
TeamStatDetail is the statistics for a ranked team
</p>


			

			

			
			
			

			

			
		
			
			
			<h2 id="Timeline">type <a href="/src/target/match.go?s=2786:2898#L50">Timeline</a></h2>
			<pre>type Timeline struct {
    FrameInterval <a href="/pkg/builtin/#int64">int64</a>   `json:&#34;frameInterval&#34;`
    Frames        []<a href="#Frame">Frame</a> `json:&#34;frames&#34;`
}</pre>