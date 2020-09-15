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
								),
								Out(
									Return(Slice("Book"), "...is the list of books."),
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
									),
									Out(
										Return("Book", "...is the according book."),
										Err(),
									),
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
				REST("v1.0.1",
					Servers(
						Server("https://capital-city-library.com", "The live server."),
						Server("https://dev.capital-city-library.com", "The development server."),
					),
					Resources(
						Resource("/books",
							"Resource to manage books.",
							Parameters(
								Header(
									Var("session", String, "...is the global valid session."),
								),
							),
							GET("Returns all books.",
								Parameters(),
								Responses(
									Response(200, "book array response",
										Header(
											Var("x-RateLimit-Limit", Int64, "... is the request limit per hour."),
											Var("x-RateLimit-Remaining", Int64, "...is the number of requests left for the time window."),
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
							DELETE("Removes all books.",
								Parameters(),
								Responses(),
							),
						),
						Resource("/books/:id",
							"Resource to manage a single book.",
							Parameters(
								Header(Var("clientId", String, "...is a comment.")),
							),
							GET("Returns a single book.",
								Parameters(),
								Responses(),
							),
							DELETE("Removes a single book.",
								Parameters(),
								Responses(),
							),
							PUT("Updates a book.",
								Parameters(),
								RequestBody(),
								Responses(),
							),

							POST("Creates a new book.",
								Parameters(
									Path(
										Var("id", UUID, "...is a comment."),
									),
									Header(
										Var("bearer", String, "...is a token bearer."),
										Var("x-special-something", String, "...is something special.").SetOptional(true),
									),
									Query(
										Var("offset", Int64, "...is a comment."),
										Var("limit", Int64, "...is a comment.").SetOptional(true),
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
											Var("x-RateLimit-Limit", Int64, "...is th request limit per hour."),
											Var("x-RateLimit-Remaining", Int64, "...is the number of requests left for the time window."),
										),
										Content("image/png", Reader),
										Content("image/jpg", Reader),
										Content("application/octet-stream", Reader),
										Content(Json, "Book"),
									),
								),
							),

						),

					)),
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
