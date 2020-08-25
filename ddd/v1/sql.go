package ddd

import "time"

type StatementSpec interface{
	statement()
}

type DataDefinitionStatement interface{
	alterOrCreateOrDropOrRename()
}

type ViewSpec struct {
}


type MigrationSpec struct {
	time       time.Time
	operations []DataDefinitionStatement
}

type Statement string

func (s Statement) alterOrCreateOrDropOrRename() {
	panic("implement me")
}

// Migration requires dateTime in ISO 8601 format (2006-01-02T15:04:05)
func Migrate(dateTime string, operations ...DataDefinitionStatement) *MigrationSpec {
	timeValue, err := time.Parse("2006-01-02T15:04:05", dateTime)
	if err != nil {
		panic("illegal format: " + dateTime)
	}

	return &MigrationSpec{time: timeValue, operations: operations}
}


func Migrations(migrations ...*MigrationSpec) []*MigrationSpec {
	return migrations
}
