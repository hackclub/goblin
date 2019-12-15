package main

import (
	"strings"
)

type BugCommand struct {
}

func (cmd BugCommand) Init() {}

func (cmd BugCommand) ShouldRespond(msg string) bool {
	keywords := []string{
		"bug",
		"problem",
		"broke",
	}

	for _, kw := range keywords {
		if strings.Contains(strings.ToLower(msg), kw) {
			return true
		}
	}
	return false
}

func (cmd BugCommand) Respond(msg string) string {
	return "Did you find a bug in Bank? Did something break?\n" +
		"We want to fix it! Go tell *@thesephist* or report it at #bank-counter"
}
