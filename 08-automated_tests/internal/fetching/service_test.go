package fetching

import (
	"errors"
	"testing"

	beerscli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"
	"github.com/CodelyTV/golang-examples/08-automated_tests/internal/storage/mock"

	"github.com/stretchr/testify/assert"
)

func TestFetchByID(t *testing.T) {

	tests := map[string]struct {
		beerRepo  beerscli.BeerRepo
		storeRepo beerscli.StoreRepo
		input     int
		want      int
		err       error
	}{
		"valid beer":            {beerRepo: buildMockBeers(), input: 127, want: 127, err: nil},
		"not found beer":        {beerRepo: buildMockBeers(), input: 99999, err: errors.New("error")},
		"error with repository": {beerRepo: buildMockError(), err: errors.New("error")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			service := NewService(tc.beerRepo, tc.storeRepo)

			b, err := service.FetchByID(tc.input)

			if tc.err != nil {
				assert.Error(t, err)
			}

			if tc.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.want, b.ProductID)
		})
	}
}

func buildMockBeers() beerscli.BeerRepo {
	mockedRepo := &mock.BeerRepoMock{
		GetBeersFunc: func() ([]beerscli.Beer, error) {
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
		},
	}

	return mockedRepo
}

func buildMockError() beerscli.BeerRepo {
	mockedRepo := &mock.BeerRepoMock{
		GetBeersFunc: func() ([]beerscli.Beer, error) {
			return []beerscli.Beer{}, errors.New("error")
		},
	}

	return mockedRepo
}

func TestGetStores(t *testing.T) {
	storeRepo := new(mock.MockStoreRepo)
	stores := []beerscli.Store{
		{
			StoreID: 1,
			Name:    "Paceña",
			Country: "BOB",
		},
	}
	storeRepo.On("GetStores").Return(stores, nil)

	service := NewService(nil, storeRepo)

	stores, err := service.FetchStores()

	assert.Nil(t, err)

	assert.Equal(t, 1, stores[0].StoreID)
	assert.Equal(t, "Paceña", stores[0].Name)
	assert.Equal(t, "BOB", stores[0].Country)
}

func TestFetchStoreByID(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
		err   error
	}{
		"valid beer":     {input: 127, want: 127, err: nil},
		"not found beer": {input: 99999, err: errors.New("error")},
		//"error with repository": {err: errors.New("error")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			storeRepo := new(mock.MockStoreRepo)
			stores := []beerscli.Store{
				{
					StoreID: 1,
					Name:    "Paceña",
					Country: "BOB",
				},
			}
			storeRepo.On("GetStores").Return(stores, nil)

			service := NewService(nil, storeRepo)

			stores, err := service.FetchStores()

			assert.Equal(t, tc.err, err)
			assert.Equal(t, "Paceña", stores[0].Name)

		})
	}
}
