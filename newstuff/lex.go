package token

// Lexemes ...
var (
	Lexemes = []string{
		"var",
		"int",
		"float",
		"string",
		"bool",
		"char",
		"object",

		"select",
		"SELECT",
		"FROM",
		"WHERE",
		":",
		"=",
		"+",
		"-",
		"*",
		"/",
		"(",
		")",
		"{",
		"}",
		"[",
		"]",
		"\"",
		"'",
		";",
		",",
		"#",
		"!",
		"<",
		">",
		"@",
		// "„",
		" ",
		"\n",
	}

	// LexemeMap is used for holding the lexemes that will be used to identify tokens in the lexer
	LexemeMap = map[string]Token{}
)
