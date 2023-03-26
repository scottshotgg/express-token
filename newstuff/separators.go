package token

// SeparatorMap holds all defined statement separators
var SeparatorMap = map[string]Token{
	",": {
		Type:   "COMMA",
		String: ",",
	},

	";": {
		Type:   "SEMICOLON",
		String: ";",
	},
}
