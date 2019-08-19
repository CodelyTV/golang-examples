package inmem

import (
	beerscli "github.com/CodelyTV/golang-examples/09-benchmarking/internal"
)

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() beerscli.BeerRepo {
	return &repository{}
}

// GetBeers fetch beers data from memory
func (r *repository) GetBeers() ([]beerscli.Beer, error) {
	return []beerscli.Beer{
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Domestic Specialty",
			"Molson",
			"Canada",
			"23.95",
			beerscli.NewBeerType("Lager"),
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Grolsch Export B.V.",
			"Canada",
			"49.50",
			beerscli.NewBeerType("Non-Alcoholic Beer"),
		),
	}, nil
}
