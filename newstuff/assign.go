package token

// AssignMap holds every assignment operator
var AssignMap = map[string]Token{
	"=": Token{
		Type:   "ASSIGN",
		String: "=",
	},

	":": Token{
		Type:   "SET",
		String: ":",
	},

	":=": Token{
		Type:   "INIT",
		String: ":=",
	},
}
