package goblin

// Generic command interface that all Goblin commands should conform to
type Command interface {
	// Any global initialization work, like seeding an RNG, takes place
	// in Init
	Init()

	// Returns true if this command should be the one handling
	// a given message to Goblin. Usually, this just performs
	// a test to see if the message contains keywords or matches
	// a regular expression.
	ShouldRespond(string) bool

	// Given a message, generate a response for Goblin.
	Respond(string) string
}
