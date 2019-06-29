package csv

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	beerscli "github.com/CodelyTV/golang-introduction/08-automated_tests/internal"
)

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() beerscli.BeerRepo {
	return &repository{}
}

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]beerscli.Beer, error) {
	f, _ := os.Open("08-automated_tests/data/beers.csv")
	reader := bufio.NewReader(f)

	var beers []beerscli.Beer

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		productID, _ := strconv.Atoi(values[0])

		beer := beerscli.NewBeer(
			productID,
			values[1],
			values[2],
			values[5],
			values[6],
			values[3],
			beerscli.NewBeerType(values[4]),
		)

		beers = append(beers, beer)
	}

	return beers, nil
}

func (r *repository) GetStores() ([]beerscli.Store, error) {
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
			values[3],
			values[4],
			values[5],
			values[6],
			values[7],
			values[8],
			values[9],
			values[10],
			values[11],
			values[12],
			values[13],
			values[14],
			values[15],
			values[16],
			values[17],
			values[18],
			values[19],
			values[20],
		)

		stores = append(stores, store)
	}

	return stores, nil
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
