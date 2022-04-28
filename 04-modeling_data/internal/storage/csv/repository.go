package csv

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	beerscli "github.com/CodelyTV/golang-examples/04-modeling_data/internal"
)

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() beerscli.BeerRepo {
	return &repository{}
}

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]beerscli.Beer, error) {
	f, error := os.Open("data/beers.csv")
	if error != nil {
		log.Fatal(error)
	}
	reader := bufio.NewReader(f)

	var beers []beerscli.Beer

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		productID, _ := strconv.Atoi(values[0])
		price, _ := strconv.ParseFloat(values[3], 64)

		beer := beerscli.NewBeer(
			productID,
			values[1],
			values[2],
			values[5],
			values[6],
			beerscli.NewBeerType(values[4]),
			price,
		)

		beers = append(beers, beer)
	}

	return beers, nil
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
