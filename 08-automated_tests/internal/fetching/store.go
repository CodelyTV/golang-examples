package fetching

import (
	beerscli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"
	"github.com/pkg/errors"
)

func (s *service) FetchStores() ([]beerscli.Store, error) {
	return s.storeRepo.GetStores()
}

func (s *service) FetchStoreByID(id int) (beerscli.Store, error) {
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
