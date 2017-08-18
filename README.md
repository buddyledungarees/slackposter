# skackposter
Post a payload to your Slack Incoming Webhook.

## Usage

You can send a message.

```go
package main

import (
	"github.com/mnkd/slackposter"
)

func main() {
    config := slackposter.Config{
        "mn",
        ":ghost:",
        "Ghost",
        "https://hooks.slack.com/services/xxxx/xxx/xxx",
    }

    slack := slackposter.NewSlack(config)
    slack.PostMessage("Hello world!")
}
```

![message](examples/message.png)

You can send a customized message (payload).
Please refer to [examples/example_payload.go](https://github.com/mnkd/slackposter/blob/master/examples/example_payload.go)

![payload](examples/payload.png)

## Example Apps
* [github-status](https://github.com/mnkd/github-status) — Notify GitHub Site Status to Slack incoming webhook.
* [prnotify](https://github.com/mnkd/prnotify) — Notify GitHub pull requests to Slack incoming webhook.
* [qiiotd](https://github.com/mnkd/qiiotd) — Qiita:Team (Qiita) の n 年前の今日の記事を Slack の Incoming webhook へ post する。
