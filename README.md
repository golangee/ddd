# architecture
This *architecture* module is especially useful to
bootstrap and maintain enterprise grade services.
The main goal is to reduce degeneration of the 
architecture and related documentation.

## planned features
*[ ] domain driven design, enforce correct dependency graph
*[ ] REST service generation
*[ ] OpenAPI generation
*[ ] Async http client generation support for easy WASM integration
*[ ] Ensure correct regeneration after changes
*[ ] MySQL Repository generation and migration support
*[ ] Generate UML and architecture Diagrams
*[ ] ...

## Alternatives

Actually, only [goa](https://goa.design/) is known.
However, *goa* neither enforces a *domain driven design*
nor does it provide a
meaningful, typesafe and autocompletion friendly
DSL. It does also not support planned features,
especially with regard to an automatic project 
documentation, deep integrations of frontends or 
other backend languages.

## domain driven design, variant 1

This is a very fluid DSL with factory methods all over the
place which integrates nicely with your favorite IDE providing
autocompletion.

DSL-Example
```go
package main

import (
	. "github.com/golangee/architecture/ddd/v1"
)

func main() {
	app := Application("BookLibrary",
		BoundedContexts(
			Context(
				"BookLending",
				"... is all about renting a book.",

				//This layer cannot have any dependencies to other layers.
				Core(
					API(
						Struct("Book", "...represents a virtual book",
							Field("ID", UUID, "unique id of a book"),
							Field("ISBN", Int64, "multiple books share the same ISBN"),
						),

						Struct("Reader", "...is a human reader.",
							Field("ID", UUID, "unique id of a library user"),
							Field("FirstName", String, "first name of user"),
							Field("LastName", String, "last name of user"),
						),

						Struct("Loan", "...is what a human does when renting a book",
							Field("ID", UUID, "unique id of a loan"),
							Field("BookId", UUID, "which book has been loaned"),
							Field("UserId", UUID, "which user has loaned it"), // überhaupt notwendig? oder
						),

						Interface("ReadBookService", "... is a service for books.",
							Func("ReadBook",
								"... represents the use case, where a Reader reads a book in the library.",
								In(Var("bookId", UUID)),
								Out(Return("Book")),
							),
						),

						Interface("BookRepo", "...is a repo with books.",
							Func("FindBook",
								"... finds one book by its unique id.",
								In(Var("bookId", UUID)),
								Out(Return("Book")),
							),

							Func("FindAllByName",
								"... returns all books with a name like text.",
								In(Var("text", String)),
								Out(Return(List("Book"))),
							),
						),

						Interface("EBookRepo", "...contains the books in PDF.",
							Func("FindBook",
								"...finds a unique pdf book",
								In(Var("id", UUID)),
								Out(Return(Reader)),
							),
						),
					),
					Factories(
						Func("NewReadBookService",
							"...creates a domain instance for the book service",
							In(Var("repo", "BookRepo")),
							Out(Return("ReadBookService")),
						),
					),
				),

				// this layer has only dependencies on the domain core
				UseCases(
					API(

						Struct("LoanReader", "...is a use case specific model.",
							Field("ID", UUID, "unique id of a library user"),
							Field("FirstName", String, "first name of user"),
							Field("LastName", String, "last name of user"),
							Field("BookId", UUID, "which book has been loaned"),
							Field("UserId", UUID, "which user has loaned it"), // überhaupt notwendig? oder
						),

						Interface("ReadBookUseCase", "...represents the use case around reading a book",
							Func("ReadBookInLibrary",
								"... represents the use case, where a Reader reads a book in the library.",
								In(Var("bookId", UUID)),
								Out(Return("Book")),
							),
						),

						Interface("BorrowBookUseCase", "...represents the use case tbd.",
							Func("BorrowBook",
								"... represents the use case, where a Reader loans a physical book.",
								In(Var("bookId", UUID), Var("readerId", UUID)),
								Out(Return("Book")),
							),

							Func("BorrowEbook",
								"... represents the use case, where a Reader loans an ebook.",
								In(Var("bookId", UUID), Var("readerId", UUID)),
								Out(Return("{{.Core}}.Book")), //TODO uses the models and services from domain core, automatically in our scope???
							),
						),

						Interface("StatisticUseCase", "...represents the use case tbd.",
							Func("AllLoaners",
								"...represents the use case, to show all loaners to the inventory executor",
								In(),
								Out(Return(List("Reader"))),
							),
						),
					),

					Factories(
						Struct("FactoryOptions", "...is the options struct for the factory, which is implemented by the dev",
							Field("flag", Int64),
						),
						Func("NewStatisticUseCase",
							"...creates a new statistics use case",
							In(Var("randomValue", Int64), Var("opts", "FactoryOptions")),
							Out(Return("StatisticUseCase")),
						),
					),
				),

				// the REST layer is a presentation layer and has only dependencies to the use case layer and therefore also to the domain layer.
				REST(
					"v1.1.1",
					Resources(
						Resource(
							"/books",
							"Resource to manage books.",
							Parameters(),
							GET("Returns all books.",
								Parameters(),
								Responses(
									Response(200, "book array response",
										Header(
											Var("x-RateLimit-Limit", Int64, "Request limit per hour."),
											Var("x-RateLimit-Remaining", Int64, "The number of requests left for the time window."),
										),
										Content(Json, List("Book")),
										Content("application/text", List("Book")),
										Content(Xml, List("Book")),
										Content("image/png", Reader),
										Content("image/jpg", Reader),
										Content(Any, Reader),
									),
									Response(404, "not found", Header()),
								),
							),
							DELETE("Removes all books",
								Parameters(),
								Responses(),
							),
						),
						Resource("/books/{id}",
							"Resource to manage a single book",
							Parameters(
								Header(Var("clientId", String)),
							),
							GET("Returns a single book.",
								Parameters(),
								Responses(),
							),
							DELETE("Removes a single book",
								Parameters(),
								Responses(),
							),
							PUT("Updates a book",
								Parameters(),
								RequestBody(),
								Responses(),
							),
							POST("Creates a new book",
								Parameters(
									Path(
										Var("id", UUID),
									),
									Header(
										Var("bearer", String),
									),
									Query(
										Var("offset", Int64),
										Var("limit", Int64),
									),
								),
								RequestBody(
									Content("image/png", Reader),
									Content("image/jpg", Reader),
									Content("application/octet-stream", Reader),
									Content(Json, "Book"),
								),
								Responses(
									Response(200, "ok, book array response",
										Header(
											Var("x-RateLimit-Limit", Int64, "Request limit per hour."),
											Var("x-RateLimit-Remaining", Int64, "The number of requests left for the time window."),
										),
										Content("image/png", Reader),
										Content("image/jpg", Reader),
										Content("application/octet-stream", Reader),
										Content(Json, "Book"),
									),
								),
							),
						),
					),
				),

				// a concrete implementation layer has only dependencies on the domainCore, especially the Repository interface and models. Does not depend on the use case or presentation.
				MySQL(
					Migrations(
						Migrate("2006-01-02T15:04:05",
							CreateTable("users",
								Columns(
									Int("id", 11, AutoIncrement()),
									Varchar("name", 255, NotNull()),
									Binary("uuid", 16),
								),
								PrimaryKey("id", "name"),
								ForeignKey("uuid", "objects", "id"),
							),
							AlterTable("users",
								AddColumns(
									Int("num", 3, NotNull()),
								),
								DropColumns("name"),
							),
						),
						Migrate("2006-01-02T15:04:06",
							Statement("CREATE TABLE ..."),
						),
					),

					// can only use domain core
					Generate(
						From("BookRepo",
							StatementFunc("FindBook", "SELECT * FROM books"),
							StatementFunc("FindAllByName",
								"SELECT * FROM users WHERE name like :text",
							),
						),
					),
				),

				S3(
					From("EBookRepo"), // TODO can only auto-implement trivial methods?
				),

				Filesystem(
					From("EBookRepo"), // TODO can only auto-implement trivial methods?
				),

			),
		),
	)

	_ = app
}
```