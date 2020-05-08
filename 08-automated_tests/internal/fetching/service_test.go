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
		repo  beerscli.BeerRepo
		input int
		want  int
		err   error
	}{
		"valid beer":            {repo: buildMockBeers(), input: 127, want: 127, err: nil},
		"not found beer":        {repo: buildMockBeers(), input: 99999, err: errors.New("error")},
		"error with repository": {repo: buildMockError(), err: errors.New("error")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			service := NewService(tc.repo)

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

func TestStoreService_FetchStoreById(t *testing.T) {
	testsCases := map[string]struct {
		repo beerscli.StoreRepo
		input int
		want string
		err error
	}{
		"valid store": {repo: buildMockStores(), input: 127, want: "Mercadona", err: nil},
		"not found beer": {repo: buildMockStores(), input: 128, err: errors.New("error")},
	}

	for name, testCase := range testsCases{
		t.Run(name, func(t *testing.T) {
			service := NewStoreService(testCase.repo)
			store, err := service.FetchStoreById(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			assert.Equal(t, testCase.want, store.Name)
		})
	}
}

func buildMockStores() beerscli.StoreRepo {
	mockedRepo := &mock.StoreRepoMock{
		GetStoresFunc: func() ([]beerscli.Store, error) {
			return []beerscli.Store{
				beerscli.NewStore(127, "Mercadona"),
				beerscli.NewStore(8520130, "Carrefour"),
			}, nil
		},
	}
	return mockedRepo
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
