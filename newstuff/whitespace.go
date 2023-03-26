package token

// WhitespaceMap holds all defined Whitespace tokens
var WhitespaceMap = map[string]Token{
	" ": {
		Type:   "SPACE",
		String: " ",
	},

	"\t": {
		Type:   "TAB",
		String: "\t",
	},

	"\n": {
		Type:   "NEWLINE",
		String: "\n",
	},
}
