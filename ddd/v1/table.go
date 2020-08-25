package ddd

type ColumnOption struct {
}

type TableOption struct {
}

func NotNull() *ColumnOption {
	return nil
}

func Nullable() *ColumnOption {
	return nil
}

func Default(literal string) *ColumnOption {
	return nil
}

func AutoIncrement() *ColumnOption {
	return nil
}

func PrimaryKey(columns ...string) *TableOption {
	return nil
}

// ForeignKey automatically issues an ON DELETE CASCADE
func ForeignKey(column, refTable, refColumn string) *TableOption {
	return nil
}

type ColumnSpec struct {
}

func Varchar(name string, len int, opts ...*ColumnOption) *ColumnSpec {
	return nil
}

func Binary(name string, len int, opts ...*ColumnOption) *ColumnSpec {
	return nil
}

func Int(name string, len int, opts ...*ColumnOption) *ColumnSpec {
	return nil
}

func Column(name string) *ColumnSpec {
	return &ColumnSpec{}
}

func Columns(columns ...*ColumnSpec) []*ColumnSpec {
	return columns
}

type TableSpec struct {
}

func (t *TableSpec) alterOrCreateOrDropOrRename() {
	panic("implement me")
}

func CreateTable(name string, cols []*ColumnSpec, opts ...*TableOption) *TableSpec {
	return nil
}

type AlterTableSpec struct{

}

func AddColumns(cols...*ColumnSpec)*AlterTableSpec{
	return nil
}

func DropColumns(cols...string)*AlterTableSpec{
return nil
}

func AlterTable(name string,alter...*AlterTableSpec)*TableSpec{
	return nil
}