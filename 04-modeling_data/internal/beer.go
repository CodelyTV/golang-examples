package beerscli

// Beer representation of beer into data struct
type Beer struct {
	ProductID int
	Name      string
	Price     float64
	Category  string
	Type      string
	Brewer    string
	Country   string
}

// NewBeer initialize struct beer
func NewBeer(productID int, name, category, beerType, brewer, country string, price float64) (b Beer) {
	b = Beer{
		ProductID: productID,
		Name:      name,
		Category:  category,
		Type:      beerType,
		Brewer:    brewer,
		Country:   country,
		Price:     price,
	}
	return
}
