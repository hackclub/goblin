package main

// Generic command interface that all Goblin commands should conform to
type Command interface {
	Init()
	ShouldRespond(string) bool
	Respond(string) string
}
