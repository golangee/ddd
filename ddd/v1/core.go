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
	name            string
	description     string
	api             []StructOrInterface
	implementations []*ServiceImplSpec
	pos             Pos
}

// Core has never any dependencies to any other layer.
func Core(api []StructOrInterface, implementations ...*ServiceImplSpec) *CoreLayerSpec {
	return &CoreLayerSpec{
		name: "Core",
		description: "Package core contains all domain specific models for the current bounded context.\n" +
			"It contains an exposed public API to be imported by other layers and an internal package \n" +
			"private implementation accessible by factory functions.",
		api:             api,
		implementations: implementations,
		pos:             capturePos("Core", -1),
	}
}

// API returns the struct or interfaces from the API definition.
func (c *CoreLayerSpec) API() []StructOrInterface {
	return c.api
}

// Pos returns the debugging position.
func (c *CoreLayerSpec) Pos() Pos {
	return c.pos
}

// Implementations returns the constructor or factory functions for the implementation of the API. Structs are only
// to describe parameters or factory options. The returned types must match the API interfaces and structs.
// The actual implementation must be performed by the developer and is not defined by the DSL.
func (c *CoreLayerSpec) Implementations() []*ServiceImplSpec {
	return c.implementations
}

// IsService returns true, if an implementation has been defined.
func (c *CoreLayerSpec) IsService(name string) bool {
	for _, impl := range c.implementations {
		if impl.Of() == name {
			return true
		}
	}
	return false
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

func (c *CoreLayerSpec) Walk(f func(obj interface{}) error) error {
	if err := f(c); err != nil {
		return err
	}

	for _, obj := range c.api {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	for _, obj := range c.implementations {
		if err := f(obj); err != nil {
			return err
		}

		if err := obj.options.Walk(f); err != nil {
			return err
		}
	}

	return nil
}
