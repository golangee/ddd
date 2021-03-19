package v1

// Tickets contains all ticket related endpoints.
//
// Generator creates a REST service.
// Path is v1.
// Instance requires <Supportiety/Ticket/domain/application/>Tickets.
type Tickets interface {
	// Method ist POST.
	// Path is tickets.
	// Require header["jwt"] as jwt.
	// Require header["app_id"] as appId.
	// Optional header["app_version"] as appVersion.
	// Parameter jwt contains a json web token.
	// Parameter appId #copy(application/Tickets.Create#Parameter appId).
	// Parameter appVersion #copy(application/Tickets.Create#Parameter appVersion).
	// Return 201, 401, 403, 400.
	// Return string #copy(application/Tickets.Create#Return string).
	Create(jwt, appId, appVersion string) string

	// Method is DELETE.
	// Path is tickets/:ticket_id.
	// Require path["ticket_id"] as ticketId.
	Delete(ticketId string)
}
