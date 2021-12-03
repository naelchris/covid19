package fetcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/naelchris/covid19/Internal/repository/covid"
)

func (h HTTPResource) QueryRequestAllData(ctx context.Context, country string) []covid.Cases {

	var (
		Cases []covid.Cases
	)

	//return error if country doesn't exists in list
	_, isExists := covid.ListCountryMapped[country]

	if !isExists {
		log.Println("[fetcherDomain][QueryRequestData] problem Do Request, err :", errors.New("invalid params"))
		return Cases
	}

	url := fmt.Sprintf("https://api.covid19api.com/dayone/country/%s", country)

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

	return Cases

}

func (h HTTPResource) QueryRequestDailyData(ctx context.Context, country string) []covid.Cases {

	var (
		Cases []covid.Cases
	)

	timeZone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return []covid.Cases{}
	}

	currentDate := time.Now()

	fromDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day()-1,
		0, 0, 0, 0, timeZone).Format(time.RFC3339)
	toDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(),
		0, 0, 0, 0, timeZone).Format(time.RFC3339)

	url := fmt.Sprintf("https://api.covid19api.com/country/%s?from=%sZ&to=%sZ", country, fromDate[:len(fromDate)-6], toDate[:len(toDate)-6])

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := h.Http.Do(req)
	if err != nil {
		log.Println("[fetcherDomain][QueryRequestDailyData] problem Do Request, err :", err)
		return Cases
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println("[fetcherDomain][QueryRequestDailyData] problem Read Body, err :", err)
		return Cases
	}

	if err := json.Unmarshal(body, &Cases); err != nil {
		log.Println("[fetcherDomain][QueryRequestDailyData] problem when umarshal json, err :", err)
		return Cases
	}

	return Cases

}
