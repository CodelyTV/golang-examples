package fetching_test

func (m *MyMockedObject) GetBeers(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}
