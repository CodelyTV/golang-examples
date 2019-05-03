package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CodelyTV/golang-introduction/06-error_handling/internal/errors"

	beerscli "github.com/CodelyTV/golang-introduction/06-error_handling/internal"
)

const (
	productsEndpoint = "/products"
	ontarioURL       = "http://ontariobeerapi.ca"
)

type beerRepo struct {
	url string
}

// NewOntarioRepository fetch beers data from csv
func NewOntarioRepository() beerscli.BeerRepo {
	return &beerRepo{url: ontarioURL}
}

func (b *beerRepo) GetBeers() (beers []beerscli.Beer, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, productsEndpoint))
	if err != nil {
		return nil, &errors.BadResponseErr{"Something happened when call the endpoint", "ontario/repository.go", 31}
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &errors.BadResponseErr{"Something happened when read the content", "ontario/repository.go", 36}
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, &errors.BadResponseErr{"Something happened when unmarshal the content", "ontario/repository.go", 41}
	}
	return
}
