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

// A CoreLayerSpec represents a stereotyped CORE Layer.
type CoreLayerSpec struct {
	name        string
	description string
	api         []StructOrInterface
	factories   []FuncOrStruct
}

// Name of the Layer
func (c *CoreLayerSpec) Name() string {
	return c.name
}

// Description of the layer
func (c *CoreLayerSpec) Description() string {
	return c.description
}

// Stereotype of the layer
func (c *CoreLayerSpec) Stereotype() Stereotype {
	return CORE
}

// Core has never any dependencies to any other layer.
func Core(api []StructOrInterface, factories []FuncOrStruct) *CoreLayerSpec {
	return &CoreLayerSpec{
		name: "core",
		description: "Package core contains all domain specific models for the current bounded context.\n" +
			"It contains an exposed public API to be imported by other layers and an internal package \n" +
			"private implementation accessible by factory functions.",
		api:       api,
		factories: factories,
	}
}
