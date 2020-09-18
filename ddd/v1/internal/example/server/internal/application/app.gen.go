// Code generated by golangee/architecture. DO NOT EDIT.

package application

import (
	loancore "example-server/internal/loan/core"
	loanusecase "example-server/internal/loan/usecase"
	core "example-server/internal/search/core"
	mysql "example-server/internal/search/mysql"
	usecase "example-server/internal/search/usecase"
	flag "flag"
	fmt "fmt"
	os "os"
)

// App is the actual application, which glues all layers together and is launched from the command line.
type App struct {
	dBTX           mysql.DBTX
	error          error
	bookRepository core.BookRepository
	searchService  core.SearchService
	bookSearch     usecase.BookSearch
	loanService    loancore.LoanService
	bookLoaning    loanusecase.BookLoaning
}

// Start launches any blocking background processes, like e.g. an http server.
func (a App) Start() error {
	return nil
}

// NewApp creates a new instance of the application and performs all parameter parsing and wiring.
func NewApp() (*App, error) {
	var err error
	options := Options{}
	options.Reset()
	if err := options.ParseEnv(); err != nil {
		return nil, err
	}
	options.ConfigureFlags()
	help := flag.Bool("help", false, "shows this help")
	flag.Parse()
	if *help {
		fmt.Println("BookLibrary")
		flag.PrintDefaults()
		os.Exit(0)
	}
	a := &App{}
	if a.dBTX, err = mysql.Open(options.SearchMysqlOptions); err != nil {
		return nil, err
	}

	if err = mysql.Migrate(a.dBTX); err != nil {
		return nil, err
	}

	if a.bookRepository, err = mysql.NewMysqlBookRepository(a.dBTX); err != nil {
		return nil, err
	}

	if a.searchService, err = core.NewSearchService(options.SearchCoreSearchServiceOpts, a.bookRepository); err != nil {
		return nil, err
	}

	if a.bookSearch, err = usecase.NewBookSearch(options.SearchUsecaseBookSearchOpts, a.searchService); err != nil {
		return nil, err
	}

	if a.loanService, err = loancore.NewLoanService(options.LoanCoreLoanServiceOpts); err != nil {
		return nil, err
	}

	if a.bookLoaning, err = loanusecase.NewBookLoaning(options.LoanUsecaseBookLoaningOpts, a.loanService); err != nil {
		return nil, err
	}

	return a, nil
}
