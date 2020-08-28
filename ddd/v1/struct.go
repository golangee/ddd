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

// A StructSpec represents a simple data class. It is usually treated as a value type.
type StructSpec struct {
	name    string
	comment string
	fields  []*FieldSpec
	pos     Pos
}

// Struct is factory for a StructSpec and the contained FieldSpec.
func Struct(name, comment string, fields ...*FieldSpec) *StructSpec {
	return &StructSpec{
		name:    name,
		comment: comment,
		fields:  fields,
		pos:     capturePos("Struct", 1),
	}
}

// Name of the struct.
func (s *StructSpec) Name() string {
	return s.name
}

// Comment of the struct.
func (s *StructSpec) Comment() string {
	return s.comment
}

// Pos for debugging.
func (s *StructSpec) Pos() Pos {
	return s.pos
}

// Fields returns the declared fields.
func (s *StructSpec) Fields() []*FieldSpec {
	return s.fields
}

// funcOrStruct is a marker for FuncOrStruct
func (s *StructSpec) funcOrStruct() {
	panic("implement me")
}

// structOrInterface is a marker for StructOrInterface
func (s *StructSpec) structOrInterface() {
	panic("implement me")
}

func (s *StructSpec) Walk(f func(obj interface{}) error) error {
	if err := f(s); err != nil {
		return err
	}

	for _, obj := range s.fields {
		if err := f(obj); err != nil {
			return err
		}
	}

	return nil
}

// A FieldSpec represents a public field or attribute of a data class. The data type of a field is itself usually
// treated as a value type.
type FieldSpec struct {
	comment  []string
	name     string
	typeName TypeName
	pos      Pos
}

// Field is a factory for a FieldSpec.
func Field(name string, typeName TypeName, comment ...string) *FieldSpec {
	return &FieldSpec{
		comment:  comment,
		name:     name,
		typeName: typeName,
		pos:      capturePos("Field", 1),
	}
}

// Name of the field.
func (f *FieldSpec) Name() string {
	return f.name
}

// Comment of the field.
func (f *FieldSpec) Comment() string {
	return strings.Join(f.comment, "\n")
}

// Pos for debugging.
func (f *FieldSpec) Pos() Pos {
	return f.pos
}

// TypeName of the field.
func (f *FieldSpec) TypeName() TypeName {
	return f.typeName
}
