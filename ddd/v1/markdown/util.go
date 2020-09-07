package markdown

import "github.com/golangee/src"

// umlifyDeclName makes some readable name of it, without package qualifier etc.
func umlifyDeclName(dec *src.TypeDecl) string {
	tmp := string(dec.Qualifier())
	for _, decl := range dec.Params() {
		tmp += umlifyDeclName(decl)
	}
	return tmp
}
