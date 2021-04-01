package mock

import (
	beerscli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"
	"github.com/stretchr/testify/mock"
)

type MockStoreRepo struct {
	mock.Mock
}

func (m *MockStoreRepo) GetStores() ([]beerscli.Store, error) {
	args := m.Called()
	res := args.Get(0)
	return res.([]beerscli.Store), args.Error(1)
}
