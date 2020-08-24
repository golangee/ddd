package ddd_test

import (
	. "github.com/golangee/ddd"
	. "github.com/golangee/ddd/sql"
	"testing"
)

func TestDDD(t *testing.T) {
	err := ApplicationDomain("BookLibrary",
		BoundedContexts(
			Context(
				"BookLending",
				"... is all about renting a book.",
				DomainCore(
					API(
						Struct("Book",
							Field("ID", UUID, "unique id of a book"),
							Field("ISBN", Int64, "multiple books share the same ISBN"),
						),

						Struct("Reader",
							Field("ID", UUID, "unique id of a library user"),
							Field("FirstName", String, "first name of user"),
							Field("LastName", String, "last name of user"),
						),

						Struct("Loan",
							Field("ID", UUID, "unique id of a loan"),
							Field("BookId", UUID, "which book has been loaned"),
							Field("UserId", UUID, "which user has loaned it"), // überhaupt notwendig? oder
						),

						Interface("ReadBookService",
							Func("ReadBook",
								"... represents the use case, where a Reader reads a book in the library.",
								In(Var("bookId", UUID)),
								Out(Return("Book")),
							),
						),
					),
					Factories(),
				),

				UseCases(
					API(
						Struct("Book",
							Field("ID", UUID, "unique id of a book"),
							Field("ISBN", Int64, "multiple books share the same ISBN"),
						),

						Struct("Reader",
							Field("ID", UUID, "unique id of a library user"),
							Field("FirstName", String, "first name of user"),
							Field("LastName", String, "last name of user"),
						),

						Struct("Loan",
							Field("ID", UUID, "unique id of a loan"),
							Field("BookId", UUID, "which book has been loaned"),
							Field("UserId", UUID, "which user has loaned it"), // überhaupt notwendig? oder
						),

						Interface("ReadBookUseCase",
							Func("ReadBookInLibrary",
								"... represents the use case, where a Reader reads a book in the library.",
								In(Var("bookId", UUID)),
								Out(Return("Book")),
							),
						),

						Interface("BorrowBookUseCase",
							Func("BorrowBook",
								"... represents the use case, where a Reader loans a physical book.",
								In(Var("bookId", UUID), Var("readerId", UUID)),
								Out(Return("Book")),
							),

							Func("BorrowEbook",
								"... represents the use case, where a Reader loans an ebook.",
								In(Var("bookId", UUID), Var("readerId", UUID)),
								Out(Return("Book")),
							),
						),

						Interface("StatisticUseCase",
							Func("AllLoaners",
								"...represents the use case, to show all loaners to the inventory executor",
								In(),
								Out(Return(List("Reader"))),
							),
						),
					),

					Factories(
						Struct("FactoryOptions",
							Field("flag", Int64),
						),
						Func("NewStatisticUseCase",
							"...creates a new statistics use case",
							In(Var("randomValue", Int64), Var("opts", "FactoryOptions")),
							Out(Return("StatisticUseCase")),
						),
					),
				),
				REST(
					"v1.1.1",
					Resources(
						Resource(
							"/books",
							"Resource to manage books.",
							GET("Returns all books.", In(),
								Responses(
									Response(200, "book array response",
										Headers(
											Header("x-RateLimit-Limit", Int64, "Request limit per hour."),
											Header("x-RateLimit-Remaining", Int64, "The number of requests left for the time window."),
										),
										ContentTypes(
											JSON(List("Book")),
											Text(List("Book")),
											XML(List("Book")),
											JPEG("io.Reader"),
											BinaryStream(List("byte")),
										),
									),
									Response(404, "book not found",
										Headers(),
										ContentTypes(ForMimeType(MimeTypeJson, "error")),
									),
								),
							),
							DELETE("Removes all books"),
						),
						Resource("/books/{id}",
							"Resource to manage a single book",
							GET("Returns a single book.", In(), Responses()),
							DELETE("Removes a single book"),
							PUT("Updates a book"),
							POST("Creates a new book"),
						),
					),
					Types(
						Struct("Book",
							Field("ID", UUID, "unique id of a book"),
							Field("ISBN", Int64, "multiple books share the same ISBN"),
						),

						CopyStruct("UseCases", "Book", Fields()),
					),
				),
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
					Types(
						CopyStruct("UseCases", "Book",
							Fields(
								Field("PKId", Int64, "autoincrement mysql id"),
								RemoveField("ID"),
							),
						),
					),
					Generate(
						Repository("Users",
							"... manages a lot of user stuff.",
							StatementFunc("FindAll",
								"...finds all users.",
								"SELECT * FROM users",
								In(),
								Out(Return(List("Book"))),
							),
							StatementFunc("FindAllById",
								"...finds all users with a name.",
								"SELECT * FROM users WHERE name = :name",
								In(Var("name", String, "...is the name to find.")),
								Out(Return(List("Book"))),
							),
						),
						CRUDRepository("CRUDUsers",
							"...is an auto-implemented crud repository",
							"users",
							"Book",
							ReadAll|CountAll,
						),
					),
				),
			),
		),
	).Generate()

	if err != nil {
		t.Fatal(err)
	}
}
