post /tickets/
    require domain/application/Tickets.create
    body:
        @doc #copy(domain/application/Tickets.create@param appId)
        @required
        app_id string

        appVersion int
    returns:
        sessionId string
        201: created
        401: unauthorized
        403: forbidden
        400: bad request

 # implementation is up to developer