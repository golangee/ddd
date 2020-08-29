# BookLibrary

This is the central service of the book library of capital city, for searching and loaning books.

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

...is a book with meta data to index and find.

##### Type *BookRepository*

...is a repository to handle books.

##### Type *SearchService*

...is the domain specific service API.

##### Factory *NewSearchService*

...is a factory to create a new SearchService.

### The context *loan*

This context is about everything around loaning or renting a book.
Only physical books can be loaned from within the library building by users.

#### The domains core layer

The core layer or API layer of the domain consists of 2 data types,
0 service or SPI interfaces and 0 actual service implementations.

##### Type *Book*

...is a book to loan or rent.

##### Type *User*

... is a library customer.

