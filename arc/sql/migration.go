package sql

import (
	"fmt"
	"github.com/golangee/architecture/arc"
	"io"
	"io/fs"
	"strings"
	"time"
)

// A Ctx describes the SQL specific environmental context.
type Ctx struct {
	// Dialect determines the SQL dialect to generate the support for.
	Dialect Dialect
	// Mod is the target module.
	Mod arc.String
	// Pkg is the target package within the module.
	Pkg arc.String
	// Migrations contains all user defined dialect specific raw migration statements.
	Migrations []*Migration
	// Repositories refers to all module-local interfaces which must be implemented (each as their own repository).
	// Note that not all methods have bindings and their implementations is either
	// omitted, abstract or otherwise stubbed out. This is a full qualified name
	// like my.company.MyType or my/company.MyType.
	Repositories []Repository
}

// A Migration represents a transactional group of sql migration statements. All of them should be applied or none.
// However due to SQL nature, many engines do not support that well with CREATE/DROP TABLE statements.
// We have no down/revert migrations, because in practice they don't make much sense and only work in few
// special cases (usually toy projects):
//  * for large database (million or even billions of rows) you surely don't want to wait for an alter-table.
//    Use the filesystems snapshot feature for a restore which can revert within a second.
//  * deleting or updating user-owned entries is never reversible.
//  * a failed migration cannot be safely undone using a down migration because many databases cannot alter tables
//    within a transaction.
type Migration struct {
	ID         time.Time
	Name       arc.String
	Statements []arc.String
}

// ParseMigrationName takes a name like 202009161147_the_initial_schema.sql and returns
// a time representing 20200916114700 and the text "the_initial_schema".
func ParseMigrationName(name string) (time.Time, string, error) {
	const dateFormat = "200601021504"
	if len(name) < len(dateFormat)+2 {
		return time.Time{}, "", fmt.Errorf("migration name to short")
	}

	t, err := time.Parse(dateFormat, name[0:len(dateFormat)])
	if err != nil {
		return time.Time{}, "", fmt.Errorf("time not parseable: %w", err)
	}

	ext := strings.LastIndex(name, ".")
	if ext < 0 {
		ext = len(name)
	}

	return t, name[len(dateFormat)+1 : ext], nil

}

// ParseStatements takes a byte sequence and splits it by ;\n to identify single statements. Every newline and
// whitespace padding per line is normalized to a single space. Statement semicolons are removed.
func ParseStatements(r io.Reader) ([]arc.String, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("unable to read statements: %w", err)
	}

	var fname string
	if fsFile, ok := r.(fs.File); ok {
		if stat, err := fsFile.Stat(); err == nil {
			fname = stat.Name()
		}
	}

	var res []arc.String
	tmp := &strings.Builder{}
	strBuf := string(buf)
	posLineBegin := 0
	lines := strings.Split(strBuf, "\n")
	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue
		}

		if tmp.Len() > 0 {
			tmp.WriteString(" ")
		}

		tmp.WriteString(trimmedLine)
		if strings.HasSuffix(trimmedLine, ";") {
			str := tmp.String()
			tmp.Reset()

			if len(strings.TrimSpace(str)) == 0 {
				continue // ignore empty statements
			}

			str = str[:len(str)-1] // remove ; suffix

			lit:=arc.NewString(str)
			lit.NodePos.Line = posLineBegin + 1
			lit.NodePos.Col = 1
			lit.NodePos.File = fname
			lit.NodeSrc = strBuf

			lit.NodeEnd.Line = i + 1
			lit.NodeEnd.Col = len(line)
			lit.NodeEnd.File = fname
			lit.NodeSrc = strBuf

			posLineBegin = i
			res = append(res, lit)
		}

	}

	// last statement may have missing ;
	if tmp.Len() > 0 {
		str := strings.TrimSpace(tmp.String())
		if len(str) == 0 {
			return res, nil
		}

		if strings.HasSuffix(str, ";") {
			str = str[:len(str)-1] // remove ; suffix
		}

		lit := core.NewStrLit(str)
		lit.NodePos.Line = posLineBegin + 1
		lit.NodePos.Col = 1
		lit.NodePos.File = fname
		lit.NodeSrc = strBuf

		lit.NodeEnd.Line = len(lines)
		lit.NodeEnd.Col = len(lines[len(lines)-1])
		lit.NodeEnd.File = fname
		lit.NodeSrc = strBuf

		res = append(res, lit)
	}

	return res, nil
}
