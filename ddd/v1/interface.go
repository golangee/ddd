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

// InterfaceSpec represents an interface contract containing multiple FuncSpec.
type InterfaceSpec struct {
	name    string
	comment string
	funcs   []*FuncSpec
	pos     Pos
}

// Interface is a factory to create an InterfaceSpec.
func Interface(name, comment string, methods ...*FuncSpec) *InterfaceSpec {
	return &InterfaceSpec{
		name:    name,
		comment: comment,
		funcs:   methods,
		pos:     capturePos("Interface", 1),
	}
}

// Walk marches over all sub objects.
func (s *InterfaceSpec) Walk(f func(obj interface{}) error) error {
	if err := f(s); err != nil {
		return err
	}

	for _, obj := range s.funcs {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	return nil
}

// structOrInterface is a marker for StructOrInterface.
func (s *InterfaceSpec) structOrInterface() {
	panic("implement me")
}

// Name returns the name of interface.
func (s *InterfaceSpec) Name() string {
	return s.name
}

// Comment returns the documentation of the interface.
func (s *InterfaceSpec) Comment() string {
	return s.comment
}

// Funcs returns the method set.
func (s *InterfaceSpec) Funcs() []*FuncSpec {
	return s.funcs
}

// FuncByName returns nil, if no such FuncSpec has been defined.
func (s *InterfaceSpec) FuncByName(name string) *FuncSpec {
	for _, spec := range s.funcs {
		if spec.name == name {
			return spec
		}
	}

	return nil
}

// Pos returns the debugging position.
func (s *InterfaceSpec) Pos() Pos {
	return s.pos
}
