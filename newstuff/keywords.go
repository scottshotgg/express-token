package token

// KeywordMap is a map of all the keywords
var KeywordMap = map[string]Token{
	"let": {
		Type:   "LET",
		String: "let",
	},

	"select": {
		Type:   "SELECT",
		String: "select",
	},

	"for": {
		Type:   "FOR",
		String: "for",
	},

	"if": {
		Type:   "IF",
		String: "if",
	},

	"in": {
		Type:   "IN",
		String: "in",
	},

	"of": {
		Type:   "OF",
		String: "of",
	},

	"function": {
		Type:   "FUNCTION",
		String: "function",
	},

	"func": {
		Type:   "FUNC",
		String: "func",
	},

	"fn": {
		Type:   "FN",
		String: "fn",
	},

	"return": {
		Type:   "RETURN",
		String: "return",
	},

	"onexit": {
		Type:   "ONEXIT",
		String: "onexit",
	},

	"onreturn": {
		Type:   "ONRETURN",
		String: "onreturn",
	},

	"onleave": {
		Type:   "ONLEAVE",
		String: "onleave",
	},

	"defer": {
		Type:   "DEFER",
		String: "defer",
	},
}
