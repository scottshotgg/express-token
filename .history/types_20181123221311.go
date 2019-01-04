package token

// TypeMap holds all defined type tokens
var TypeMap = map[string]Token{
	VarType: Token{
		Type: Type,
		Value: Value{
			Type:   VarType,
			String: VarType,
		},
	},
	IntType: Token{
		Type: Type,
		Value: Value{
			Type:   IntType,
			String: IntType,
		},
	},
	FloatType: Token{
		Type: Type,
		Value: Value{
			Type:   FloatType,
			String: FloatType,
		},
	},
	CharType: Token{
		Type: Type,
		Value: Value{
			Type:   CharType,
			String: CharType,
		},
	},
	StringType: Token{
		Type: Type,
		Value: Value{
			Type:   StringType,
			String: StringType,
		},
	},
	BoolType: Token{
		Type: Type,
		Value: Value{
			Type:   BoolType,
			String: BoolType,
		},
	},
	ObjectType: Token{
		Type: Type,
		Value: Value{
			Type:   ObjectType,
			String: ObjectType,
		},
	},
	// StructType: Token{
	// 	Type: Type,
	// 	Value: Value{
	// 		Type:   StructType,
	// 		String: StructType,
	// 	},
	// },
	ArrayType: Token{
		Type: Type,
		Value: Value{
			Type:   ArrayType,
			String: ArrayType,
		},
	},
	IntArrayType: Token{
		Type: Type,
		Value: Value{
			Type:   ArrayType,
			Acting: IntType,
			String: ArrayType,
		},
	},
	StringArrayType: Token{
		Type: Type,
		Value: Value{
			Type:   ArrayType,
			Acting: StringType,
			String: ArrayType,
		},
	},
}
