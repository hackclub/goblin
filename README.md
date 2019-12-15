# Goblin

Goblin is a Slack bot that integrates with Hack Club Bank 🐲

## Commands

- `@goblin help`: Display help text.
- `@goblin hello`: Make some witty comment about money in the cloud.
- `@goblin stats`: Reports HCB stats.
- `@goblin apply`: Apply to Hack Club Bank! Shows the application link.
- `@goblin faq`: Shows a link to the FAQ page.
- `@goblinn bug`: Report a bug! Redirects user to the #bank-counter channel.

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
