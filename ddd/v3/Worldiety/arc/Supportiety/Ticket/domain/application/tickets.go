package application

// Processing directives are \n+
//  - Introduced by [...]
//  - Generator creates a [...]
//  - Instance requires [...]
//  - Parameter <name> [...]
//  - Return <name|type> [...]
//  - Error <name> [...]

// CannotAllocateTicket indicates that a Ticket cannot be allocated.
//
// Generator creates a wrapping error.
type CannotAllocateTicket error


// Tickets encapsulates all ticket related functions.
//
// Introduced by doc/en/context/Ticket/ManageTickets.
// Generator creates a stub.
// Instance requires core/TicketRepository
type Tickets interface {
	// Create allocates a new ticket and a unique id.
	//
	// Introduced by doc/en/context/Ticket/ManageTickets#AdminDeletesTicket.
	// Parameter appId is the application id.
	// Parameter appVersion is the application version.
	// Return string contains the newly allocated identifier.
	// Error CannotAllocateTicket.
	Create(appId, appVersion string) (string, error)
}
