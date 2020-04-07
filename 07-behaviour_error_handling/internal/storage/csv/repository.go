package csv

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	beerscli "github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal"
	"github.com/golang-examples/07-behaviour_error_handling/internal/errors"
)

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() beerscli.BeerRepo {
	return &repository{}
}

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]beerscli.Beer, error) {
	f, err := os.Open("07-behaviour_error_handling/data/beers.csv")

	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error opening csv file")
	}

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

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, err = reader.ReadLine()

	if err != nil {
		return nil, errors.WrapDataUnreacheable(err,"error reading the response from %s", reader))
	}

	return
}
