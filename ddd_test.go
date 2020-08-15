package ddd_test

import (
	. "github.com/golangee/ddd"
	"testing"
)

func TestDDD(t *testing.T) {
	err := Application("MySuperApplication",
		Domains(
			Domain("Dashboards",
				"... is about users and their dashboards.",
				Persistence(
					Repositories(
						Interface("UserRepository",
							Method("FindAll",
								"...returns all entries.",
								In(),
								Out(),
							),
							Method("FindOne",
								"...returns the first matching entry or an error",
								In(),
								Out(),
							),
						),
						Interface("BookStore",
							Method("ReadPDF",
								"...opens a book to read",
								In(Param("file", "...is the filename", String)),
								Out(
									Return(Reader),
									Return(Error),
								),
							),
						),
						Interface("DashboardRepository",
							Method("FindAll",
								"...returns all entries.",
								In(
									Param("offset", "...the index to return from", Int64),
									Param("limit", "...returns at most", Int64),
								),
								Out(
									Param("names", "all names", List(String)),
									Param("err", "returned if something went wrong", Error),
								),
							),
							Method("FindOne",
								"...returns the first matching entry or an error",
								In(
									Param(
										"id",
										"the id of entity to find",
										UUID,
									),
								),
								Out(
									Return(List(String)),
								),
							),
						),
						Interface("DeviceRepository",
							Method("Count",
								"...enumerates all entries.",
								In(),
								Out(Return(Int64)),
							),
							Method("FindOne",
								"...returns the first matching entry or an error",
								In(),
								Out(
									Return(Int64),
									Return(Error),
								),
							),
						),
					),
					Types(
						Type("Device",
							Fields(
								Field("Id", "...is unique per device.", UUID),
								Field("Name", "...is an arbitrary non unique name.", String),
								Field("power", "...is the power consumption in Ah", Int64),
							),
						),
					),
					Implementations(
						MySQL("BookStore",
							Schema(
								Migrate(20200815153330,
									`CREATE TABLE books (id binary(16), name VARCHAR(255)) PRIMARY KEY (id);
										-- and many more stuff

										`,

								),
								Migrate(20200815153330, "ALTER TABLE books ADD COLUMN isbn BIGINT"),
							),
							Implement("FindOne", Statement("SELECT id, name FROM books WHERE id=:id")),
							Implement("Insert", DefaultCreate("books")),
							Implement("Remove", DefaultDelete("books")),
						),
						Filesystem("BookStore"),
					),
				),
			),

			Domain("Portfolios",
				"... is about Portfolio management.",
				Persistence(
					Repositories(
						Interface("PortfolioRepository",
							Method("FindAll",
								"...returns all entries.",
								In(),
								Out(Return("Portfolio")),
							),
							Method("FindOne",
								"...returns the first matching entry or an error",
								In(),
								Out(Return("Portfolio")),
							),
						),
					),
					Types(
						Type("Portfolio",
							Fields(
								Field("Id", "unique id", UUID),
								Field("Name", "human readable string", String),
							),
						),
					),
					Implementations(),
				),
			),

		),
	).Generate()

	if err != nil {
		t.Fatal(err)
	}
}
