package parser

// SQL contains sql independent declaration, but the generators are dialect specific.
type SQL struct {
	Database   String               `"database" "=" @@`
	Implements []*SQLImplementation `@@+`
}

type SQLImplementation struct {
	Type Path `"impl" @@ "{"`

	Configure []*FieldWithDefault `("configure" "{" @@* "}")?`
	Inject    []*Field            `("inject" "{" @@* "}")?`
	Private   []*Field            `("private" "{" @@* "}")?`

	SQLFunc []*SQLFunc ` @@* "}"`
}

type SQLFunc struct {
	Name Ident           `@@`
	SQL  String          `@@`
	In   []SQLFuncInVar  `( "(" @@ ("," @@)* ")" )?`
	Out  []SQLFuncOutVar `( "=>" "(" @@ ("," @@)* ")" )?`
}

type SQLFuncInVar struct {
	Selector []LooperIdent `( @@ ("." @@)* )?`
	IsLooper bool          `(@SliceLooper)?`
}

type LooperIdent struct {
	Ident    Ident `@@`
	IsLooper bool  `(@SliceLooper)?`
}

type SQLFuncOutVar struct {
	Ident Ident `"." @@?`
}
