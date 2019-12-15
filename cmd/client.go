package main

import (
	"strings"
	"time"
)

type Client struct {
	startedAt time.Time

	Commands []Command
}

func NewClient() *Client {
	return &Client{
		startedAt: time.Now(),
		Commands: []Command{
			StatsCommand{},
		},
	}
}

func (c Client) Uptime() time.Duration {
	return time.Since(c.startedAt)
}

func (c Client) Respond(msg string) string {
	for _, cmd := range c.Commands {
		lower := strings.ToLower(msg)
		if strings.Contains(lower, cmd.Name()) {
			return cmd.Respond(msg)
		}
	}

	return "Gobble gobble! Not sure what that means..."
}
