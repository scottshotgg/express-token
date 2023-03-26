package token

// AssignMap holds every assignment operator
var AssignMap = map[string]Token{
	"=": {
		Type:   "ASSIGN",
		String: "=",
	},

	":": {
		Type:   "SET",
		String: ":",
	},

	":=": {
		Type:   "INIT",
		String: ":=",
	},
}
