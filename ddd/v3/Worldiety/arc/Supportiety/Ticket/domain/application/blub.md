# Tickets

## Glossary

.An AppName is a string something like cewe fotowelt.
```go
type AppName string
```

.An AppVersion is a strict monotonic increasing version number
```go
type AppVersion int64
```

## Domain

### core

#### CompanyService
-> requires CompanyRepository

#### CompanyRepository

### Application 

#### PortfolioService
-> requires core.CompanyService

##### Story 1

As a CFWA I need to *create a ticket* using my _app name_ and my _app version_, so that I get a unique **ticket
identifier** starting with a configurable prefix depending on app name.

```go

type CreateATicket func(appName AppName, appVersion AppVersion) (TicketId, error)
```

#### CompanyService

## Adapter

## REST

### POST /v1/tickets
see [[Story 1]], returns 201, 401, 403