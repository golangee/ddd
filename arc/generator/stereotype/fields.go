package stereotype

import "github.com/golangee/src/ast"

// Field contains all stereotype annotations for single field instances.
type Field struct {
	field *ast.Field
}

func FieldFrom(field *ast.Field) Field {
	return Field{field: field}
}

func (c Field) Unwrap() *ast.Field {
	return c.field
}

// SetEnvironmentVariable connects the given field with a specific name from the environment variable.
func (c Field) SetEnvironmentVariable(environmentVariable string) {
	c.field.PutValue(kEnvironmentVariable, environmentVariable)
}

// EnvironmentVariable extracts the environment variable.
func (c Field) EnvironmentVariable() (string, bool) {
	val := c.field.Value(kEnvironmentVariable)
	if val == nil {
		return "", false
	}

	return val.(string), true
}

// SetProgramFlagVariable connects the given field with a specific name from the program flag variable.
func (c Field) SetProgramFlagVariable(flagVariable string) {
	c.field.PutValue(kProgramFlagVariable, flagVariable)
}

// ProgramFlagVariable extracts the program flag variable.
func (c Field) ProgramFlagVariable() (string, bool) {
	val := c.field.Value(kProgramFlagVariable)
	if val == nil {
		return "", false
	}

	return val.(string), true
}

// SetSQLColumnName sets an sql column name.
func (c Field) SetSQLColumnName(tableName string) Field {
	c.field.PutValue(kSQLColumnName, tableName)
	return c
}

// SQLColumnName extracts the sql column name.
func (c Field) SQLColumnName() (string, bool) {
	val := c.field.Value(kSQLColumnName)
	if val == nil {
		return "", false
	}

	return val.(string), true
}
