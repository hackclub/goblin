package goblin

import (
	"strings"
)

type HelpCommand struct {
}

func (cmd HelpCommand) Init() {}

func (cmd HelpCommand) ShouldRespond(msg string) bool {
	return strings.Contains(strings.ToLower(msg), "help")
}

func (cmd HelpCommand) Respond(msg string) string {
	lines := []string{
		"*Hi! I'm Goblin, the Hack Club Bank bot*",
		"Here's what you can ask me to do:",
		"- `help`: show this help message",
		"- `hello` or `hi`: I'll say hello back!",
		"- `stats`: show a summary of important Bank stats",
		"- `apply` or `join`: want to sign up for Hack Club Bank? I'll show you to the application.",
		"- `faq` or `question`: I'll show you to the Bank FAQ page",
		"- `bug` or `broken`: did something break on Hack Club Bank? I'll show you who to talk to",
		"",
		"Have another question I can't answer? You can find my full source code at https://github.com/hackclub/goblin :computer:",
	}

	return strings.Join(lines, "\n")
}
