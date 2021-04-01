package fetching

import (
	beerscli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"
	"github.com/pkg/errors"
)

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
