package token

// SeparatorMap holds all defined statement separators
var SeparatorMap = map[string]Token{
	",": Token{
		Type:   "COMMA",
		String: ",",
	},

	";": Token{
		Type:   "SEMICOLON",
		String: ";",
	},
}
