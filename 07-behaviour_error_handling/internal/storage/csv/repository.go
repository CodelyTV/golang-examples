package csv

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	beerscli "github.com/CodelyTV/golang-introduction/07-behaviour_error_handling/internal"
	"github.com/CodelyTV/golang-introduction/07-behaviour_error_handling/internal/errors"
)

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() beerscli.BeerRepo {
	return &repository{}
}

const fileName string = "07-behaviour_error_handling/data/beers.csv"

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]beerscli.Beer, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error opening file with name %s", fileName)
	}

	reader := bufio.NewReader(f)

	var beers []beerscli.Beer

	for line, err := readLine(reader); line != nil; line, err = readLine(reader) {

		if err != nil {
			return nil, errors.WrapDataUnreacheable(err, "there was problems while reading lines %s", fileName)
		}

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

func readLine(reader *bufio.Reader) (line []byte, err error) {
	line, _, err = reader.ReadLine()
	return
}
