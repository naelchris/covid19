package covid

import (
	"context"
	"database/sql"
	"log"
)

type DomainItf interface {
	AddCases(ctx context.Context, data InsertCasesRequest) (CasesResponse, error)
}

type DBResourceItf interface {
	AddCases(ctx context.Context, data InsertCasesRequest) (CasesResponse, error)
}

type Domain struct {
	Storage DBResourceItf
}

func InitDomain(db *sql.DB) Domain {
	return Domain{
		Storage: storage{
			ClassDB: db,
		},
	}
}

func (d Domain) AddClass(ctx context.Context, data InsertCasesRequest) (CasesResponse, error) {
	resp, err := d.Storage.AddCases(ctx, data)
	if err != nil {
		log.Println("[ClassUsecase][AddClass] problem when querying to database, err :", err)
		return resp, err
	}

	return resp, nil
}
