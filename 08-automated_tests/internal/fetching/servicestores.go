package fetching

import (
	storescli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"
	"github.com/pkg/errors"
)

// ServiceStores provides store fetching operations
type ServiceStores interface {
	// FetchStores fetch all stores from repository
	FetchStores() ([]storescli.Store, error)
	// FetchByID filter all stores and get only the store that match with given id
	FetchByID(id int) (storescli.Store, error)
}

type serviceStore struct {
	sR storescli.StoreRepo
}

// NewServiceStore creates an adding service with the necessary dependencies
func NewServiceStore(r storescli.StoreRepo) ServiceStores {
	return &serviceStore{r}
}

func (s *serviceStore) FetchStores() ([]storescli.Store, error) {
	return s.sR.GetStores()
}

func (s *serviceStore) FetchByID(id int) (storescli.Store, error) {
	stores, err := s.FetchStores()
	if err != nil {
		return storescli.Store{}, err
	}

	for _, store := range stores {
		if store.StoreID == id {
			return store, nil
		}
	}

	return storescli.Store{}, errors.Errorf("Store %d not found", id)
}
