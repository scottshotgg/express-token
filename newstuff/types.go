package token

// TypeMap holds all defined type tokens
var TypeMap = map[string]Token{
	VarType: Token{
		Type:   VarType,
		String: VarType,
	},

	IntType: Token{
		Type:   IntType,
		String: IntType,
	},

	FloatType: Token{
		Type:   FloatType,
		String: FloatType,
	},

	CharType: Token{
		Type:   CharType,
		String: CharType,
	},

	StringType: Token{
		Type:   StringType,
		String: StringType,
	},

	BoolType: Token{
		Type:   BoolType,
		String: BoolType,
	},

	ObjectType: Token{
		Type:   ObjectType,
		String: ObjectType,
	},

	StructType: Token{
		Type:   StructType,
		String: StructType,
	},

	ArrayType: Token{
		Type:   ArrayType,
		String: ArrayType,
	},
}
