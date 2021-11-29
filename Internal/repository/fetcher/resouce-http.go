package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/naelchris/covid19/Internal/repository/covid"
)

func (h HTTPResource) QueryRequestData(ctx context.Context) []covid.Cases {

	var (
		Cases []covid.Cases
	)

	url := "https://api.covid19api.com/dayone/country/indonesia"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := h.Http.Do(req)
	if err != nil {
		log.Println("[fetcherDomain][QueryRequestData] problem Do Request, err :", err)
		return Cases
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println("[fetcherDomain][QueryRequestData] problem Read Body, err :", err)
		return Cases
	}

	if err := json.Unmarshal(body, &Cases); err != nil {
		log.Println("[fetcherDomain][QueryRequestData] problem when umarshal json, err :", err)
		return Cases
	}

	fmt.Println(Cases)

	return Cases

}
