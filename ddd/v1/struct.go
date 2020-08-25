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

// A StructSpec represents a simple data class. It is usually treated as a value type.
type StructSpec struct {
	name    string
	comment string
	fields  []*FieldSpec
}

// funcOrStruct is a marker for FuncOrStruct
func (s *StructSpec) funcOrStruct() {
	panic("implement me")
}

// structOrInterface is a marker for StructOrInterface
func (s *StructSpec) structOrInterface() {
	panic("implement me")
}

// Name of the type.
func (s *StructSpec) Name() string {
	return s.name
}

// Struct is factory for a StructSpec and the contained FieldSpec.
func Struct(name, comment string, fields ...*FieldSpec) *StructSpec {
	return &StructSpec{
		name:    name,
		comment: comment,
		fields:  fields,
	}
}

// A FieldSpec represents a public field or attribute of a data class. The data type of a field is itself usually
// treated as a value type.
type FieldSpec struct {
	comment  []string
	name     string
	typeName TypeName
}

// Field is a factory for a FieldSpec.
func Field(name string, typeName TypeName, comment ...string) *FieldSpec {
	return &FieldSpec{
		comment:  comment,
		name:     name,
		typeName: typeName,
	}
}
