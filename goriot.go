package goriot

/*
TODO :
	Faire des structs summoner, game, etc..
	Généralisé les liens de l'API ?
*/
import (
	"io/ioutil"
	"log"
	"net/http"
)

type RiotAPI struct {
	APIkey string
	Region string
}

func Get(key string, region string) *RiotAPI {
	return &RiotAPI{
		APIkey: key,
		Region: region,
	}
}

func (api *RiotAPI) GetInfosSummoner(summoner string) (retour string) {
	res, err := http.Get("https://" + api.Region + ".api.pvp.net/api/lol/" + api.Region + "/v1.4/" + "summoner/by-name/" + summoner + "?api_key=" + api.APIkey)
	if err != nil {
		log.Fatal(err)
	}
	infos, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	retour = string(infos)

	return
}

// func bonSummoner(summoner string) string {

// }
/* `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`
*/
