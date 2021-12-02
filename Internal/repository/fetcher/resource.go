package fetcher

import (
	"context"
	"net/http"
	"time"

	"github.com/naelchris/covid19/Internal/repository/covid"
)

type DomainItf interface {
	QueryRequestData(ctx context.Context) []covid.Cases
}

type ResourceDomainItf interface {
	QueryRequestData(ctx context.Context) []covid.Cases
}

type Domain struct {
	HttpDomain ResourceDomainItf
}

func InitDomain() Domain {
	return Domain{
		HttpDomain: HTTPResource{
			Http: http.Client{
				Timeout: time.Second * 2, //timeout after 2 second
			},
		},
	}
}

func (d Domain) QueryRequestData(ctx context.Context) []covid.Cases {
	cases := d.HttpDomain.QueryRequestData(ctx)
	return cases
}
