package beerscli

// Beer representation of beer into data struct
type Beer struct {
	ProductID int
	Name      string
	Price     float64
	Category  string
	Type      *BeerType
	Brewer    string
	Country   string
}

type BeerType int

const (
	Unknown BeerType = iota
	Lager
	Malt
	Ale
	FlavouredMalt
	Stout
	Porter
	NonAlcoholic
)

func (t BeerType) String() string {
	return toString[t]
}

// NewBeerType initialize a type from enum beerTypes
func NewBeerType(t string) *BeerType {
	beerType := toID[t]
	return &beerType
}

var toString = map[BeerType]string{
	Lager:         "Lager",
	Malt:          "Malt",
	Ale:           "Ale",
	FlavouredMalt: "Flavoured Malt",
	Stout:         "Stout",
	Porter:        "Stout",
	NonAlcoholic:  "Non-Alcoholic",
	Unknown:       "unknown",
}

var toID = map[string]BeerType{
	"Lager":          Lager,
	"Malt":           Malt,
	"Ale":            Ale,
	"Flavoured Malt": FlavouredMalt,
	"Stout":          Stout,
	"Porter":         Porter,
	"Non-Alcoholic":  NonAlcoholic,
	"unknown":        Unknown,
}

// BeerRepo definiton of methods to access a data beer
type BeerRepo interface {
	GetBeers() ([]Beer, error)
}

// NewBeer initialize struct beer
func NewBeer(productID int, name, category, brewer, country string, beerType *BeerType, price float64) (b Beer) {
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
