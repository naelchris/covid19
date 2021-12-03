package covid

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type DomainItf interface {
	AddCases(ctx context.Context, data Cases) (Cases, error)
	AddCasesBulk(ctx context.Context, data []Cases) error
	GetCasesByDay(ctx context.Context, country string, date time.Time, dateRange int) ([]Cases, error)
	GetCasesByMonths(ctx context.Context, country string, year int, startMonth int, monthRange int) ([]CasesSummary, error)
}

type DBResourceItf interface {
	AddCases(ctx context.Context, data Cases) (Cases, error)
	AddCasesBulk(ctx context.Context, data []Cases) error
	GetCasesByDay(ctx context.Context, country string, date time.Time, dateRange int) ([]Cases, error)
	GetCasesByMonth(ctx context.Context, country string, startDate string, endDate string) (CasesSummary, error)
}

type Domain struct {
	Storage DBResourceItf
}

func InitDomain(db *sql.DB) Domain {
	return Domain{
		Storage: storage{
			CasesDB: db,
		},
	}
}

func (d Domain) AddCases(ctx context.Context, data Cases) (Cases, error) {
	resp, err := d.Storage.AddCases(ctx, data)
	if err != nil {
		log.Println("[CovidDomain][AddCases] problem when querying to database, err :", err)
		return resp, err
	}

	return resp, nil
}

func (d Domain) AddCasesBulk(ctx context.Context, data []Cases) error {
	err := d.Storage.AddCasesBulk(ctx, data)
	if err != nil {
		log.Println("[CovidDomain][AddCasesBulk] problem when querying to database, err :", err)
		return err
	}

	return nil
}

func (d Domain) GetCasesByDay(ctx context.Context, country string, date time.Time, dateRange int) ([]Cases, error) {
	resp, err := d.Storage.GetCasesByDay(ctx, country, date, dateRange)
	if err != nil {
		log.Println("[ClassUsecase][GetCasesByDay] problem when querying to database, err :", err)
		return resp, err
	}

	return resp, nil
}

//TODO : add parameter month, monthrange
func (d Domain) GetCasesByMonths(ctx context.Context, country string, year int, startMonth int, monthRange int) ([]CasesSummary, error) {
	var (
		result []CasesSummary
	)

	timeZone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return []CasesSummary{}, err
	}

	for i := startMonth; i <= startMonth+monthRange+1; i++ {
		startTime := time.Date(year, time.Month(i), 0, 0, 0, 0, 0, timeZone)
		//parse
		fromDate := fmt.Sprintf("%d-%02d-%02d", startTime.Year(), startTime.Month(), startTime.Day())

		cases, err := d.Storage.GetCasesByMonth(ctx, country, fromDate, fromDate)
		if err != nil {
			log.Println("[GetCasesByMonth Domain][GetCasesByMonths] err, ", err)
			return result, err
		}

		cases.Date = startTime

		//append result
		result = append(result, cases)
	}

	return result, nil
}
