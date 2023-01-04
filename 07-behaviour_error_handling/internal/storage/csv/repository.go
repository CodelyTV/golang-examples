package csv

import (
	"bufio"
	"github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal/errors"
	"os"
	"strconv"
	"strings"

	beerscli "github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal"
)

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() beerscli.BeerRepo {
	return &repository{}
}

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]beerscli.Beer, error) {
	f, err := os.Open("07-behaviour_error_handling/data/beers2.csv")

	if err != nil {
		return nil, errors.WrapFileDataUError(err, "error getting beers to %s", "csv data")
	}

	reader := bufio.NewReader(f)

	var beers []beerscli.Beer

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		productID, err := strconv.Atoi(values[0])

		if err != nil {
			return nil, errors.WrapFormatDataError(err, "can't format product Id with Atoi method")
		}

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
	line, _, _ = reader.ReadLine()

	return
}
