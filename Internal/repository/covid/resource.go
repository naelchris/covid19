package covid

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DomainItf interface {
	AddCases(ctx context.Context, data Cases) (Cases, error)
	GetCasesByDay(ctx context.Context, country string, date time.Time, dateRange int) ([]Cases, error)
}

type DBResourceItf interface {
	AddCases(ctx context.Context, data Cases) (Cases, error)
	GetCasesByDay(ctx context.Context, country string, date time.Time, dateRange int) ([]Cases, error)
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
		log.Println("[ClassUsecase][AddClass] problem when querying to database, err :", err)
		return resp, err
	}

	return resp, nil
}

func (d Domain) GetCasesByDay(ctx context.Context, country string, date time.Time, dateRange int) ([]Cases, error) {
	resp, err := d.Storage.GetCasesByDay(ctx, country, date, dateRange)
	if err != nil {
		log.Println("[ClassUsecase][GetCasesByDay] problem when querying to database, err :", err)
		return resp, err
	}

	return resp, nil
}
