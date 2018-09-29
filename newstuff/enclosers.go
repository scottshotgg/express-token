package token

// EncloserMap holds all valid encloser tokens
var EncloserMap = map[string]Token{
	"(": Token{
		Type:   "R_PAREN",
		String: "(",
	},

	")": Token{
		Type:   "L_PAREN",
		String: ")",
	},

	"{": Token{
		Type:   "R_BRACE",
		String: "{",
	},

	"}": Token{
		Type:   "L_BRACE",
		String: "}",
	},

	"[": Token{
		Type:   "R_BRACKET",
		String: "[",
	},

	"]": Token{
		Type:   "L_BRACKET",
		String: "]",
	},

	"<": Token{
		Type:   "LTHAN",
		String: "<",
	},

	">": Token{
		Type:   "RTHAN",
		String: ">",
	},

	"`": Token{
		Type:   "GRAVE",
		String: "`",
	},

	"~": Token{
		Type:   "TILDE",
		String: "~",
	},

	"'": Token{
		Type:   "SQUOTE",
		String: "'",
	},

	"\"": Token{
		Type:   "DQUOTE",
		String: "\"",
	},

	"@": Token{
		Type:   "AT",
		String: "@",
	},
}
