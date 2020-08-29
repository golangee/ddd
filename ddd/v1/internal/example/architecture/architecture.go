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
									Var("ctx", Ctx, "...provides the timeout handling."),
								),
								Out(
									Return(Slice("Book"), "...is the list of books."),
									Return(Error, "...returns an implementation specific failure."),
								),
							),
						),
						Interface("SearchService", "...is the domain specific service API.",
							Func("Search", "...inspects each book for the key words.",
								In(
									Var("ctx", Ctx, "...provides the timeout handling."),
									Var("query", String, "...contains the query to search for."),
								),
								Out(
									Return(Slice("Book"), "...is the list of found books."),
									Return(Error, "...returns an implementation specific failure."),
								),
							),

						),
					),
					Factories(
						Struct("Opts", "...is an option type for the factory.",
							Field("Flag", Bool, "...for something."),
						),
						Func("NewSearchService", "...is a factory to create a new SearchService.",
							In(
								Var("opts", "Opts", "... contains options to configure the instance."),
								Var("repo", "BookRepository", "... is the repo implementation to use."),
							),
							Out(
								Return("BookRepository", "...is a package private instance."),
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
