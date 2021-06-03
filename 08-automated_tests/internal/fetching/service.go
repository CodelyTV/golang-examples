package fetching

import (
	beerscli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"
)

// Service provides beer fetching operations
type Service interface {
	// FetchBeers fetch all beers from repository
	FetchBeers() ([]beerscli.Beer, error)
	// FetchByID filter all beers and get only the beer that match with given id
	FetchByID(id int) (beerscli.Beer, error)

	// FetchStores fetch all stores from repository
	FetchStores() ([]beerscli.Store, error)
	// FetchStoreByID filter all stores and get only the store that match with given id
	FetchStoreByID(id int) (beerscli.Store, error)
}

type service struct {
	bR        beerscli.BeerRepo
	storeRepo beerscli.StoreRepo
}

// NewService creates an adding service with the necessary dependencies
func NewService(beerRepo beerscli.BeerRepo, storeRepo beerscli.StoreRepo) Service {
	return &service{beerRepo, storeRepo}
}
