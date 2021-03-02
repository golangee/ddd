package sql

// An Adapter combines a dialect with multiple migrations and repository definitions.
type Adapter struct {
	Dialect      Dialect
	Migrations   Migrations
	Repositories []Repository
}
