package stereotype

import "github.com/golangee/src/ast"

type valueKey string

const key valueKey = "stereotype"

type envKey string

const ekey envKey = "envKeyName"

const flagKey envKey = "flagKeyName"

const (
	ConfigureStruct         = "ConfigureStruct"
	MySQL                   = "MySQL"
	Database                = "Database"
)

func SetEnvName(node ast.Node, name string) {
	node.PutValue(ekey, name)
}

func GetEnvName(node ast.Node) string {
	v := node.Value(ekey)
	if s, ok := v.(string); ok {
		return s
	}

	return ""
}

func SetFlagName(node ast.Node, name string) {
	node.PutValue(flagKey, name)
}

func GetFlagName(node ast.Node) string {
	v := node.Value(flagKey)
	if s, ok := v.(string); ok {
		return s
	}

	return ""
}

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
