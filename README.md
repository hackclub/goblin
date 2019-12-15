# Goblin

Goblin is a Slack bot that integrates with Hack Club Bank 🐲

## Commands

- `/goblin hello`: Make some witty comment
- `/goblin stats`: Reports HCB stats
- `/goblin apply`: Apply to Hack Club Bank!
- `/goblin faq`: Go to the FAQ page.

## Build

>Currently, there's a bug in goblin's dependency nlopes/slack caused by Slack releasing an API update
that contains a new `rich_text` field after their rich text editor rollout. Until the mainline of the
slack Go library fixes this, build goblin with

```sh
go get -u github.com/nlopes/slack@d06c2a2b3249b44a9c5dee8485f5a87497beb9ea
```

>to build with a version of nlopes/slack that works with the new rich text API.

Build with `go build cmd/goblin.go`.

Run with `./goblin`. Goblin should automatically connect to Slack and start handling messages.
