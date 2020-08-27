package ddd

import "strconv"

// A ColumnOption represents properties like Nullable, NotNull, Default or AutoIncrement.
type ColumnOption struct {
	nullable       bool
	defaultLiteral string
	defaultValue   bool
	autoIncrement  bool
}

// A ColumnOptionFunc modifies a ColumnOption
type ColumnOptionFunc func(opt *ColumnOption)

// NotNull sets nullable to false, which is the default.
func NotNull() ColumnOptionFunc {
	return func(opt *ColumnOption) {
		opt.nullable = false
	}
}

// Nullable sets nullable to true.
func Nullable() ColumnOptionFunc {
	return func(opt *ColumnOption) {
		opt.nullable = true
	}
}

// Default will apply a default literal when no such value is available.
func Default(literal string) ColumnOptionFunc {
	return func(opt *ColumnOption) {
		opt.defaultLiteral = literal
		opt.defaultValue = true
	}
}

// AutoIncrement automatically increases an integer primary key.
func AutoIncrement() ColumnOptionFunc {
	return func(opt *ColumnOption) {
		opt.autoIncrement = true
	}
}

// ColumnSpec is the model to represent a SQL column definition.
type ColumnSpec struct {
	name     string
	len      int
	dataType string
	ColumnOption
}

// newColumnSpec is a factory for a ColumnSpec and already applies the ColumnOptionFunc.
func newColumnSpec(name, dataType string, len int, opts ...ColumnOptionFunc) *ColumnSpec {
	col := &ColumnSpec{
		name:     name,
		dataType: dataType,
		len:      len,
	}

	for _, opt := range opts {
		opt(&col.ColumnOption)
	}

	return col
}

// Varchar creates the according sql column.
func Varchar(name string, len int, opts ...ColumnOptionFunc) *ColumnSpec {
	return newColumnSpec(name, "VARCHAR", len, opts...)
}

// Binary creates the according sql column.
func Binary(name string, len int, opts ...ColumnOptionFunc) *ColumnSpec {
	return newColumnSpec(name, "BINARY("+strconv.Itoa(len)+")", len, opts...)
}

// Int creates the according sql column.
func Int(name string, len int, opts ...ColumnOptionFunc) *ColumnSpec {
	return newColumnSpec(name, "INT("+strconv.Itoa(len)+")", len, opts...)
}

// Columns is a factory for a slice of ColumnSpec.
func Columns(columns ...*ColumnSpec) []*ColumnSpec {
	return columns
}

// A TableOption represents table related properties like PrimaryKey or a ForeignKey constraint.
type TableOption struct {
	primaryKey  []string
	foreignKeys []ForeignKeyConstraint
}

// PrimaryKey sets the current primary key to the named local columns.
func PrimaryKey(columns ...string) TableSpecFunc {
	return func(opt *TableOption) {
		opt.primaryKey = columns
	}
}

// ForeignKeyConstraint describes a relation between
type ForeignKeyConstraint struct {
	localColumn       string
	refTable          string
	refColumn         string
	onDeleteRefOption string
	onUpdateRefOption string
}

// ForeignKey adds another constraint and automatically issues an ON DELETE and ON UPDATE CASCADE
func ForeignKey(column, refTable, refColumn string) TableSpecFunc {
	return func(opt *TableOption) {
		opt.foreignKeys = append(opt.foreignKeys, ForeignKeyConstraint{
			localColumn:       column,
			refTable:          refTable,
			refColumn:         refColumn,
			onDeleteRefOption: "CASCADE",
			onUpdateRefOption: "CASCADE",
		})
	}
}

// TableSpec describes the structure of a SQL table.
type TableSpec struct {
	name string
	cols []*ColumnSpec
	TableOption
}

func (t *TableSpec) alterOrCreateOrDropOrRename() {
	panic("marker")
}

// A ColumnOptionFunc modifies a ColumnOption
type TableSpecFunc func(opt *TableOption)

// CreateTable is factory for a TableSpec and directly applies the options.
func CreateTable(name string, cols []*ColumnSpec, opts ...TableSpecFunc) *TableSpec {
	table := &TableSpec{
		name:        name,
		cols:        cols,
		TableOption: TableOption{},
	}

	for _, opt := range opts {
		opt(&table.TableOption)
	}

	return table
}

// AlterTableSpec contains all modifications which should be applied to the named table.
type AlterTableSpec struct {
	name           string
	addedColumns   []*ColumnSpec
	droppedColumns []string
}

// alterOrCreateOrDropOrRename is the marker interface for the DataDefinitionStatement.
func (a *AlterTableSpec) alterOrCreateOrDropOrRename() {
	panic("marker")
}

// AlterTableSpecFunc modifies an AlterTableSpec.
type AlterTableSpecFunc func(spec *AlterTableSpec)

// AlterTable is a factory for a new AlterTableSpec and applies all AlterTableSpecFunc on it.
func AlterTable(name string, alterOpts ...AlterTableSpecFunc) *AlterTableSpec {
	table := &AlterTableSpec{
		name: name,
	}
	for _, opt := range alterOpts {
		opt(table)
	}

	return table
}

// AddColumns appends the columns when applied.
func AddColumns(cols ...*ColumnSpec) AlterTableSpecFunc {
	return func(spec *AlterTableSpec) {
		spec.addedColumns = append(spec.addedColumns, cols...)
	}
}

// DropColumns removes the columns when applied.
func DropColumns(cols ...string) AlterTableSpecFunc {
	return func(spec *AlterTableSpec) {
		spec.droppedColumns = append(spec.droppedColumns, cols...)
	}
}
