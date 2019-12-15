package main

import (
	"strings"
)

type FaqCommand struct {
}

func (cmd FaqCommand) Init() {}

func (cmd FaqCommand) ShouldRespond(msg string) bool {
	keywords := []string{
		"question",
		"faq",
	}

	for _, kw := range keywords {
		if strings.Contains(strings.ToLower(msg), kw) {
			return true
		}
	}
	return false
}

func (cmd FaqCommand) Respond(msg string) string {
	return "Looking for the Hack Club Bank FAQs?\nIt's here 👉 bank.hackclub.com/faq :bank-hackclub:"
}
