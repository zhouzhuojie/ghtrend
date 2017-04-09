package main

import (
	"os"

	"github.com/jasonlvhit/gocron"
)

var (
	// config the gmail auth from env
	userEmail    = os.Getenv("userEmail")
	userPassword = os.Getenv("userPassword")

	// config which languages you are interested in
	// separated by ","
	// e.g. languages=go,python,javascript
	ghLanguages = os.Getenv("languages")
)

func main() {

	// send the first time when we run
	sendGithubTrendMail()

	// send every day at utc 3:30 am
	gocron.Every(1).Day().At("03:30").Do(sendGithubTrendMail)
	<-gocron.Start()
}

func sendGithubTrendMail() {
	html := CrawlGithubTrendingPages()
	SendGithubTrendMail(html)
}
