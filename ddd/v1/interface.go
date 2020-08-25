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
}

// structOrInterface is a marker for StructOrInterface.
func (s *InterfaceSpec) structOrInterface() {
	panic("implement me")
}

// Name returns the name of interface.
func (s *InterfaceSpec) Name() string {
	return s.name
}

// Interface is a factory to create an InterfaceSpec.
func Interface(name, comment string, methods ...*FuncSpec) *InterfaceSpec {
	return &InterfaceSpec{
		name:    name,
		comment: comment,
		funcs:   methods,
	}
}
