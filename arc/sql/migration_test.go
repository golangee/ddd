package sql

import (
	"github.com/golangee/architecture/adl/saa/v1/core"
	"io"
	"reflect"
	"testing"
	"time"
)

func TestParseMigrationName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		args    args
		want    time.Time
		want1   string
		wantErr bool
	}{
		{
			args:    args{name: "202009161147_the_initial_schema.sql"},
			want:    time.Date(2020, 9, 16, 11, 47, 0, 0, time.UTC),
			want1:   "the_initial_schema",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.name, func(t *testing.T) {
			got, got1, err := ParseMigrationName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMigrationName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMigrationName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseMigrationName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParseStatements(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []core.StrLit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStatements(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStatements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStatements() got = %v, want %v", got, tt.want)
			}
		})
	}
}