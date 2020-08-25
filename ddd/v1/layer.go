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

// A Layer can be of very different kind. Its dependencies depends on the Stereotype and its concrete payload
// on the actual kind of the layer. Different IMPLEMENTATION types provides distinct types.
type Layer interface {
	// Name of the layer
	Name() string

	// Description of the layer
	Description() string

	// Stereotype of the layer
	Stereotype() Stereotype
}
