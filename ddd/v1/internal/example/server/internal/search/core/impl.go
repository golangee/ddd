package core

import (
	"context"
	"fmt"
)

type MySearchServiceImpl struct {
	SearchServiceMock
}

func (s MySearchServiceImpl) Search(ctx context.Context, query string) ([]Book, error) {
	fmt.Println("search success for: " + query)
	return nil, nil
}

func init() {
	fmt.Println("2")
	SearchServiceFactory = func(opts SearchServiceOpts, bookRepository BookRepository) (SearchService, error) {
		return MySearchServiceImpl{}, nil
	}
}
