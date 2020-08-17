package ddd_test

import (
	. "github.com/golangee/ddd"
	"testing"
)

func TestDDD(t *testing.T) {
	err := ApplicationDomain("BookLibrary",
		BoundedContexts(
			Context(
				"BookLending",
				"... is all about renting a book.",
				Requires(),
				UseCases(
					Func("ReadBookInLibrary",
						"... represents the use case, where a Reader reads a book in the library.",
						In(Var("bookId", UUID)),
						Out(Return("Book")),
					),
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
					Func("AllLoaners",
						"...represents the use case, to show all loaners to the inventory executor",
						In(),
						Out(Return(List("Reader"))),
					),
				),
				Types(
					Type("Book",
						Field("ID", "unique id of a book", UUID),
						Field("ISBN", "multiple books share the same ISBN", Int64),
					),
					Type("Reader",
						Field("ID", "unique id of a library user", UUID),
						Field("FirstName", "first name of user", String),
						Field("LastName", "last name of user", String),
					),
					Type("Loan",
						Field("ID", "unique id of a loan", UUID),
						Field("BookId", "which book has been loaned", UUID),
						Field("UserId", "which user has loaned it", UUID), // überhaupt notwendig? oder
					),
				),
				Presentations(
					REST("v1.0.0",
						Resources(
							GET("/books/{id}/reader", "Returns the current reader of the book",
								In(Var("id", String)),
								Out(Return("Reader")),
							),
							GET("/books/{id}", "Returns a single book",
								In(Int32("id", "id of the book")),
								Out(),
							),
							DELETE("/books/{id}/borrows/{readerId}", "returns a book by removing a reader from the borrowers list."),
						),
						Types(
							Type("Reader",
								Field("ID", "...the id of the reader", UUID),
								Field("Name", "...first- and lastname of the user", UUID),
							),
						),
					),
				),
				Service(), //the application service is another layer of separation, which contains usually vertical constraints like user-management which is not domain specific
				Persistence(
					Repositories(
						Interface("ReaderRepository",
							Func("FindAll",
								"...returns all readers.", In(), Out(),
							),
							Func("FindOne", "...returns the first matching reader or an error",
								In(), Out(),
							),
						),
						Interface("LoanRepository",
							Func("ReadPDF", "...opens a book to read",
								In(Var("file", String, "...is the filename")), Out(Return(Reader), Return(Error)),
							),
						),
						Interface("BookRepository",
							Func("FindAll", "...returns all entries.",
								In(
									Int32("offset2"),
									Int32("offset3", "...some comment"),
									Var("offset", Int64, "...the index to return from"),
									Var("limit", Int64, "...returns at most"),
								),
								Out(
									Var("names", List(String), "all names"),
									Var("err", Error, "returned if something went wrong"),
								),
							),
							Func("FindOne",
								"...returns the first matching entry or an error",
								In(Var("id", UUID, "the id of entity to find")),
								Out(Return(List(String))),
							),
						),
					),
					Types(
						Type("Device",
							Field("Id", "...is unique per device.", UUID),
							Field("Name", "...is an arbitrary non unique name.", String),
							Field("power", "...is the power consumption in Ah", Int64),
						),
					),
					Implementations(
						MySQL("BookStore",
							Schema(
								Migrate(20200815153330,
									`CREATE TABLE books (id binary(16), name VARCHAR(255)) PRIMARY KEY (id);
										-- and many more stuff

										`,
								),
								Migrate(20200815153330, "ALTER TABLE books ADD COLUMN isbn BIGINT"),
							),
							Implement("FindOne", Statement("SELECT id, name FROM books WHERE id=:id")),
							Implement("Insert", DefaultCreate("books")),
							Implement("Remove", DefaultDelete("books")),
						),
						Filesystem("BookStore"),
					),
				),
			),

			Context("BookSearch",
				"... is about searching books.",
				Requires(),
				UseCases(),
				Types(
					Type("Book",
						Field("ISBNS", "hardcover, softcover and ebooks have different ISBNs but same content", List(UUID)),
						Field("Author", "name of author", String),
						Field("Title", "...of the book", String),
						Field("Tags", "... for the search", List(String)),
					),
					Type("SearchResult",
						Field("Relevance", "...priority of the result", Int64),
					),
				),

				Presentations(),
				Service(),
				Persistence(
					Repositories(
						Interface("BookRepository",
							Func("FindAll",
								"...returns all entries.",
								In(),
								Out(Return("Book")), //cross-reference between layers is possible but creates redundant types
							),
							Func("FindOne",
								"...returns the first matching entry or an error",
								In(),
								Out(Return("Book")), //cross-reference between layers is possible but creates redundant types
							),
						),
						Interface("SearchRepository",
							Func("Search",
								"...returns all entries.",
								In(Var("query", String)),
								Out(Return(List("SearchResult"))), //cross-reference between layers is possible but creates redundant types
							),
						),
					),
					Types(
						Type("Portfolio",
							Field("Id", "unique id", UUID),
							Field("Name", "human readable string", String),
						),
					),
					Implementations(),
				),
			),
			Context("BookSystem",
				"...is for statistic guys, which create reports for the higher business men",
				Requires( // this is the firewall / open host coupling mechanism which is full-fillable by BookLending
					Interface("LoanerLister",
						Func("AllLoaners",
							"...gets all loaners",
							In(),
							Out(Return(List("Reader"))),
						),
					),
					Interface("BookSearcher",
						Func("Search",
							"...returns all entries.",
							In(Var("query", String)),
							Out(Return(List("SearchResult"))), //cross-reference between layers is possible but creates redundant types
						),
					),
				),
				UseCases(
					Func("CalculateInteractions",
						"... represents the use case, where a statistic guy creates a report",
						In(),
						Out(),
					),
				),
				Types(),
				Presentations(),
				Service(),
				Persistence(Repositories(), Types(), Implementations()),
			),
		),
	).Generate()

	if err != nil {
		t.Fatal(err)
	}
}
