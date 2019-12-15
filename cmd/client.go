package main

import (
	"time"
)

type Client struct {
	startedAt time.Time

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
