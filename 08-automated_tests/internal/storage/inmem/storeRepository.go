package inmem

import beerscli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"

type storeRepository struct {
}

func NewStoreRepository() beerscli.StoreRepo {
	return &storeRepository{}
}

func (s storeRepository) GetStores() ([]beerscli.Store, error) {
	store1 := beerscli.NewStore(
		127,
		 "Mercadona",
		)
	store2 := beerscli.NewStore(
		8520130,
		"Carrefour",
		)
	return []beerscli.Store{store1, store2}, nil
}
