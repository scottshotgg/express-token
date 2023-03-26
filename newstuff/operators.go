package token

// OperatorMap holds all defined operator tokens
var OperatorMap = map[string]Token{
	"+": {
		Kind:   "ADD",
		String: "+",
	},

	"-": {
		Kind:   "SUB",
		String: "-",
	},

	"*": {
		Kind:   "MULT",
		String: "*",
	},

	"/": {
		Kind:   "DIV",
		String: "/",
	},

	// "%": {
	// 	Type:   PriOp,
	// 	Kind:   "mod",
	// 	String: "%",
	// },

	// "^": {
	// 	Type:   PriOp,
	// 	Kind:   "exp",
	// 	String: "^",
	// },

	// "!": {
	// 	Type:   Bang,
	// 	Kind:   "bang",
	// 	String: "!",
	// },

	// "?": {
	// 	Type:   QuestionMark,
	// 	Kind:   "qm",
	// 	String: "!",
	// },

	// "_": {
	// 	Type:   Underscore,
	// 	Kind:   "underscore",
	// 	String: "_",
	// },

	// // FIXME: DOLLA DOLLA BILLS YALL: define this
	// "$": {
	// 	Type:   DDBY,
	// 	Kind:   "ddby",
	// 	String: "$",
	// },

	// "&": {
	// 	Type:   Ampersand,
	// 	Kind:   "op_3",
	// 	String: "&",
	// },

	// "|": {
	// 	Type:   Pipe,
	// 	Kind:   "op_3",
	// 	String: "|",
	// },

	// "#": {
	// 	Type:   Hash,
	// 	Kind:   "op_3",
	// 	String: "#",
	// },

	".": {
		Kind:   "PERIOD",
		String: ".",
	},

	// "==": {
	// 	Type:   IsEqual,
	// 	Kind:   "is_equal",
	// 	String: "==",
	// },

	// // Increment
	// "++": {
	// 	Type:   Increment,
	// 	Kind:   "increment",
	// 	String: "++",
	// },

	// // VECTOR OPERANDS
	// ".+": {
	// 	Type:   "VEC_ADD",
	// 	Kind:   "op_3",
	// 	String: ".+",
	// },

	// ".-": {
	// 	Type:   "VEC_SUB",
	// 	Kind:   "op_4",
	// 	String: ".-",
	// },

	// ".*": {
	// 	Type:   "VEC_MULT",
	// 	Kind:   "op_3",
	// 	String: ".*",
	// },

	// "./": {
	// 	Type:   "VEC_DIV",
	// 	Kind:   "op_3",
	// 	String: "./",
	// },
}
