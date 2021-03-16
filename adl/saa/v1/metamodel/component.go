package metamodel

/*
// A Component is quite similar to a BoundedContext. However, the structure enforces a subset of a BC and
// also makes some opinionated decisions which some may call "antipattern". Also this is more a pre-compile time
// component, because it has no real runtime model later (no component model like OSGI etc).
type Component struct {
	// Domain is not allowed to use anything else.
	Domain Domain

	// UseCases is only allowed to use Domain
	UseCases UseCases

	// Presentation is only allowed to use UseCases
	Presentation Presentation

	// Generators contains all emitters to build this server. This may be something like a Spring Boot or Go server
	// generator or Android or iOS.
	Generators []Generator
}

type Presentation struct {
	Http Http
	RCP  RCP
}

// RCP is a Rich Client Platform like Android or iOS.
type RCP struct {
	MVPVMs []MVPVM
}

// MVPVM is model-view-presenter-viewmodel.
// See also https://herbertograca.com/2017/08/17/mvc-and-its-variants/#model-view-presenter-view_model.
// The Model part is provides by Component.UseCases.
type MVPVM struct {
	View       types.Stub  // Views are only allowed to call Presenters
	Presenter  types.Stub  // Presenters will only return ViewModels. It uses UseCases.
	ViewModels []types.DTO // ViewModels contains View specific DTOs
}

type Http struct {
	Rest Rest
}
type Rest struct {
}

// UseCases declares the domains (application) layer. This layer offers everything for concrete use cases and
// uses the Domain to do its job. It usually transforms information into use case specific DTOs.
type UseCases struct {
	// Services provides a kind of interface which manifests itself as a frame for a developer to get implemented.
	// It is more a kind of an abstract class.
	Services []types.Stub

	// Ports declares dependencies of the use cases to the outside world. These are satisfied by the dependency
	// injection.
	Ports []types.Interface

	// DTOs declares just some data transfer objects, suited to serve each use case properly.
	DTOs []types.DTO
}

// Domain declares the domain (core) layer. The one, which should be implemented first, by a developer. It has
// no dependencies to other written things. It provides the base operations within the domain. Sometimes it is
// hard to distinguish from UseCases.
type Domain struct {
	// Services provides a kind of interface which manifests itself as a frame for a developer to get implemented.
	// It is more a kind of an abstract class.
	Services []types.Stub

	// Ports declares dependencies of the use cases to the outside world. These are satisfied by the dependency
	// injection.
	Ports []types.Interface

	// Entities contains all Entity declarations, which should be more or less anemic (besides basic validations
	// and modifications). An Entity is always represented by an ID but is still a copy and is not a singleton, like
	// a User. It is always a mistake to include business rules and aggregate roots here (in the sense of DDD).
	Entities []types.Entity

	// Values contains all Value declarations, which are usually anemic (besides basic validations). A value
	// has no ID, like an E-Mail or a Phone number.
	Values []types.Value

	// Adapter implementations within this context.
	Adapter Adapter
}

// Adapter aggregates context-bound Port implementations. A Port definition may
// get implemented from different Adapters which may not belong to the same category (e.g. there may be
// a mysql and postgres adapter but also an IPC adapter to another microservice which is perhaps event driven).
type Adapter struct {
	Persistence Persistence
}

// Persistence contains adapter implementations associated to the persistence layer.
type Persistence struct {
	// SQL provides Adapter implementations for a subset of the ports declared above. There may be ports, which
	// will never be implemented within the Domain.
	SQL []sql.Adapter
}

*/
