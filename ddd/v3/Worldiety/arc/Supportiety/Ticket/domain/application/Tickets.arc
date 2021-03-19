@doc ...encapsulates all ticket related functions.
@epic #copy(doc/en/context/Ticket/ManageTickets@doc)
@epic As a SupportietyAdmin or Application I need to manage tickets.
type Tickets stub {
    @doc ...deletes the given ticket identified by its id.
    @story #copy(doc/en/context/Ticket/ManageTickets.DeleteBySecId@doc)
    @story As a SupportietyAdmin I need to <<delete>> all tickets using
        a <user security id>, so that I can comply to the DSGVO/GDPR.
    @param id ...is the unique ticket id.
    @error NotFound ...is returned, if the expected id is not known.
    @error NotAccessible ...is returned, if the entity is not available in some form.
    fn delete(id SecId)

    @doc ...create allocates a new ticket and a unique id.
    @param appId ... is the application id.
    @param appVersion ... is the application version.
    @error CannotAllocateTicket ...is returned, if something went wrong.
    @error NotAccessible #copy(.delete.NotAccessible)
    @return TicketId ...contains the unique ticket id.
    fn create(appId ApplicationId, appVersion AppVersion) -> (TicketId)
}