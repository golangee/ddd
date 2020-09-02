package main

import (
	. "github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/golang"
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
					Factories(
						Struct("Opts", "...is an option type for the factory.",
							Field("Flag", Bool, "...is for something."),
						),
						Func("NewSearchService", "...is a factory to create a new SearchService.",
							In(
								Var("opts", "Opts", "... contains options to configure the instance."),
								Var("repo", "BookRepository", "... is the repo implementation to use."),
							),
							Out(
								Return("SearchService", "...is a package private instance."),
							),
						),
					),
				),

				// the use case layer
				UseCases(
					UseCase("BookSearch",
						"...provides all user stories around involved in searching books.",
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
						Story("As a search I want to see book details, because I need have to proof the relevance of the result.",
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
					),
					Factories(),
				),
				UseCases(
					UseCase("BookLoaning", "...provides all stories around loaning books.",
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
}
