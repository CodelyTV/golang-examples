package beerscli

type Store struct {
	StoreID    int    `json:"store_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Phone      string `json:"phone"`
}

// StoreRepo definiton of methods to access a data store
type StoreRepo interface {
	GetStores() ([]Store, error)
}

// NewStore initialize struct beer
func NewStore(storeID int, name, address, city, postalCode, phone string) (s Store) {
	s = Store{
		StoreID:    storeID,
		Name:       name,
		Address:    address,
		City:       city,
		PostalCode: postalCode,
		Phone:      phone,
	}
	return
}
