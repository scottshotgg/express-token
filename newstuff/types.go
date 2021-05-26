package token

// TypeMap holds all defined type tokens
var TypeMap = map[string]Token{
	VarType: {
		Type:   VarType,
		String: VarType,
	},

	IntType: {
		Type:   IntType,
		String: IntType,
	},

	FloatType: {
		Type:   FloatType,
		String: FloatType,
	},

	CharType: {
		Type:   CharType,
		String: CharType,
	},

	StringType: {
		Type:   StringType,
		String: StringType,
	},

	BoolType: {
		Type:   BoolType,
		String: BoolType,
	},

	ObjectType: {
		Type:   ObjectType,
		String: ObjectType,
	},

	StructType: {
		Type:   StructType,
		String: StructType,
	},

	ArrayType: {
		Type:   ArrayType,
		String: ArrayType,
	},
}
