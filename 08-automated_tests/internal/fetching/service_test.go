package fetching

import (
	"errors"
	"testing"

	beerscli "github.com/CodelyTV/golang-introduction/08-automated_tests/internal"
	//"github.com/CodelyTV/golang-introduction/08-automated_tests/internal/storage/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestFetchStoreByID(t *testing.T) {

	tests := map[string]struct {
		repo  beerscli.StoreRepo
		input int
		want  int
		err   error
	}{
		"valid store":                 {repo: buildMockStores(), input: 200, want: 200, err: nil},
		"not found store":             {repo: buildMockStores(), input: 99999, err: errors.New("error")},
		"error with store repository": {repo: buildMockStores(), err: errors.New("error")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			service := NewStoreService(tc.repo)

			b, err := service.FetchStoreByID(tc.input)

			if tc.err != nil {
				assert.Error(t, err)
			}

			if tc.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.want, b.StoreID)
		})
	}
}

func buildMockStores() beerscli.StoreRepo {
	mockedRepo := new(storeRepoMock)
	return mockedRepo
}

func (m *storeRepoMock) GetStores() ([]beerscli.Store, error) {
	store := []beerscli.Store{
		beerscli.NewStore(
			200,
			"Store",
			"Mi casa",
			"Madrid",
			"28982",
			"10", "10",
			"9", "6",
			"9", "6",
			"9", "6",
			"9", "6",
			"9", "6",
			"9", "6",
			"9", "6"),
	}

	return store, nil
}

func buildMockBeers() beerscli.BeerRepo {
	mockedRepo := new(beerRepoMock)
	return mockedRepo
}

func (b *beerRepoMock) GetBeers() ([]beerscli.Beer, error) {

	beers := []beerscli.Beer{
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
	}

	return beers, nil
}

type storeRepoMock struct {
	mock.Mock
}

type beerRepoMock struct {
	mock.Mock
}

func buildMockError() beerscli.BeerRepo {
	mockedRepo := new(beerRepoMock)
	mockedRepo.On("GetBeers", errors.New("error"))
	return mockedRepo
}
