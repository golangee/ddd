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
      * [The use case or application layer](#the-use-case-or-application-layer)
        * [BookSearch](#booksearch)
      * [REST API *v1.0.1*](#rest-api-v101)
        * [/books](#books)
          * [*GET* /books](#get-books)
          * [*DELETE* /books](#delete-books)
        * [/books/:id](#booksid)
          * [*GET* /books/:id](#get-booksid)
          * [*DELETE* /books/:id](#delete-booksid)
          * [*PUT* /books/:id](#put-booksid)
          * [*POST* /books/:id](#post-booksid)
    * [The context *loan*](#the-context-loan)
      * [The domains core layer](#the-domains-core-layer)
        * [Type *Book*](#type-book)
        * [Type *User*](#type-user)
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
2 service or SPI interfaces and 0 actual service implementations.

##### Type *Book*

The data class *Book* is a book with meta data to index and find.

##### Type *BookRepository*

The SPI interface *BookRepository* is a repository to handle books.

##### Type *SearchService*

The SPI interface *SearchService* is the domain specific service API.

![search core API](uml-search-core-api.gen.svg?raw=true)

#### The use case or application layer

The following use case is defined.

##### BookSearch

The use case *BookSearch* provides all user stories involved in searching books.
It contains 3 user stories.

|As a/an|I want to...|So that...|
|---|---|---|
|searcher|search for keywords|I must not know the title or author|
|searcher|have an autocomplete|I get support while typing my keywords|
|searcher|see book details|I need to proof the relevance of the result|


![use case-BookSearch](uml-use-case-booksearch.gen.svg?raw=true)



![iface-BookSearch](uml-iface-booksearch.gen.svg?raw=true)

#### REST API *v1.0.1*

Package rest contains the REST specific implementation for the current bounded context.
It depends only from the use cases and transitively on the core API.

##### /books

Resource to manage books.

###### *GET* /books

Returns all books.

```bash
curl -v -X GET https://capital-city-library.com/api/v1/books

```
###### *DELETE* /books

Removes all books.

```bash
curl -v -X DELETE https://capital-city-library.com/api/v1/books

```
##### /books/:id

Resource to manage a single book.

###### *GET* /books/:id

Returns a single book.

```bash
curl -v -X GET https://capital-city-library.com/api/v1/books/:id

```
###### *DELETE* /books/:id

Removes a single book.

```bash
curl -v -X DELETE https://capital-city-library.com/api/v1/books/:id

```
###### *PUT* /books/:id

Updates a book.

```bash
curl -v -X PUT https://capital-city-library.com/api/v1/books/:id

```
###### *POST* /books/:id

Creates a new book.

```bash
curl -v -X POST https://capital-city-library.com/api/v1/books/:id

```
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

![loan core API](uml-loan-core-api.gen.svg?raw=true)

#### The use case or application layer

The following use case is defined.

##### BookLoaning

The use case *BookLoaning* provides all stories around loaning books.
It contains 2 user stories.

|As a/an|I want to...|So that...|
|---|---|---|
|book loaner|scan the books barcode|I can take it with me|
|library staff|check a customers library card|I can ensure that only actual customers can enter and loan books|


![use case-BookLoaning](uml-use-case-bookloaning.gen.svg?raw=true)



![iface-BookLoaning](uml-iface-bookloaning.gen.svg?raw=true)

