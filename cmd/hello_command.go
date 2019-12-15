package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type HelloCommand struct {
}

func (cmd HelloCommand) Init() {
	rand.Seed(time.Now().Unix())
}

func (cmd HelloCommand) ShouldRespond(msg string) bool {
	hellos := []string{
		"hello",
		"hi",
		"hey",
	}

	for _, greeting := range hellos {
		if strings.Contains(strings.ToLower(msg), greeting) {
			return true
		}
	}
	return false
}

func (cmd HelloCommand) Respond(msg string) string {
	autonyms := []string{
		"the hivemind",
		"a cloud full of money",
		"Hack Club's pot of gold",
		"a sentient stack of dollars",
		"the Hack Club Federal Reserve",
		"a money-crazed virus 🤑",
		"a cloud raining money",
		"a pile of money in the cloud",
		"Hack Club Smoothmunny",
		"Hack Club ezBUCKS",
		"Hack Club Money Bucket",
		"a mattress stuffed with 100 dollar bills",
		"Hack Club Dollaringos",
		"the Hack Club Dolla Store",
	}

	return fmt.Sprintf("Hey there! I'm Goblin from Bank, %s", autonyms[rand.Intn(len(autonyms))])
}
