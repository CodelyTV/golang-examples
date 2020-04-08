package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	storescli "github.com/CodelyTV/golang-examples/08-automated_tests/internal"
	"github.com/CodelyTV/golang-examples/08-automated_tests/internal/errors"
)

const (
	productsEndpointStore = "/stores"
)

// NewOntarioStoreRepository fetch stores data from api
func NewOntarioStoreRepository() storescli.StoreRepo {
	return &beerRepo{url: ontarioURL}
}

func (b *beerRepo) GetStores() (stores []storescli.Store, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, productsEndpointStore))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", productsEndpointStore)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading the response from %s", productsEndpointStore)
	}

	err = json.Unmarshal(contents, &stores)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into stores")
	}
	return
}
