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

package ddd_test

import (
	. "github.com/golangee/architecture/ddd/v1"
	"testing"
)

func TestDDD(t *testing.T) {
	app := Application("BookLibrary", "A story about a book library.",
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
								In(Var("bookId", UUID, "...a comment.")),
								Out(Return("Book", "...a comment.")),
							),
						),

						Interface("BookRepo", "...is a repo with books.",
							Func("FindBook",
								"... finds one book by its unique id.",
								In(Var("bookId", UUID, "...a comment.")),
								Out(Return("Book", "...a comment.")),
							),

							Func("FindAllByName",
								"... returns all books with a name like text.",
								In(Var("text", String, "...a comment.")),
								Out(Return(Slice("Book"), "...a comment."))),
							),
						),

						Interface("EBookRepo", "...contains the books in PDF.",
							Func("FindBook",
								"...finds a unique pdf book",
								In(Var("id", UUID, "...a comment.")),
								Out(Return(Reader, "...a comment.")),
							),
						),
					),
					Factories(
						Func("NewReadBookService",
							"...creates a domain instance for the book service",
							In(Var("repo", "BookRepo", "...a comment.")),
							Out(Return("ReadBookService", "...a comment.")),
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
								In(Var("bookId", UUID, "...a comment.")),
								Out(Return("Book", "...a comment.")),
							),
						),

						Interface("BorrowBookUseCase", "...represents the use case tbd.",
							Func("BorrowBook",
								"... represents the use case, where a Reader loans a physical book.",
								In(Var("bookId", UUID, "...a comment."), Var("readerId", UUID, "...a comment.")),
								Out(Return("Book", "...a comment.")),
							),

							Func("BorrowEbook",
								"... represents the use case, where a Reader loans an ebook.",
								In(Var("bookId", UUID, "...a comment."), Var("readerId", UUID, "...a comment.")),
								Out(Return("{{.Core}}.Book", "...a comment.")), //TODO uses the models and services from domain core, automatically in our scope???
							),
						),

						Interface("StatisticUseCase", "...represents the use case tbd.",
							Func("AllLoaners",
								"...represents the use case, to show all loaners to the inventory executor",
								In(),
								Out(Return(Slice("Reader", "...a comment."))),
							),
						),
					),

					Factories(
						Struct("FactoryOptions", "...is the options struct for the factory, which is implemented by the dev",
							Field("flag", Int64),
						),
						Func("NewStatisticUseCase",
							"...creates a new statistics use case",
							In(Var("randomValue", Int64, "...a comment."), Var("opts", "FactoryOptions", "...a comment.")),
							Out(Return("StatisticUseCase", "...a comment.")),
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
										Content(Json, Slice("Book")),
										Content("application/text", Slice("Book")),
										Content(Xml, Slice("Book")),
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
								Header(Var("clientId", String, "...a comment.")),
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
										Var("id", UUID, "...a comment."),
									),
									Header(
										Var("bearer", String, "...a comment."),
									),
									Query(
										Var("offset", Int64, "...a comment."),
										Var("limit", Int64, "...a comment."),
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
									Varchar("name", 255, NotNull(), Default("unnamed")),
									Binary("uuid", 16, Nullable()),
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