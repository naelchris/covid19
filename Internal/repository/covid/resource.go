package covid

import (
	"context"
	"database/sql"
	"log"
)

type DomainItf interface {
	AddCases(ctx context.Context, data Cases) (Cases, error)
}

type DBResourceItf interface {
	AddCases(ctx context.Context, data Cases) (Cases, error)
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
