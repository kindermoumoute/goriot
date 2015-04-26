package goriot

import (
	"net/http"
	"io/ioutil"
	"log"
)

var APIkey = ""


func GetInfosSummoner(summoner string) string {
	var retour string = ""
	res, err := http.Get("https://euw.api.pvp.net/api/lol/euw/v1.4/summoner/by-name/" + summoner + "?api_key=" + APIkey)
	if err != nil {
		log.Fatal(err)
	}
	infos, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	retour = string(infos)

	return retour
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