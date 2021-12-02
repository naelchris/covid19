package covid

import (
	"context"
	"errors"
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
		log.Println("[CasesUsecase][AddCases] fail to create cases err:", err)
		return resp, err
	}

	return resp, nil
}

func (uc *CovidUsecase) GetCasesByDay(ctx context.Context, country string, year int, month int, startDate int, dateRange int) (resp []covid.CasesSummary, err error) {
	timeZone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return
	}

	date := time.Date(year, time.Month(month), startDate-1, 0, 0, 0, 0, timeZone)

	caseData, err := uc.covidDomain.GetCasesByDay(ctx, country, date, dateRange+1)
	if err != nil {
		log.Println("[CasesUsecase][GetCasesByDay] fail to get cases by day err:", err)
		return
	}

	for i := range caseData {
		if i == 0 {
			continue
		}
		summary := covid.CasesSummary{
			Confirmed:         caseData[i].Confirmed,
			Deaths:            caseData[i].Deaths,
			Recovered:         caseData[i].Recovered,
			Active:            caseData[i].Active,
			IncreaseConfirmed: caseData[i].Confirmed - caseData[i-1].Confirmed,
			IncreaseDeaths:    caseData[i].Deaths - caseData[i-1].Deaths,
			IncreaseRecovered: caseData[i].Recovered - caseData[i-1].Recovered,
			IncreaseActive:    caseData[i].Active - caseData[i-1].Active,
			Date:              caseData[i].Date,
		}

		resp = append(resp, summary)
	}

	return
}

func (uc *CovidUsecase) GetCaseIncrement(ctx context.Context, country string, year int, month int, date int) (summary covid.CasesSummary, err error) {
	timeZone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return
	}

	currentDate := time.Date(year, time.Month(month), date-1, 0, 0, 0, 0, timeZone)

	resp, err := uc.covidDomain.GetCasesByDay(ctx, country, currentDate, 1)
	if err != nil {
		log.Println("[CasesUsecase][GetCaseIncrement] fail to get cases increment err:", err)
		return
	}

	//return error if cannot received today's data or yesterday's data
	if len(resp) != 2 {
		log.Println("[CasesUsecase][GetCaseIncrement] fail to get cases increment err:", err)
		return summary, errors.New("invalid params")
	}

	summary = covid.CasesSummary{
		IncreaseConfirmed: resp[1].Confirmed - resp[0].Confirmed,
		IncreaseDeaths:    resp[1].Deaths - resp[0].Deaths,
		IncreaseRecovered: resp[1].Recovered - resp[0].Recovered,
		IncreaseActive:    resp[1].Active - resp[0].Active,
		Date:              resp[1].Date,
	}

	return
}
