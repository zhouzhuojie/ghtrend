package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/zhouzhuojie/ghtrend/crawl"
	"github.com/zhouzhuojie/ghtrend/mail"
)

func main() {

	// send the first time when we run
	scheduledSendGithubTrendMail()

	// send every day at 5:30 am
	gocron.Every(1).Day().At("05:30").Do(scheduledSendGithubTrendMail)

	_, time := gocron.NextRun()
	fmt.Println(time)

	<-gocron.Start()
}

func scheduledSendGithubTrendMail() {
	html := crawl.CrawlGithubTrendingPages()
	err := mail.SendGithubTrendMail(html)
	if err != nil {
		fmt.Println(err)
	}
}
