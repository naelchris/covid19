package covid

import (
	"context"
	"log"
	"time"

	covid "github.com/naelchris/covid19/Internal/repository/covid"
)

type CovidUsecase struct {
	covidDomain covid.DomainItf
}

func NewCovidUsecase(covidDomain covid.DomainItf) *CovidUsecase {
	return &CovidUsecase{
		covidDomain: covidDomain,
	}
}

func (uc *CovidUsecase) AddCases(ctx context.Context, casesData covid.Cases) (covid.Cases, error) {
	casesData.Date = time.Now()
	resp, err := uc.covidDomain.AddCases(ctx, casesData)
	if err != nil {
		log.Println("[CasesUsecase][AddCases] failt to create cases err", err)
		return resp, err
	}

	return resp, nil
}
