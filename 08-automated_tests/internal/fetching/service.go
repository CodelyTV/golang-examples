package fetching

import (
	beerscli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"
	"github.com/pkg/errors"
)

// Service provides beer fetching operations
type Service interface {
	// FetchBeers fetch all beers from repository
	FetchBeers() ([]beerscli.Beer, error)
	// FetchByID filter all beers and get only the beer that match with given id
	FetchByID(id int) (beerscli.Beer, error)
}

type StoreService interface {
	FetchStores() ([] beerscli.Store, error)
	FetchStoreById(id int) (beerscli.Store, error)
}

type service struct {
	bR beerscli.BeerRepo
}

type storeService struct {
	storeRepo beerscli.StoreRepo
}

// NewService creates an adding service with the necessary dependencies
func NewService(r beerscli.BeerRepo) Service {
	return &service{r}
}

func NewStoreService(r beerscli.StoreRepo) StoreService {
	return &storeService{r}
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

func (service *storeService) FetchStores() ([]beerscli.Store, error) {
	return service.storeRepo.GetStores()
}

func (service *storeService) FetchStoreById(id int) (beerscli.Store, error) {
	stores, err := service.storeRepo.GetStores()
	if err != nil {
		return beerscli.Store{}, err
	}

	for _, store := range stores {
		if store.ProductId == id {
			return store, nil
		}
	}

	return beerscli.Store{}, errors.Errorf("Beer %d not found, so not sold in any store", id)
}
