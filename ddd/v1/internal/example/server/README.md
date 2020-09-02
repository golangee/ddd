# BookLibrary

This is the central service of the book library of capital city, for searching and loaning books.

## Index

* [BookLibrary](#booklibrary)
  * [Index](#index)
  * [Architecture](#architecture)
    * [The context *search*](#the-context-search)
      * [The domains core layer](#the-domains-core-layer)
        * [Type *Book*](#type-book)
        * [Type *BookRepository*](#type-bookrepository)
        * [Type *SearchService*](#type-searchservice)
        * [Factory *NewSearchService*](#factory-newsearchservice)
      * [UML](#uml)
      * [The use case or application layer](#the-use-case-or-application-layer)
        * [BookSearch](#booksearch)
    * [The context *loan*](#the-context-loan)
      * [The domains core layer](#the-domains-core-layer)
        * [Type *Book*](#type-book)
        * [Type *User*](#type-user)
      * [UML](#uml)
      * [The use case or application layer](#the-use-case-or-application-layer)
        * [BookLoaning](#bookloaning)


## Architecture

The server is organized after the domain driven design principles.
It is separated into the following 2 bounded contexts.

### The context *search*

This context is about everything around searching for books.
A search can be issued from a users home or from within the library building by users or
employees. It allows access to any known book, which may not be even available physically,
like ebooks or new publications.

#### The domains core layer

The core layer or API layer of the domain consists of 1 data types,
2 service or SPI interfaces and 1 actual service implementations.

##### Type *Book*

The data class *Book* is a book with meta data to index and find.

##### Type *BookRepository*

The SPI interface *BookRepository* is a repository to handle books.

##### Type *SearchService*

The API interface *SearchService* is the domain specific service API.

##### Factory *NewSearchService*

The API factory method *NewSearchService* is a factory to create a new SearchService.

#### UML

![search core API](uml-search-core-api.gen.svg?raw=true)

#### The use case or application layer

The following use case is defined.

##### BookSearch

The use case *BookSearch* provides all user stories around involved in searching books.

### The context *loan*

This context is about everything around loaning or renting a book.
Only physical books can be loaned from within the library building by users.

#### The domains core layer

The core layer or API layer of the domain consists of 2 data types,
0 service or SPI interfaces and 0 actual service implementations.

##### Type *Book*

The data class *Book* is a book to loan or rent.

##### Type *User*

The data class *User* is a library customer.

#### UML

![loan core API](uml-loan-core-api.gen.svg?raw=true)

#### The use case or application layer

The following use case is defined.

##### BookLoaning

The use case *BookLoaning* provides all stories around loaning books.

