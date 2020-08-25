package ddd

// A Stereotype in the sense of UML defines a usage context for a UseCaseLayerSpec.
type Stereotype string

const(
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
