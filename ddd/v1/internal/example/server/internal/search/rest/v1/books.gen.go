// Code generated by golangee/architecture. DO NOT EDIT.

package rest

import (
	http "net/http"
)

// BooksGetContext provides the specific http request and response context including already parsed parameters.
type BooksGetContext struct {
	// Request contains the raw http request.
	Request *http.Request
	// Writer contains a reference to the raw http response writer.
	Writer http.ResponseWriter
	// Session contains the parsed header parameter for 'session'.
	Session string
}

// BooksDeleteContext provides the specific http request and response context including already parsed parameters.
type BooksDeleteContext struct {
	// Request contains the raw http request.
	Request *http.Request
	// Writer contains a reference to the raw http response writer.
	Writer http.ResponseWriter
	// Session contains the parsed header parameter for 'session'.
	Session string
}

// Books represents the REST resource api/v1/books.
// Resource to manage books.
type Books interface {
	// Get represents the http GET request on the /books resource.
	// Returns all books.
	Get(ctx BooksGetContext) error
	// Delete represents the http DELETE request on the /books resource.
	// Removes all books.
	Delete(ctx BooksDeleteContext) error
}
