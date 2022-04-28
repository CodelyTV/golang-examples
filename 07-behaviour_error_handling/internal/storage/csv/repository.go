package csv

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	beerscli "github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal"
	"github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal/errors"
)

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() beerscli.BeerRepo {
	return &repository{}
}

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]beerscli.Beer, error) {
	const filePath = "07-behaviour_error_handling/data/beers.csv"
	f, err := os.Open(filePath)

	if err != nil {
		return nil, errors.WrapFileHandlingError(err, "error opening file %s", filePath)
	}

	reader := bufio.NewReader(f)

	var beers []beerscli.Beer

	var line []byte

	for line, err = readLine(reader); line != nil; line, err = readLine(reader) {
		if err != nil {
			return nil, err
		}
		values := strings.Split(string(line), ",")

		var productID int
		productID, err = strconv.Atoi(values[0])

		if err != nil {
			return nil, errors.WrapConversionError(err, "error while converting 'productId' from ascii to integer: %s", values[0])
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

func readLine(reader *bufio.Reader) (line []byte, err error) {
	line, _, err = reader.ReadLine()

	if err != nil {
		return nil, errors.WrapFileHandlingError(err, "could not read next line of buffer")
	}
	return
}
