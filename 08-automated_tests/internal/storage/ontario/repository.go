package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beerscli "github.com/CodelyTV/golang-introduction/08-automated_tests/internal"
	"github.com/CodelyTV/golang-introduction/08-automated_tests/internal/errors"
)

const (
	productsEndpoint = "/products"
	storesEndpoint   = "/stores"
	ontarioURL       = "http://ontariobeerapi.ca"
)

type beerRepo struct {
	url string
}

type storeRepo struct {
	url string
}

// NewOntarioRepository fetch beers data from csv
func NewOntarioRepository() (beerscli.BeerRepo, beerscli.StoreRepo) {
	return &beerRepo{url: ontarioURL}, &storeRepo{url: ontarioURL}
}

func (b *beerRepo) GetBeers() (beers []beerscli.Beer, err error) {
	contents, err := get(b.url, productsEndpoint)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return
}

func (b *storeRepo) GetStores() (stores []beerscli.Store, err error) {
	contents, err := get(b.url, storesEndpoint)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &stores)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into stores")
	}
	return
}

func get(url, endpoint string) ([]byte, error) {
	response, err := http.Get(fmt.Sprintf("%v%v", url, endpoint))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", endpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading the response from %s", endpoint)
	}

	return contents, nil
}
