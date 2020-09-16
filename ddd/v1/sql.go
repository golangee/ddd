package ddd

// A MySQLLayerSpec represents a stereotyped IMPLEMENTATION Layer.
type MySQLLayerSpec struct {
	name        string
	description string
	migrations  []*MigrationSpec
	genSpecs    []*RepoSpec
	pos         Pos
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

// Migrations returns the underlying slice of migrations.
func (u *MySQLLayerSpec) Migrations() []*MigrationSpec {
	return u.migrations
}

// Repositories returns the underlying slice of repositories.
func (u *MySQLLayerSpec) Repositories() []*RepoSpec {
	return u.genSpecs
}

// Pos returns the debug position.
func (u *MySQLLayerSpec) Pos() Pos {
	return u.pos
}

// Walks loops over self and migrations and repositories.
func (u *MySQLLayerSpec) Walk(f func(obj interface{}) error) error {
	if err := f(u); err != nil {
		return err
	}

	for _, obj := range u.migrations {
		if err := f(obj); err != nil {
			return err
		}
	}

	for _, repo := range u.genSpecs {
		if err := repo.Walk(f); err != nil {
			return err
		}
	}

	return nil
}

// MySQL is a factory for a MySQLLayerSpec. An implementation can only ever import Core API.
func MySQL(migrations []*MigrationSpec, genSpecs []*RepoSpec) *MySQLLayerSpec {
	return &MySQLLayerSpec{
		name:        "Mysql",
		description: "Package mysql contains specific repository implementations for the mysql dialect.",
		migrations:  migrations,
		genSpecs:    genSpecs,
		pos:         capturePos("MySQL", 1),
	}
}

// A RepoSpec refers to an interface and contains GenFuncSpec to map sql statements to method declarations and
// therefore also to input and output parameters.
type RepoSpec struct {
	interfaceName string
	funcs         []*GenFuncSpec
	pos           Pos
}

// Repository is a factory for a RepoSpec which refers to an interface and contains GenFuncSpec to map
// sql statements to method declarations and
// therefore also to input and output parameters.
func Repository(interfaceName string, funcs ...*GenFuncSpec) *RepoSpec {
	return &RepoSpec{
		interfaceName: interfaceName,
		funcs:         funcs,
		pos:           capturePos("Repository", 1),
	}
}

// Walks loops over self and implementations.
func (s *RepoSpec) Walk(f func(obj interface{}) error) error {
	if err := f(s); err != nil {
		return err
	}

	for _, obj := range s.funcs {
		if err := f(obj); err != nil {
			return err
		}
	}

	return nil
}

// InterfaceName returns the name of an repository or SPI interface defined in the core layer.
func (s *RepoSpec) InterfaceName() string {
	return s.interfaceName
}

// Pos returns the debug position.
func (s *RepoSpec) Pos() Pos {
	return s.pos
}

// Implementations returns the mapping of a SPI function, an SQL query and the according in/out parameters.
func (s *RepoSpec) Implementations() []*GenFuncSpec {
	return s.funcs
}

// ImplementationByName returns either the GenFuncSpec or nil if no such implementation is defined.
func (s *RepoSpec) ImplementationByName(name string) *GenFuncSpec {
	for _, spec := range s.funcs {
		if spec.name == name {
			return spec
		}
	}

	return nil
}

// Repositories is a factory for a slice of RepoSpec.
func Repositories(genSpecs ...*RepoSpec) []*RepoSpec {
	return genSpecs
}

// A GenFuncSpec combines a function name (more precisely an interface method)
// with a dialect specific sql raw statement or query.
type GenFuncSpec struct {
	name      string
	statement RawStatement
	in        []MapInParamSpec
	row       []MapRowParamSpec
	pos       Pos
}

// MapFunc maps a concrete interface method to a SQL statement. It also requires a mapping
// between the prepared statement parameters and the parameter input names. If a parameter is
// a struct, use a dot notation to reference a field. Same is true for the out parameter type.
// The slice type is automatically detected, however the order of the column selection to the
// actual return type of a single row must be also declared.
func MapFunc(name string, statement RawStatement, in []MapInParamSpec, row []MapRowParamSpec) *GenFuncSpec {
	return &GenFuncSpec{
		name:      name,
		statement: statement,
		in:        in,
		row:       row,
		pos:       capturePos("MapFunc", 1),
	}
}

// Params describes the mapping of function input parameters and prepared statements.
func (s *GenFuncSpec) Params() []MapInParamSpec {
	return s.in
}

// Row describes the mapping of the function output parameters and an sql result set row.
func (s *GenFuncSpec) Row() []MapRowParamSpec {
	return s.row
}

// Name returns the name of the function.
func (s *GenFuncSpec) Name() string {
	return s.name
}

// RawStatement returns the raw sql statement.
func (s *GenFuncSpec) RawStatement() RawStatement {
	return s.statement
}

// Pos returns the debug position.
func (s *GenFuncSpec) Pos() Pos {
	return s.pos
}

// MapInParamSpec introduces just a new string type to provide type safety.
type MapInParamSpec string

// Prepare aggregates a bunch of already defined parameter names and evaluates them in the given order
// for the prepared statement. The following notations are valid:
//  * Prepare("myParamName")
//  * Prepare("myParamName.Field")
//  * Prepare("myParamName.MyMethod()")
// Note that only the first parameter name is validated. Everything else in a calling chain is ignored, because
// at generation time, we cannot evaluate external dependencies anyway. However, your project may become uncompilable
// if no such field or method exists.
func Prepare(names ...MapInParamSpec) []MapInParamSpec {
	return names
}

// MapRowParamSpec introduces a new string type to provide type safety.
type MapRowParamSpec string

// MapRow aggregates a list of instructions to fill a single row result in definition order.
// It provides the following conventions:
//  * MapRow(".") allocates and returns a primitive or supported base type, as declared by the function signature.
//  * MapRow(".Field") allocates the result type and parses into the named field (which must be a pointer).
//  * MapRow("&.Field") allocates the result type and parses into the named field (which must not be a pointer).
func MapRow(references ...MapRowParamSpec) []MapRowParamSpec {
	return references
}

// A MigrationSpec combines a version (which is a parsed time.Time) and a bunch of DataDefinitionStatement.
type MigrationSpec struct {
	dateTime   string
	statements []RawStatement
	pos        Pos
}

// Migration requires dateTime in ISO 8601 format (2006-01-02T15:04:05)
func Migrate(dateTime string, statements ...RawStatement) *MigrationSpec {
	return &MigrationSpec{dateTime: dateTime, statements: statements, pos: capturePos("Migrate", 1)}
}

// DateTime returns a string in ISO 8601 format
func (m *MigrationSpec) DateTime() string {
	return m.dateTime
}

// Pos returns the debug position.
func (m *MigrationSpec) Pos() Pos {
	return m.pos
}

// RawStatements returns the set of ordered statements for this migration. A migration is not allowed to change
// secured by a checksum, if applied.
func (m *MigrationSpec) RawStatements() []RawStatement {
	return m.statements
}

// RawStatement represents a SQL raw statement.
type RawStatement string

// Migrations is a factory for a slice of MigrationSpec.
func Migrations(migrations ...*MigrationSpec) []*MigrationSpec {
	return migrations
}
