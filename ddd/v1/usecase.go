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

// A UseCaseLayerSpec represents a stereotyped USECASE Layer.
type UseCaseLayerSpec struct {
	name        string
	description string
	api         []StructOrInterface
	factories   []FuncOrStruct
}

// Name of the layer.
func (u *UseCaseLayerSpec) Name() string {
	return u.name
}

// Description of the layer.
func (u *UseCaseLayerSpec) Description() string {
	return u.description
}

// Stereotype of the layer.
func (u *UseCaseLayerSpec) Stereotype() Stereotype {
	return USECASE
}

// UseCases is a factory for a UseCaseLayerSpec. A use case can only ever import Core API.
func UseCases(api []StructOrInterface, factories []FuncOrStruct) *UseCaseLayerSpec {
	return &UseCaseLayerSpec{
		name: "usecases",
		description: "Package usecases contains all domain specific use cases for the current bounded context.\n" +
			"It contains an exposed public API to be imported by PRESENTATION layers. \n" +
			"It provides a private implementation of the use cases accessible by factory functions.",
		api:       api,
		factories: factories,
	}
}
