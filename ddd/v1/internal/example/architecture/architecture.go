package main

import (
	"github.com/golangee/architecture"
	. "github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/golang"
	"github.com/golangee/architecture/ddd/v1/markdown"
	"log"
)

func main() {
	spec := Application(
		"BookLibrary",
		"This is the central service of the book library of capital city, for searching and loaning books.",
		BoundedContexts(
			Context("security",
				"This context provides authentication (is a user the one he pretends to be?) and authorization "+
					"(has a user a specific role or is allowed to access an un(!)specific resource?).",
				Core(
					API(
						Struct("User", "...represents an authenticated user.",
							Field("ID", UUID, "...is the unique id of a user."),
							Field("Name", String, "...is the firstname of a user."),
							Field("Age", Int, "...is the age in years of a user."),
						),
						Interface("SecurityService", "...can authenticate and authorize a user from various sources.",
							Func("FromBearer", "...authenticates from a bearer token.",
								In(
									WithContext(),
									Var("bearer", String, "...is the bearer token from e.g. a http header variable."),
								),
								Out(
									Return("User", "...is the authenticated user."),
									Err(),
								),
							),

						),
					),
					Implementation("SecurityService",
						Requires(),
						Options(
							Field("KeycloakSecret", String, "...is an app secret for keycloak."),
							Field("KeycloakUrl", String, "...is the url of the keycloak service."),
						),
					),
				),

			),
			Context("search",
				"This context is about everything around searching for books.\n"+
					"A search can be issued from a users home or from within the library building by users or\n"+
					"employees. It allows access to any known book, which may not be even available physically,\n"+
					"like ebooks or new publications.",

				// the domain core providing the API
				Core(
					API(
						Struct("Book", "...is a book with meta data to index and find.",
							Field("ID", UUID, "...is the unique id of a book."),
							Field("Title", String, "...is the title for the book."),
							Field("Special", "github.com/google/uuid.UUID", "...is a test for importing a custom type."),
							Field("Tags", Slice(String), "...to describe a book."),
						),
						Interface("BookRepository",
							"...is a repository to handle books.",
							Func("ReadAll", "...returns all books.",
								In(
									WithContext(),
									Var("offset", Int64, "...is the offset to return the entries for paging."),
									Var("limit", Int64, "...is the maximum amount of entries to return."),
								),
								Out(
									Return(Slice("Book"), "...is the list of books."),
									Err(),
								),
							),

							Func("Count", "...enumerates all stored elements.",
								In(WithContext()),
								Out(
									Return(Int64, "...is the actual count."),
									Err(),
								),
							),

							Func("FindOne", "...finds exactly one entry.",
								In(
									WithContext(),
									Var("id", UUID, "...is the data transfer object to read into."),
								),
								Out(
									Return("Book", "...the found book."),
									Err(),
								),
							),

							Func("Insert", "...adds some stuff.",
								In(
									WithContext(),
									Var("dto", "Book", "...the book to save."),
								),
								Out(
									Err(),
								),
							),
						),
						Interface("SearchService", "...is the domain specific service API.",
							Func("Search", "...inspects each book for the key words.",
								In(
									WithContext(),
									Var("query", String, "...contains the query to search for."),
								),
								Out(
									Return(Slice("Book"), "...is the list of found books."),
									Err(),
								),
							),

						),
					),
					Implementation(
						"SearchService",
						Requires("BookRepository"),
						Options(
							Field("FulltextSearch", Bool, "...is a flag to enable fulltext search in items.").
								SetDefault("false").
								SetJsonName("full-text-SEARCH"),
							Field("Namespace", String, "...is a weired option.").
								SetDefault("\"some ugly stuff\""),
							Field("MyInt64", Int64, "... is an integer with 8 byte."),
							Field("MyFloat64", Float64, "... is a float with 64 bits.").
								SetDefault("5"),
							Field("MyDuration", Duration, "... is a duration."),
						),
					),

				),

				MySQL(
					Migrations(
						Migrate("2020-09-16T11:47:00", "Creates the initial schema.",
							"CREATE TABLE book (`id` BINARY(16), name VARCHAR(255))",
							"CREATE TABLE book3 (id BINARY(16))",
							"CREATE TABLE book3 (id JSON)",
							"SELECT blub from BLA",
							"INSERT INTO book3 VALUES(1,2)",
							"ALTER TABLE BLA ADD COLUMN (id TEXT)",
						),

						Migrate("2020-09-17T11:47:00", "Adding another table to support other books.",
							"CREATE TABLE book5 (id BINARY(16), name VARCHAR(255))",
						),
					),

					Repositories(
						Repository("BookRepository",
							MapFunc("ReadAll", "SELECT * FROM book LIMIT ? OFFSET ? ",
								Prepare("limit", "offset"),
								MapRow("&.ID", "&.Title"),
							),
							MapFunc("Count", "SELECT count(*) FROM book",
								Prepare(),
								MapRow("&."),
							),
							MapFunc("Insert", "INSERT INTO book(id) VALUES (?)",
								Prepare("dto.ID"),
								MapRow(),
							),

							MapFunc("FindOne", "SELECT * FROM book WHERE id=?",
								Prepare("id"),
								MapRow("&.ID", "&.Title"),
							),
						),
					),
				),

				// the use case layer
				UseCases(
					Epic("BookSearch",
						"...provides all user stories involved in searching books.",
						Stories(

							Story("As a searcher, I want to search for keywords, so that I must not know the title or author.",
								Func("FindByTags", "...searches for tags only.",
									In(
										WithContext(),
										Var("query", String, "...provides tokens to search for, separated by spaces or commas."),
									),
									Out(
										Return(Slice("Book"), "...is a list of books which match to the query."),
										Err(),
									),
								),
							),
							Story("As a searcher, I want to have an autocomplete so that I get support while typing my keywords.",
								Func("Autocomplete", "...proposes autocompletion values.",
									In(
										WithContext(),
										Var("text", String, "...is the text to autocomplete."),
									),
									Out(
										Return(Slice("AutoCompleteValue"), "...is a list of proposals."),
										Err(),
									),
								),
								Struct("AutoCompleteValue",
									"...represents an auto completed value.",
									Field("Value", String, "...is the value to complete."),
									Field("Score", Float32, "...the probability of importance."),
									Field("Synonyms", Slice(String), "...alternative search suggestions."),
								),
							),
							Story("As a searcher, I want to see book details, because I need to proof the relevance of the result.",
								Func("Details", "...returns the details of a book.",
									In(
										WithContext(),
										Var("id", UUID, "...is the Id of a book."),
										Var("user", "AuthUser", "...is the authenticated user."),
									),
									Out(
										Return("Book", "...is the according book."),
										Err(),
									),
								),
								Struct("AuthUser", "...represents an authenticated user.",
									Field("Age", Int, "...is the age of a user and determines if he can view the book or not."),
								),
							),

							Story("As a book admin, I want to change a title, because the book has a typo.",
								Func("ChangeBookTitle", "...changes the book title.",
									In(
										Var("titleModel", "BookTitleSpec", "... is to short."),
									),
									Out(
										Return("Book", "... the updated book."),
										Err(),
									),

								),
								Struct("BookTitleSpec", "...is for changing book titles.",
									Field("Title", String, "... is a title.").SetOptional(true),
								),

							),
						),
						Options(
							Field("EpicFeatureFlag", Bool, "...turns the magic feature on, if set to true."),
						),
					),

				),
				REST(
					Version("v1.0.1",
						"The initial base API provides all the things, which must be accessible by a mobile App.",
						Middlewares(
							Middleware("validateClientId",
								"...is here to check the validity of the client id.",
								Parse(
									Header(
										Var("clientId", String, "...is a client id"),
									),
								),
								Provide(),
							),
							Middleware("authenticate",
								"...takes a bearer or whatever and provides basic user properties.",
								Parse(
									Header(
										Var("authorization", String, "...is the jwt bearer token.").SetOptional(true),
										Var("cookie", String, "...is some cookie thingy.").SetOptional(true),
									),
									Path(),
									Query(),
								),
								Provide(
									Var("authUser", "AuthUser", "...is the user model from the use case layer."),
								),


							),
						),
						Resource("/books/:id/title",
							POST(
								After("validateClientId", "authenticate"),
								Invoke("BookSearch", "ChangeBookTitle",
									HeaderVar("x-uuid", "id"),
									BodyVar(Json, "titleModel"),
								),
								Response(
									HeaderVar("x-some-special", ".Title"),
									BodyVar(Json, ".Chapter"),
								),
							),
						),
					),
					Version("v2.0.3",
						"The v2 API introduces incompatible use cases for the usage, which is not only for mobile but also for web apps.",
						Middlewares(),
					),
				),


			),
			Context("loan",
				"This context is about everything around loaning or renting a book.\n"+
					"Only physical books can be loaned from within the library building by users.",
				Core(
					API(
						Struct("Book", "...is a book to loan or rent.",
							Field("ID", UUID, "...is the unique id of a book."),
							Field("ISBN", Int64, "...the international number."),
							Field("LoanedBy", Optional(UUID), "...is either nil or the user id."),
						),
						Struct("User", "... is a library customer.",
							Field("ID", UUID, "...is the unique id of the user."),
						),
						Interface("LoanService", "...provides stuff to loan all the things.",
							Func("LoanIt", "...loans a book.",
								In(),
								Out(),
							),
						),
					),
					Implementation(
						"LoanService",
						Requires(),
						Options(
							Field("Test", String, "...a test string."),
						),
					),
				),
				UseCases(
					Epic("BookLoaning", "...provides all stories around loaning books.",
						Stories(
							Story("As a book loaner, I have to scan the books barcode, so that I can take it with me.",
								Func("Rent", "...loans a book.",
									In(
										WithContext(),
										Var("bookId", UUID, "...is the id of the book."),
										Var("userId", UUID, "...is the id of the user, who loans the book."),
									),
									Out(Err()),
								),
							),
							Story("As a library staff, I have to check a customers library card, so that I can ensure that only actual customers can enter and loan books.",
								Func("CheckCustomerId", "...validates if the user is registered and active.",
									In(
										WithContext(),
										Var("userId", UUID, "...is the users id."),
									),
									Out(Err()),
								),
							),
						),
						Options(),

					),
				),
			),
		),
	)

	opts := golang.Options{
		ServerDir: "../server",
		ClientDir: "../client",
	}

	if err := golang.Generate(opts, spec); err != nil {
		log.Fatal(err)
	}

	const uml = false

	if uml {
		prj, err := architecture.Detect()
		if err != nil {
			log.Fatal(err)
		}

		if err := markdown.Generate(prj.File(opts.ServerDir), spec); err != nil {
			log.Fatal(err)
		}
	}

}
