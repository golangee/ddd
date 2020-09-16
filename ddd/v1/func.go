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
	pos     Pos
}

// Func is a factory for a FuncSpec.
func Func(name string, comment string, inSpec InParams, outSpec OutParams) *FuncSpec {
	return &FuncSpec{
		name:    name,
		comment: comment,
		in:      inSpec,
		out:     outSpec,
		pos:     capturePos("Func", 1),
	}
}

// funcOrStruct is a marker for FuncOrStruct.
func (m *FuncSpec) funcOrStruct() {
	panic("implement me")
}

// Name returns the name of the function.
func (m *FuncSpec) Name() string {
	return m.name
}

func (m *FuncSpec) Walk(f func(obj interface{}) error) error {
	if err := f(m); err != nil {
		return err
	}

	for _, obj := range m.in {
		if err := f(obj); err != nil {
			return err
		}
	}

	for _, obj := range m.out {
		if err := f(obj); err != nil {
			return err
		}
	}

	return nil
}

// In returns the according params.
func (m *FuncSpec) In() []*ParamSpec {
	return m.in
}

// InByName returns nil or the spec.
func (m *FuncSpec) InByName(name string) *ParamSpec {
	for _, spec := range m.in {
		if spec.name == name {
			return spec
		}
	}

	return nil
}

// Out returns the according params.
func (m *FuncSpec) Out() []*ParamSpec {
	return m.out
}

// Comment returns functions documentation.
func (m *FuncSpec) Comment() string {
	return m.comment
}

// Pos returns the debugging position.
func (m *FuncSpec) Pos() Pos {
	return m.pos
}
