package mysql

// TODO this is declared in core and must be imported correctly
type Ticket struct {
	ID, Firstname, Lastname, PwdHash, Role string
}

// RepoImpl provides a mysql based repository adapter implementation.
//
// Generator creates a mysql implementation.
// Implements TicketRepository.
//
//arc:generate(adapter/mysql)
//arc:implements(core/TicketRepository)
type RepoImpl interface {

	// Query SELECT id, firstname, lastname, phash, role FROM users LIMIT ? OFFSET ?.
	// Prepare limit, offset.
	// Row Ticket.ID, Ticket.Firstname, Ticket.Lastname, Ticket.PwdHash, Ticket.Role.
	//
	//arc:sql.query(SELECT id, firstname, lastname, phash, role FROM users LIMIT ? OFFSET ?)
	//arc:sql.prepare(limit, offset)
	//arc:sql.row(Ticket.ID, Ticket.Firstname, Ticket.Lastname, Ticket.PwdHash, Ticket.Role)
	//
	//@sql:query
	FindAll(limit, offset int) ([]Ticket, error)
}
