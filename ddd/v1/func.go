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

// A FuncSpec describes a function, either as a real package level function (usually a factory method or constructor)
// or as a struct scoped member function (a method).
type FuncSpec struct {
	name    string
	comment string
	in      []*ParamSpec
	out     []*ParamSpec
}

// funcOrStruct is a marker for FuncOrStruct.
func (m *FuncSpec) funcOrStruct() {
	panic("implement me")
}

// Name returns the name of the function.
func (m *FuncSpec) Name() string {
	return m.name
}

// Func is a factory for a FuncSpec.
func Func(name string, comment string, inSpec InParams, outSpec OutParams) *FuncSpec {
	return &FuncSpec{
		name:    name,
		comment: comment,
		in:      inSpec,
		out:     outSpec,
	}
}
