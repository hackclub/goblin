package main

// Generic command interface that all Goblin commands should conform to
type Command interface {
	Name() string
	Respond(string) string
}
