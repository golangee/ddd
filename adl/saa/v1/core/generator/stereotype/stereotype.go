package stereotype

import "github.com/golangee/src/ast"

type valueKey string

const key valueKey = "stereotype"

const (
	ConfigureStruct = "ConfigureStruct"
	MySQL           = "MySQL"
	Database        = "Database"
)

func Put(node ast.Node, stereotypes ...string) {
	list := Get(node)
	for _, stereotype := range stereotypes {
		has := false
		for _, s := range list {
			if stereotype == s {
				has = true
				break
			}
		}

		if !has {
			list = append(list, stereotype)
		}
	}

	node.PutValue(key, list)
}

func Get(node ast.Node) []string {
	if v, ok := node.Value(key).([]string); ok {
		return append([]string{}, v...)
	}

	return nil
}

func Has(node ast.Node, stereotype string) bool {
	for _, s := range Get(node) {
		if s == stereotype {
			return true
		}
	}

	return false
}