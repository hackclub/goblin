package goblin

import (
	"time"
)

// Client represents a Goblin bot instance, that'll handle
// all incoming messages.
type Client struct {
	startedAt time.Time

	// Commands is a list of registered Goblin commands
	Commands []Command
}

func NewClient() *Client {
	c := &Client{
		startedAt: time.Now(),
		Commands: []Command{
			HelpCommand{},
			StatsCommand{},
			ApplyCoommand{},
			FaqCommand{},
			BugCommand{},
			// Add new commands here
		},
	}

	for _, cmd := range c.Commands {
		cmd.Init()
	}

	return c
}

func (c Client) Uptime() time.Duration {
	return time.Since(c.startedAt)
}

func (c Client) Respond(msg string) string {
	for _, cmd := range c.Commands {
		if cmd.ShouldRespond(msg) {
			return cmd.Respond(msg)
		}
	}

	return HelloCommand{}.Respond(msg)
}
