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

// A TypeName is a string suited to identify a custom type (Struct or Interface) or for build-in types like Int64
// or String.
type TypeName string

const (
	Int64   TypeName = "int64"
	Int32   TypeName = "int32"
	Bool    TypeName = "bool"
	Float32 TypeName = "float32"
	Float64 TypeName = "float64"
	Int16   TypeName = "int16"
	Int8    TypeName = "int8"
	Byte    TypeName = "byte"
	String  TypeName = "string"
	UUID    TypeName = "uuid"
	Error   TypeName = "error"
	Reader  TypeName = "io.Reader"
	Writer  TypeName = "io.Writer"
	Ctx     TypeName = "context.Context"
)

// StructOrInterface is a marker interface to declare a sum type of accepting either Struct or Interface types.
type StructOrInterface interface {
	Name() string
	// Walk through all nested objects, including self as root.
	Walk(f func(obj interface{}) error) error

	structOrInterface()
}

// FuncOrStruct is a marker interface to declare a sum type of accepting either Struct or Func types.
type FuncOrStruct interface {
	Name() string
	// Walk through all nested objects, including self as root.
	Walk(f func(obj interface{}) error) error
	funcOrStruct()
}

// Slice is generic convenience wrapper to create a slice type of t.
func Slice(t TypeName) TypeName {
	return "[]" + t
}

// Optional is just a pointer to the type.
func Optional(t TypeName) TypeName {
	return "*" + t
}

// API is a factory for a slice of StructOrInterface types.
func API(specs ...StructOrInterface) []StructOrInterface {
	return specs
}

// Factories is a factory for a slice of FuncOrStruct types.
func Factories(specs ...FuncOrStruct) []FuncOrStruct {
	return specs
}
