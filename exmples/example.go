package main

import (
	"github.com/m-nakada/slackposter"
)

func main() {
	config := slackposter.Config{
		"yourchannel",
		":ghost:",
		"Ghost",
		"https://hooks.slack.com/services/xxxx/xxx/xxx",
	}

	slack := slackposter.NewSlack(config)
	slack.PostMessage("Hello world!")
}
