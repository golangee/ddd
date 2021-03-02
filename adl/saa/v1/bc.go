package spec

// A BoundedContext defines a coherent model and clear boundaries. It defines a context in which
// the model is clear and fully described (see Eric Evans in "Domain-Driven Design" page 335ff).
// Dependencies, which are upstream connections to other boundaries,
// must be as weak as possible and isolated through an anticorruption layer, which is at least an interface or
// the definition of incoming events mapped to service endpoints.
//
// A BoundedContext must always be treated as if it has a different technical or physical location, like
// a different process on another machine. This must be kept true, even if the BoundedContexts manifests itself
// just as a simple package within a monolith. Here I introduce the first very important differences or clearifications
// in definition:
//
//   * A BoundedContext may be a module (like a Go module or a Gradle Module or a Maven Artifact) but it
//     may also just be a simple package in a monolith.
//   * A BoundedContext defines public Services which expressed as an Interface within other bounded contexts
//     always create an anti-corruption layer.
//   * If you want to express a Microservice you have to model it as a BoundedContext. This 1:1 relation can be
//     considered a misconception from Evans concept, but it is my simplification.
//   * The defined ubiquitous language has a 1:1 mapping to its BoundedContext.
//   * BoundedContext can (and should) create dependencies through interfaces (anti-corruption layer) and in general
//     there should be nothing like a global user interface, application or infrastructure layer. Divide and conquer
//     by user stories and put aggregating use-case views into their own bounded context(s). See also Vernon, page 532
//     in "Domain-Driven Design".
//   * Shared kernels are also considered BoundedContexts. They share actually the same properties but are
//     usually provided as a kind of module or library dependency to other contexts. Just like programming languages,
//     frameworks or other external resources, they are upstream in a bounded context hierarchy.
//   * There is no such thing like a single "Core Domain" (see Vernon, page 50). Larger and diversified companies
//     may have multiple core domains and interviews may provide contradictory information.
//   * The awareness of specialisation in subdomains is essentially irrelevant. Bounded contexts may cut them
//     arbitrarily anyway and there may be even situations where you are responsible to validate and trigger
//     a change-management due to higher prioritized (non-)functional requirements. Evans and Vernon hold
//     the misconception, that domain experts are always right and that it is possible to reconcile their
//     requirements. Prepare yourself that you have a (partially) better understanding of your customers business
//     and that you find yourself in a consultant role just because you asked unspoken questions.
//   * This is probably the most important rule: Always favor an anemic model over the object oriented domain
//     driven design model. Even though Vernon and Fowler call this an antipattern, my experience have proved them
//     wrong: it is the lesser of all evils in practice. When it comes to software quality, Hibernate et al.
//     represent an actually failed concept. Projects with non-functional requirements equal to toy projects
//     can be solved in an excellent way. But as soon quality constraints like maintainability, memory-, cpu-
//     consumption or big O problems start to matter, you are usually screwed. If you think this is inflammatory
//     just take a look at their bug tracker (https://hibernate.atlassian.net/issues/?filter=-5) and ask yourself
//     how many of those problems can be solved trivially, if you would not have picked this kind of abstraction
//     layer. Then ask your customer, if he is willing to pay you twiddling with it. For sure, there are examples
//     where a non-ORM approach looks infeasible, however, personally I've not encountered them in practice yet.
type BoundedContext struct {
	Name               Identifier
	Comment            String    // Comment should describe why this fragment exists.
	UbiquitousLanguage *Glossary // The UbiquitousLanguage defines a vocabulary using a simple Glossary.
	// Ports describe interfaces as anti-corruption layers. They are solely (private) dependencies for the Services.
	// They are part of the hexagonal-like architecture.
	Ports       []*Interface
	Services    []*Stub   // Services are stubs finally implemented by the developer. This is the anemic root.
	EntityTypes []*Entity // EntityTypes contains all declared entity types. This is the anemic model.
	ValueTypes  []*Value  // ValueTypes contains all declared value types.

	Adapters []Adapter
}

type Adapter interface {
	Name() Identifier
	Implements() []Identifier // returns all identifiers which will be implemented by this adapter.
	Stereotype()
}
