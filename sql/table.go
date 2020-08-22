package sql

type TableSpec struct {
}

func (t *TableSpec) createOrAlterSchema() {
	panic("implement me")
}

func CreateTable(name string)*TableSpec{
	return &TableSpec{}
}


