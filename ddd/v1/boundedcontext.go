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

// A BoundedContextSpec contains the information about the layers within a bounded context. It contains different
// layers which may dependencies with each other, depending on their stereotype.
type BoundedContextSpec struct {
	name        string
	description string
	layers      []Layer
}

// BoundedContexts is a factory to create a slice of BoundedContextSpec from variable amount of arguments.
func BoundedContexts(contexts ...*BoundedContextSpec) []*BoundedContextSpec {
	return contexts
}

// Context is a factory for a BoundedContextSpec.
func Context(name string, description string, layers ...Layer) *BoundedContextSpec {
	return &BoundedContextSpec{
		name:        name,
		description: description,
		layers:      layers,
	}
}
