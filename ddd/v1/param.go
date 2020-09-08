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

type NameValidationKind int

const (
	NVGoPublicIdentifier NameValidationKind = iota
	NVGoPrivateIdentifier
	NVHttpHeaderParameter
	NVHttpQueryParameter
	NVHttpPathParameter
)

// A ParamSpec represents an incoming or outgoing function parameter.
type ParamSpec struct {
	name               string
	typeName           TypeName
	comment            []string
	pos                Pos
	nameValidationKind NameValidationKind
	optional           bool
}

// Var is a factory for a ParamSpec.
func Var(name string, typeName TypeName, comment string) *ParamSpec {
	return &ParamSpec{
		name:               name,
		nameValidationKind: NVGoPrivateIdentifier,
		typeName:           typeName,
		comment:            []string{comment},
		pos:                capturePos("Var", 1),
	}
}

// WithContext is a shortcut to Var("ctx","context.Context", "...is the context to control timeouts and cancellations.")
func WithContext() *ParamSpec {
	return &ParamSpec{
		name:               "ctx",
		nameValidationKind: NVGoPrivateIdentifier,
		typeName:           Ctx,
		comment:            []string{"...is the context to control timeouts and cancellations."},
		pos:                capturePos("Var", 1),
	}
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

// Optional returns true, if the parameter can be nil or the default value.
func (p *ParamSpec) Optional() bool {
	return p.optional
}

// SetOptional sets the optional flag to the param. If ParamSpec.TypeName indicates a
// pointer, a nil value is acceptable. Otherwise this is also very important for
// REST parameter where a missing or unparseable parameter would otherwise cause
// an early exit.
func (p *ParamSpec) SetOptional(optional bool) *ParamSpec {
	p.optional = optional
	return p
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

// NameValidationKind returns the validation hint
func (p *ParamSpec) NameValidationKind() NameValidationKind {
	return p.nameValidationKind
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
