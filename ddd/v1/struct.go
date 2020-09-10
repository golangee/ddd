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
	comment        string
	name           string
	typeName       TypeName
	optional       bool
	pos            Pos
	defaultLiteral string
	jsonName       string
}

// Field is a factory for a FieldSpec.
func Field(name string, typeName TypeName, comment string) *FieldSpec {
	return &FieldSpec{
		comment:  comment,
		name:     name,
		typeName: typeName,
		pos:      capturePos("Field", 1),
	}
}

// SetOptional indicates if this field may be nil, if it declares a pointer type or if the default value makes
// any sense. This also influences also how other values are parsed or omitted.
func (f *FieldSpec) SetOptional(optional bool) *FieldSpec {
	f.optional = optional
	return f
}

// Optional returns the optional flag.
func (f *FieldSpec) Optional() bool {
	return f.optional
}

// SetDefault contains a literal to be used to initialize the value, if it is optional and absent or unparseable.
// It also sets automatically SetOptional to true
func (f *FieldSpec) SetDefault(literal string) *FieldSpec {
	f.defaultLiteral = literal
	f.SetOptional(true)
	return f
}

// Default returns the default literal
func (f *FieldSpec) Default() string {
	return f.defaultLiteral
}

// SetJsonName sets the json serialization name. Default is empty and conventionally the "package private" name
// (camel case with lower case starting)
func (f *FieldSpec) SetJsonName(name string) *FieldSpec {
	f.jsonName = name
	return f
}

// JsonName returns an alternative name to use or use the default if empty.
func (f *FieldSpec) JsonName() string {
	return f.jsonName
}

// Name of the field.
func (f *FieldSpec) Name() string {
	return f.name
}

// Comment of the field.
func (f *FieldSpec) Comment() string {
	return f.comment
}

// Pos for debugging.
func (f *FieldSpec) Pos() Pos {
	return f.pos
}

// TypeName of the field.
func (f *FieldSpec) TypeName() TypeName {
	return f.typeName
}
