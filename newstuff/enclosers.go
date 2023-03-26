package token

// EncloserMap holds all valid encloser tokens
var EncloserMap = map[string]Token{
	"(": {
		Type:   "R_PAREN",
		String: "(",
	},

	")": {
		Type:   "L_PAREN",
		String: ")",
	},

	"{": {
		Type:   "R_BRACE",
		String: "{",
	},

	"}": {
		Type:   "L_BRACE",
		String: "}",
	},

	"[": {
		Type:   "R_BRACKET",
		String: "[",
	},

	"]": {
		Type:   "L_BRACKET",
		String: "]",
	},

	"<": {
		Type:   "LTHAN",
		String: "<",
	},

	">": {
		Type:   "RTHAN",
		String: ">",
	},

	"`": {
		Type:   "GRAVE",
		String: "`",
	},

	"~": {
		Type:   "TILDE",
		String: "~",
	},

	"'": {
		Type:   "SQUOTE",
		String: "'",
	},

	"\"": {
		Type:   "DQUOTE",
		String: "\"",
	},

	"@": {
		Type:   "AT",
		String: "@",
	},
}
