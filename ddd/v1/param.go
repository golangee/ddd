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

// A ParamSpec represents an incoming or outgoing function parameter.
type ParamSpec struct {
	name     string
	typeName TypeName
	comment  []string
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

// Var is a factory for a ParamSpec.
func Var(name string, typeName TypeName, comment ...string) *ParamSpec {
	return &ParamSpec{
		name:     name,
		typeName: typeName,
		comment:  comment,
	}
}

// Return is a factory for an unnamed and undocumented ParamSpec.
func Return(typeName TypeName) *ParamSpec {
	return Var("", typeName)
}
