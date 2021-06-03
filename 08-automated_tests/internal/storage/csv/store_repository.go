package csv

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	beerscli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"
)

type storeRepo struct {
}

// NewCSVStoreRepo initialize csv repository
func NewCSVStoreRepo() beerscli.StoreRepo {
	return &storeRepo{}
}

// GetStores fetch stores data from csv
func (r *storeRepo) GetStores() ([]beerscli.Store, error) {
	f, _ := os.Open("08-automated_tests/data/stores.csv")
	reader := bufio.NewReader(f)

	var stores []beerscli.Store

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		storeID, _ := strconv.Atoi(values[0])

		store := beerscli.NewStore(
			storeID,
			values[1],
			values[2],
		)

		stores = append(stores, store)
	}

	return stores, nil
}
