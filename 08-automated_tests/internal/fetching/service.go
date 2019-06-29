package fetching

import (
	beerscli "github.com/CodelyTV/golang-introduction/08-automated_tests/internal"
	"github.com/pkg/errors"
)

// Service provides beer fetching operations
type Service interface {
	// FetchBeers fetch all beers from repository
	FetchBeers() ([]beerscli.Beer, error)
	// FetchByID filter all beers and get only the beer that matches with given id
	FetchByID(id int) (beerscli.Beer, error)
}

type StoresService interface {
	// FetchStores fetch all stores from repository
	FetchStores() ([]beerscli.Store, error)
	// FetchStoreById filter all stores and get only the store that matches with given id
	FetchStoreByID(id int) (beerscli.Store, error)
}

type service struct {
	bR beerscli.BeerRepo
}

type storeService struct {
	sR beerscli.StoreRepo
}

// NewService creates an adding service with the necessary dependencies
func NewService(r beerscli.BeerRepo) Service {
	return &service{r}
}

// NewStoreService creates an adding store service with the necessary dependencies
func NewStoreService(r beerscli.StoreRepo) StoresService {
	return &storeService{r}
}

func (s *storeService) FetchStores() ([]beerscli.Store, error) {
	return s.sR.GetStores()
}

func (s *storeService) FetchStoreByID(id int) (beerscli.Store, error) {
	stores, err := s.FetchStores()
	if err != nil {
		return beerscli.Store{}, err
	}

	for _, store := range stores {
		if store.StoreID == id {
			return store, nil
		}
	}

	return beerscli.Store{}, errors.Errorf("Store %d not found", id)
}

func (s *service) FetchBeers() ([]beerscli.Beer, error) {
	return s.bR.GetBeers()
}

func (s *service) FetchByID(id int) (beerscli.Beer, error) {
	beers, err := s.FetchBeers()
	if err != nil {
		return beerscli.Beer{}, err
	}

	for _, beer := range beers {
		if beer.ProductID == id {
			return beer, nil
		}
	}

	return beerscli.Beer{}, errors.Errorf("Beer %d not found", id)
}
