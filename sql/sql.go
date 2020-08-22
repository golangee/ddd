package sql

import "time"



type ViewSpec struct {
}

type SchemaOperation interface {
	createOrAlterSchema()
}

type MigrationSpec struct {
	time       time.Time
	operations []SchemaOperation
}

// Migration requires dateTime in ISO 8601 format (2006-01-02T15:04:05)
func Migrate(dateTime string, operations ...SchemaOperation) *MigrationSpec {
	timeValue, err := time.Parse("2006-01-02T15:04:05", dateTime)
	if err != nil {
		panic("illegal format: " + dateTime)
	}

	return &MigrationSpec{time: timeValue, operations: operations}
}

func Migrations(migrations ...*MigrationSpec) []*MigrationSpec {
	return migrations
}
