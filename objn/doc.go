// Package objn provides an OBJect Notation interface. The notation
// is entirely declarative and only consists of the following elements:
//  - everything is a Node and a Pos (position) to locate it.
//  - a Map contains distinct key/value combinations
//  - a Seq contains an ordered array resp. sequence of Nodes
//  - a Lit represents a scalar resp. literal value
//  - a Pkg may contain other Pkg or Doc nodes.
//  - a Doc contains the actual root Map, Seq or Lit.
// There may be multiple implementations, like yaml, json or even source code.
package objn
