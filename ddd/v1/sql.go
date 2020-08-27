package ddd

import "time"

// A MySQLLayerSpec represents a stereotyped IMPLEMENTATION Layer.
type MySQLLayerSpec struct {
	name        string
	description string
	migrations  []*MigrationSpec
	genSpecs    []*GenSpec
}

// Name of the layer.
func (u *MySQLLayerSpec) Name() string {
	return u.name
}

// Description of the layer.
func (u *MySQLLayerSpec) Description() string {
	return u.description
}

// Stereotype of the layer.
func (u *MySQLLayerSpec) Stereotype() Stereotype {
	return IMPLEMENTATION
}

// MySQL is a factory for a MySQLLayerSpec. An implementation can only ever import Core API.
func MySQL(migrations []*MigrationSpec, genSpecs []*GenSpec) *MySQLLayerSpec {
	return &MySQLLayerSpec{
		name:        "mysql",
		description: "Package mysql contains specific repository implementations for the mysql dialect.",
		migrations:  migrations,
		genSpecs:    genSpecs,
	}
}

// A GenSpec refers to an interface and contains GenFuncSpec to map sql statements to method declarations and
// therefore also to input and output parameters.
type GenSpec struct {
	typeName TypeName
	funcs    []*GenFuncSpec
}

// From is a factory for a GenSpec which refers to an interface and contains GenFuncSpec to map
// sql statements to method declarations and
// therefore also to input and output parameters.
func From(typeName TypeName, funcs ...*GenFuncSpec) *GenSpec {
	return &GenSpec{
		typeName: typeName,
		funcs:    funcs,
	}
}

// Generate is a factory for a slice of GenSpec.
func Generate(genSpecs ...*GenSpec) []*GenSpec {
	return genSpecs
}

// A GenFuncSpec combines a function name (more precisely an interface method)
// with a dialect specific sql raw statement or query.
type GenFuncSpec struct {
	name      string
	statement string
}

// StatementFunc is a factory for GenFuncSpec.
func StatementFunc(name, statement string) *GenFuncSpec {
	return &GenFuncSpec{
		name:      name,
		statement: statement,
	}
}

// DataDefinitionStatement is a sum type for an ALTER, CREATE, DROP or RENAME statement.
type DataDefinitionStatement interface {
	alterOrCreateOrDropOrRename()
}

// A MigrationSpec combines a version (which is a parsed time.Time) and a bunch of DataDefinitionStatement.
type MigrationSpec struct {
	time       time.Time
	operations []DataDefinitionStatement
}

// Statement represents a SQL raw statement.
type Statement string

// alterOrCreateOrDropOrRename is the marker method for the according sum type.
func (s Statement) alterOrCreateOrDropOrRename() {
	panic("marker")
}

// Migration requires dateTime in ISO 8601 format (2006-01-02T15:04:05)
func Migrate(dateTime string, operations ...DataDefinitionStatement) *MigrationSpec {
	format := "2006-01-02T15:04:05" //TODO better place for a validator?
	timeValue, err := time.Parse(format, dateTime)
	if err != nil {
		panic("illegal format: " + dateTime + " expected " + format)
	}

	return &MigrationSpec{time: timeValue, operations: operations}
}

// Migrations is a factory for a slice of MigrationSpec.
func Migrations(migrations ...*MigrationSpec) []*MigrationSpec {
	return migrations
}
