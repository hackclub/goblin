> 2020-08-25: Goblin is no longer being maintained, supported, or developed by Hack Club HQ.

***
 
# Goblin

Goblin is a Slack bot that integrates with Hack Club Bank 🐲

## Commands

- `@goblin help`: Display help text.
- `@goblin hello`: Make some witty comment about money in the cloud.
- `@goblin stats`: Reports HCB stats.
- `@goblin apply`: Apply to Hack Club Bank! Shows the application link.
- `@goblin faq`: Shows a link to the FAQ page.
- `@goblin bug`: Report a bug! Redirects user to the #bank-counter channel.

## Build

_Currently, there's a bug in goblin's dependency nlopes/slack caused by Slack releasing an API update
that contains a new `rich_text` field after their rich text editor rollout. Until the mainline of the
slack Go library fixes this, build goblin with_

```sh
go get -u github.com/nlopes/slack@d06c2a2b3249b44a9c5dee8485f5a87497beb9ea
```

_...to build with a version of nlopes/slack that works with the new rich text API._

Build with `go build cmd/goblin.go`.

Run with `./goblin`. Goblin should automatically connect to Slack and start handling messages!

## Deploy

We deploy goblin as a systemd service on Linux:

1. Create a new user `goblin` on Linux with `sudo adduser goblin`
2. Build the executable and place it at `/home/goblin/goblin`
3. Copy over the systemd service file `goblin.service` to `/etc/systemd/system/goblin.service` and replace the `your api token` with your Slack API Bot OAuth token
4. Start the service with `sudo systemctl start goblin`

## Add a command

To add a new command to Goblin:

1. Add a new file that'll define that command. Let's say you're adding a `goodbye` command to say
goodbye to Slack users. First, add `pkg/goblin/goodbye_command.go`.
2. In that file, define a command that conforms to the `gobline.Command` interface defined in
`pkg/goblin/command.go`. You can base your command on `help_command.go`, which is a pretty bare-bones command
that knows how to respond to a particular message and generate a response.
3. Add your command to the client in `pkg/goblin/client.go`. Add a new line where it says
`// Add new commands here` to create a new instance of your command.
4. Add a line or two about the command you just added in both `help_command.go` and `README.md` to help
users discover your new command.
