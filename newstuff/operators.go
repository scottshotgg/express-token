package token

// OperatorMap holds all defined operator tokens
var OperatorMap = map[string]Token{
	"+": Token{
		Kind:   "ADD",
		String: "+",
	},

	"-": Token{
		Kind:   "SUB",
		String: "-",
	},

	"*": Token{
		Kind:   "MULT",
		String: "*",
	},

	"/": Token{
		Kind:   "DIV",
		String: "/",
	},

	// "%": Token{
	// 	Type:   PriOp,
	// 	Kind:   "mod",
	// 	String: "%",
	// },

	// "^": Token{
	// 	Type:   PriOp,
	// 	Kind:   "exp",
	// 	String: "^",
	// },

	// "!": Token{
	// 	Type:   Bang,
	// 	Kind:   "bang",
	// 	String: "!",
	// },

	// "?": Token{
	// 	Type:   QuestionMark,
	// 	Kind:   "qm",
	// 	String: "!",
	// },

	// "_": Token{
	// 	Type:   Underscore,
	// 	Kind:   "underscore",
	// 	String: "_",
	// },

	// // FIXME: DOLLA DOLLA BILLS YALL: define this
	// "$": Token{
	// 	Type:   DDBY,
	// 	Kind:   "ddby",
	// 	String: "$",
	// },

	// "&": Token{
	// 	Type:   Ampersand,
	// 	Kind:   "op_3",
	// 	String: "&",
	// },

	// "|": Token{
	// 	Type:   Pipe,
	// 	Kind:   "op_3",
	// 	String: "|",
	// },

	// "#": Token{
	// 	Type:   Hash,
	// 	Kind:   "op_3",
	// 	String: "#",
	// },

	".": Token{
		Kind:   "PERIOD",
		String: ".",
	},

	// "==": Token{
	// 	Type:   IsEqual,
	// 	Kind:   "is_equal",
	// 	String: "==",
	// },

	// // Increment
	// "++": Token{
	// 	Type:   Increment,
	// 	Kind:   "increment",
	// 	String: "++",
	// },

	// // VECTOR OPERANDS
	// ".+": Token{
	// 	Type:   "VEC_ADD",
	// 	Kind:   "op_3",
	// 	String: ".+",
	// },

	// ".-": Token{
	// 	Type:   "VEC_SUB",
	// 	Kind:   "op_4",
	// 	String: ".-",
	// },

	// ".*": Token{
	// 	Type:   "VEC_MULT",
	// 	Kind:   "op_3",
	// 	String: ".*",
	// },

	// "./": Token{
	// 	Type:   "VEC_DIV",
	// 	Kind:   "op_3",
	// 	String: "./",
	// },
}
