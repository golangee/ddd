// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ddd

import "strings"

// A ParamSpec represents an incoming or outgoing function parameter.
type ParamSpec struct {
	name     string
	typeName TypeName
	comment  []string
	pos      Pos
}

// Var is a factory for a ParamSpec.
func Var(name string, typeName TypeName, comment string) *ParamSpec {
	return &ParamSpec{
		name:     name,
		typeName: typeName,
		comment:  []string{comment},
		pos:      capturePos("Var", 1),
	}
}

// WithContext is a shortcut to Var("ctx","context.Context", "...is the context to control timeouts and cancellations.")
func WithContext() *ParamSpec {
	return Var("ctx", Ctx, "...is the context to control timeouts and cancellations.")
}

// Return is a factory for an unnamed and undocumented ParamSpec.
func Return(typeName TypeName, comment string) *ParamSpec {
	p := Var("", typeName, comment)
	p.pos = capturePos("Return", 1)
	return p
}

// Err is a shortcut for Return(Error, "...is returned to indicate a violation of pre- or invariant and represents an implementation specific failure.")
func Err() *ParamSpec {
	return Return(Error, "...indicates a violation of pre- or invariants and represents an implementation specific failure.")
}

// Name returns the name of parameter.
func (p *ParamSpec) Name() string {
	return p.name
}

// Comment returns the parameters documentation.
func (p *ParamSpec) Comment() string {
	return strings.Join(p.comment, "\n")
}

// TypeName returns the declared type of the parameter.
func (p *ParamSpec) TypeName() TypeName {
	return p.typeName
}

// Pos returns debugging position.
func (p *ParamSpec) Pos() Pos {
	return p.pos
}

// InParams is just here for type safety and auto completion.
type InParams []*ParamSpec

// OutParams is just here for type safety and auto completion.
type OutParams []*ParamSpec

// In is a factory for a slice of ParamSpec converted to InParams.
func In(params ...*ParamSpec) InParams {
	return params
}

// Out is a factory for a slice of ParamSpec converted to OutParams.
func Out(params ...*ParamSpec) OutParams {
	return params
}
