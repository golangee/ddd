// Package yast (yet another abstract syntax tree) provides a model for an
// abstract syntax tree of a simple object notation, like json or yaml
// configurations. Each element is represented by a Node which may have attached positions to emit helpful
// error messages. Also each Node may have its own Stereotype which allows to carry more information besides
// the inherent structure itself - without having to misuse fields itself which would complicate the modelling
// of literals and sequences even more.
//
// Intentionally, the ast model consists of concrete types and not just by interfaces to ensure that
// modifications can be applied trivially. This also allows to merge multiple parsed sources (e.g. from json,
// yaml, xml or even source code) into a common representation.
package yast
