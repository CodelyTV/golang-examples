package beerscli

// Store representation of beer into data struct
type Store struct {
	StoreID int    `json:"product_id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// StoreRepo definiton of methods to access a data beer
type StoreRepo interface {
	GetStores() ([]Store, error)
}

// NewStore initialize struct beer
func NewStore(productID int, name, category, brewer, country, price string) (b Store) {
	b = Store{
		StoreID: productID,
		Name:    name,
		Country: country,
	}
	return
}
