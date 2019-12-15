package goblin

import (
	"strings"
)

type ApplyCoommand struct {
}

func (cmd ApplyCoommand) Init() {}

func (cmd ApplyCoommand) ShouldRespond(msg string) bool {
	keywords := []string{
		"apply",
		"join",
	}

	for _, kw := range keywords {
		if strings.Contains(strings.ToLower(msg), kw) {
			return true
		}
	}
	return false
}

func (cmd ApplyCoommand) Respond(msg string) string {
	return "You can sign up for Hack Club Bank at hackclub.com/bank#apply :bank-hackclub:"
}
