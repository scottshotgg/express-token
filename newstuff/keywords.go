package token

// KeywordMap is a map of all the keywords
var KeywordMap = map[string]Token{
	"let": Token{
		Type:   "LET",
		String: "let",
	},

	"select": Token{
		Type:   "SELECT",
		String: "select",
	},

	"for": Token{
		Type:   "FOR",
		String: "for",
	},

	"if": Token{
		Type:   "IF",
		String: "if",
	},

	"in": Token{
		Type:   "IN",
		String: "in",
	},

	"of": Token{
		Type:   "OF",
		String: "of",
	},

	"function": Token{
		Type:   "FUNCTION",
		String: "function",
	},

	"func": Token{
		Type:   "FUNC",
		String: "func",
	},

	"fn": Token{
		Type:   "FN",
		String: "fn",
	},

	"return": Token{
		Type:   "RETURN",
		String: "return",
	},

	"onexit": Token{
		Type:   "ONEXIT",
		String: "onexit",
	},

	"onreturn": Token{
		Type:   "ONRETURN",
		String: "onreturn",
	},

	"onleave": Token{
		Type:   "ONLEAVE",
		String: "onleave",
	},

	"defer": Token{
		Type:   "DEFER",
		String: "defer",
	},
}
