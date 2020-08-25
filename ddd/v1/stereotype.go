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

// A Stereotype in the sense of UML defines a usage context for a UseCaseLayerSpec.
type Stereotype string

const (
	// CORE defines the domain specific API within a bounded context. It has never a dependency to other layers.
	// If it needs something from the outside, it has to define a driven SPI (service provider interface).
	CORE Stereotype = "core"

	// A USECASE layer is only allowed to import dependencies from the CORE.
	USECASE Stereotype = "usecase"

	// PRESENTATION layer is only allowed to import from the USECASE layer and transitively from the CORE.
	// It provides its driven interface to the outside, e.g. a REST interface.
	PRESENTATION Stereotype = "presentation"

	// IMPLEMENTATION layer is only allowed to import SPI interfaces and types from the CORE layer. This is
	// typically used for MySQL or S3 implementations but may even issue RPC calls to other external services.
	IMPLEMENTATION Stereotype = "implementation"
)
