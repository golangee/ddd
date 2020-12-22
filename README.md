# architecture
This *architecture* module is especially useful to
bootstrap and maintain enterprise grade services.
The main goal is to reduce degeneration of the 
architecture and related documentation.

## planned features
- [ ] domain driven design, enforce correct dependency graph
- [ ] REST service generation
- [ ] OpenAPI generation
- [ ] Async http client generation support for easy WASM integration
- [ ] Ensure correct regeneration after changes
- [ ] MySQL Repository generation and migration support
- [ ] Generate UML and architecture Diagrams
- [ ] ...

## Alternatives

* [goa](https://goa.design/) seems to be very similar and quite popular.
However, *goa* neither enforces a *domain driven design*
nor does it provide a
meaningful, typesafe and autocompletion friendly
DSL. It does also not support planned features,
especially with regard to an automatic project 
documentation, deep integrations of frontends or 
other backend languages.
* [gozero](https://github.com/tal-tech/go-zero/blob/master/readme-en.md) shares also some basic ideas and seems
to enforce a better structure than goa, but otherwise it lacks the same features.
They also use a custom API DSL, which have its own pros and cons. Also, most of its documentation seems to be in chinese,
which makes it harder to understand. In a random sample, the source code itself seems not be documented at all. 
* [kok](https://github.com/RussellLuo/kok) is a generator toolkit for [gokit](https://github.com/go-kit/kit)

## domain driven design, variant 1

This is a very fluid DSL with factory methods all over the
place which integrates nicely with your favorite IDE providing
autocompletion.

See the example: 
* [handwritten architecture](ddd/v1/internal/example/architecture)
* [generated Server](ddd/v1/internal/example/server)
