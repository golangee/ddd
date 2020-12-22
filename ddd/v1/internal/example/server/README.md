# BookLibrary

This is the central service of the book library of capital city, for searching and loaning books.

## Index

* [BookLibrary](#booklibrary)
  * [Index](#index)
  * [Architecture](#architecture)
    * [The context *security*](#the-context-security)
      * [The domains core layer](#the-domains-core-layer)
        * [Type *User*](#type-user)
        * [Type *AuthenticationService*](#type-authenticationservice)
    * [The context *search*](#the-context-search)
      * [The domains core layer](#the-domains-core-layer)
        * [Type *Book*](#type-book)
        * [Type *BookRepository*](#type-bookrepository)
        * [Type *SearchService*](#type-searchservice)
        * [Factory *SearchService*](#factory-searchservice)
      * [Persistence layer: *Mysql*](#persistence-layer-mysql)
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
        * [Type *LoanService*](#type-loanservice)
        * [Factory *LoanService*](#factory-loanservice)
      * [The use case or application layer](#the-use-case-or-application-layer)
        * [BookLoaning](#bookloaning)
  * [usage](#usage)
    * [FulltextSearch](#fulltextsearch)
    * [Namespace](#namespace)
    * [MyInt64](#myint64)
    * [MyFloat64](#myfloat64)
    * [MyDuration](#myduration)
    * [Test](#test)


## Architecture

The server is organized after the domain driven design principles.
It is separated into the following 3 bounded contexts.

### The context *security*

This context provides authentication (is a user the one he pretends to be?) and authorization (has a user a specific role or is allowed to access an un(!)specific resource?).

#### The domains core layer

The core layer or API layer of the domain consists of 1 data types,
1 service or SPI interfaces and 0 actual service implementations.

##### Type *User*

The data class *User* represents an authenticated user.

##### Type *AuthenticationService*

The SPI interface *AuthenticationService* can authenticate a user from various sources.

![security core API](uml-security-core-api.gen.svg?raw=true)

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

##### Factory *SearchService*

The API factory method *SearchServiceFactory* creates an instance.
It requires the interfaces *BookRepository* as dependencies.
The instance must be configured using the following options:
 * FulltextSearch (...is a flag to enable fulltext search in items.)
 * Namespace (...is a weired option.)
 * MyInt64 (... is an integer with 8 byte.)
 * MyFloat64 (... is a float with 64 bits.)
 * MyDuration (... is a duration.)

![search core API](uml-search-core-api.gen.svg?raw=true)

#### Persistence layer: *Mysql*

The *mysql* persistence layer for the domain *search* consists of
2 migrations and has been evolved as shown in the following table:

|Introduced at|Description|Statements|
|---|---|---|
|2020-09-16T11:47:00|Creates the initial schema.|1x alter<br>3x create<br>1x insert<br>1x other<br>|
|2020-09-17T11:47:00|Adding another table to support other books.|1x create<br>|
After applying all migrations (as of 2020-09-17T11:47:00), the final sql schema looks like follows:

![Mysql-er](uml-mysql-er.gen.svg?raw=true)

#### The use case or application layer

The following use case is defined.

##### BookSearch

The use case *BookSearch* provides all user stories involved in searching books.
It contains 4 user stories.

|As a/an|I want to...|So that...|
|---|---|---|
|searcher|search for keywords|I must not know the title or author|
|searcher|have an autocomplete|I get support while typing my keywords|
|searcher|see book details|I need to proof the relevance of the result|
|book admin|change a title|the book has a typo|


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
1 service or SPI interfaces and 1 actual service implementations.

##### Type *Book*

The data class *Book* is a book to loan or rent.

##### Type *User*

The data class *User* is a library customer.

##### Type *LoanService*

The API interface *LoanService* provides stuff to loan all the things.

##### Factory *LoanService*

The API factory method *LoanServiceFactory* creates an instance.
The instance must be configured using the following options:
 * Test (...a test string.)

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

## usage

The application can be launched from the command line. One can display any available options using the *-help* flag:

```bash
booklibrary -help
```
The application can be configured using the following command line or environment options.
Currently, there are 6 options.
At first, the default value is loaded into the variable.
Afterwards the environment variable is considered and finally the command line argument takes precedence.

### FulltextSearch

The bounded context *search* declares in the layer *Core* the *bool* option **FulltextSearch** which is a flag to enable fulltext search in items.
The default value is false.
The environment variable *SEARCH_CORE_FULLTEXTSEARCH* is evaluated, if present and is only overridden by the command line argument *search-core-fulltextsearch*.

Example

```bash
export SEARCH_CORE_FULLTEXTSEARCH=false
booklibrary -search-core-fulltextsearch=false
```
### Namespace

The bounded context *search* declares in the layer *Core* the *string* option **Namespace** which is a weired option.
The default value is "some ugly stuff".
The environment variable *SEARCH_CORE_NAMESPACE* is evaluated, if present and is only overridden by the command line argument *search-core-namespace*.

Example

```bash
export SEARCH_CORE_NAMESPACE="some ugly stuff"
booklibrary -search-core-namespace="some ugly stuff"
```
### MyInt64

The bounded context *search* declares in the layer *Core* the *int64* option **MyInt64** which is an integer with 8 byte.
The default value is The default value is 0 (zero).
The environment variable *SEARCH_CORE_MYINT64* is evaluated, if present and is only overridden by the command line argument *search-core-myint64*.

Example

```bash
export SEARCH_CORE_MYINT64=0
booklibrary -search-core-myint64=0
```
### MyFloat64

The bounded context *search* declares in the layer *Core* the *float64* option **MyFloat64** which is a float with 64 bits.
The default value is 5.
The environment variable *SEARCH_CORE_MYFLOAT64* is evaluated, if present and is only overridden by the command line argument *search-core-myfloat64*.

Example

```bash
export SEARCH_CORE_MYFLOAT64=5
booklibrary -search-core-myfloat64=5
```
### MyDuration

The bounded context *search* declares in the layer *Core* the *time.Duration* option **MyDuration** which is a duration.
The default value is The default value is 0s.
The environment variable *SEARCH_CORE_MYDURATION* is evaluated, if present and is only overridden by the command line argument *search-core-myduration*.

Example

```bash
export SEARCH_CORE_MYDURATION=0s
booklibrary -search-core-myduration=0s
```
### Test

The bounded context *loan* declares in the layer *Core* the *string* option **Test** which a test string.
The default value is The default value is the empty string.
The environment variable *LOAN_CORE_TEST* is evaluated, if present and is only overridden by the command line argument *loan-core-test*.

Example

```bash
export LOAN_CORE_TEST="lorem ipsum"
booklibrary -loan-core-test="lorem ipsum"
```
