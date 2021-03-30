package stereotype

import "github.com/golangee/src/ast"

type Struct struct {
	obj *ast.Struct
}

func StructFrom(obj *ast.Struct) Struct {
	return Struct{obj: obj}
}

// SetConfiguration marks this struct as a public configuration object. It provides environmental and program flags.
func (s Struct) SetIsConfiguration(isPublicConfig bool) Struct {
	s.obj.PutValue(kConfiguration, isPublicConfig)
	return s
}

// IsConfiguration returns only true, if this struct shall be used as a configurational object.
func (s Struct) IsConfiguration() bool {
	v := s.obj.Value(kConfiguration)
	if f, ok := v.(bool); ok {
		return f
	}

	return false
}

// SetIsDatabaseConfiguration marks this struct as a public configuration object. It provides environmental and program flags.
func (s Struct) SetIsDatabaseConfiguration(isDbConfig bool) Struct {
	s.obj.PutValue(kDBConfiguration, isDbConfig)
	return s
}

// IsDBConfiguration returns only true, if this struct shall be used as a configurational object.
func (s Struct) IsDatabaseConfiguration() bool {
	v := s.obj.Value(kDBConfiguration)
	if f, ok := v.(bool); ok {
		return f
	}

	return false
}

// SetMySQLRelated marks this struct as mysql related.
func (s Struct) SetMySQLRelated(isDbConfig bool) Struct {
	s.obj.PutValue(kMySQLRelated, isDbConfig)
	return s
}

// MySQLRelated returns only true, if this struct is mysql related.
func (s Struct) MySQLRelated() bool {
	v := s.obj.Value(kMySQLRelated)
	if f, ok := v.(bool); ok {
		return f
	}

	return false
}

// SetSQLTableName sets an sql table name
func (s Struct) SetSQLTableName(tableName string) Struct {
	s.obj.PutValue(kSQLTableName, tableName)
	return s
}

// SQLTableName extracts the sql table name.
func (s Struct) SQLTableName() (string, bool) {
	val := s.obj.Value(kSQLTableName)
	if val == nil {
		return "", false
	}

	return val.(string), true
}
