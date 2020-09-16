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

// BoundedContexts is a factory to create a slice of BoundedContextSpec from variable amount of arguments.
func BoundedContexts(contexts ...*BoundedContextSpec) []*BoundedContextSpec {
	return contexts
}

// A BoundedContextSpec contains the information about the layers within a bounded context. It contains different
// layers which may dependencies with each other, depending on their stereotype.
type BoundedContextSpec struct {
	name        string
	description string
	layers      []Layer
	pos         Pos
}

// Context is a factory for a BoundedContextSpec.
func Context(name string, description string, layers ...Layer) *BoundedContextSpec {
	return &BoundedContextSpec{
		name:        name,
		description: description,
		layers:      layers,
		pos:         capturePos("Context", 1),
	}
}

// DomainServices picks up the CoreLayerSpec and returns those InterfaceSpec which have an implementation entry point.
func (s *BoundedContextSpec) DomainServices() []*InterfaceSpec {
	var res []*InterfaceSpec
	for _, layer := range s.layers {
		if api, ok := layer.(*CoreLayerSpec); ok {
			for _, structOrInterface := range api.API() {
				if api.IsService(structOrInterface.Name()) {
					res = append(res, structOrInterface.(*InterfaceSpec))
				}
			}
		}
	}
	return res
}

// SPIServices returns those interfaces which are not a Core Service but a service provider interface aka
// driven port.
func (s *BoundedContextSpec) SPIServices() []*InterfaceSpec {
	var res []*InterfaceSpec
	for _, layer := range s.layers {
		if api, ok := layer.(*CoreLayerSpec); ok {
			for _, structOrInterface := range api.API() {
				if !api.IsService(structOrInterface.Name()) {
					if iface, ok := structOrInterface.(*InterfaceSpec); ok {
						res = append(res, iface)
					}

				}
			}
		}
	}
	return res
}

func (s *BoundedContextSpec) Layers() []Layer {
	return s.layers
}

// Name of the context.
func (s *BoundedContextSpec) Name() string {
	return s.name
}

// Description of the context.
func (s *BoundedContextSpec) Description() string {
	return s.description
}

// Pos returns debug position.
func (s *BoundedContextSpec) Pos() Pos {
	return s.pos
}

func (s *BoundedContextSpec) Walk(f func(obj interface{}) error) error {
	if err := f(s); err != nil {
		return err
	}

	for _, layer := range s.layers {
		if err := layer.Walk(f); err != nil {
			return err
		}
	}

	return nil
}
