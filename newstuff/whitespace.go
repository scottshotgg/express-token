package token

// WhitespaceMap holds all defined Whitespace tokens
var WhitespaceMap = map[string]Token{
	" ": Token{
		Type:   "SPACE",
		String: " ",
	},

	"\t": Token{
		Type:   "TAB",
		String: "\t",
	},

	"\n": Token{
		Type:   "NEWLINE",
		String: "\n",
	},
}
