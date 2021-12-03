package fetcher

import (
	"context"
	"net/http"
	"time"

	"github.com/naelchris/covid19/Internal/repository/covid"
)

type DomainItf interface {
	QueryRequestAllData(ctx context.Context, country string) []covid.Cases
	QueryRequestDailyData(ctx context.Context, country string) []covid.Cases
}

type ResourceDomainItf interface {
	QueryRequestAllData(ctx context.Context, country string) []covid.Cases
	QueryRequestDailyData(ctx context.Context, country string) []covid.Cases
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

func (d Domain) QueryRequestAllData(ctx context.Context, country string) []covid.Cases {
	cases := d.HttpDomain.QueryRequestAllData(ctx, country)
	return cases
}

func (d Domain) QueryRequestDailyData(ctx context.Context, country string) []covid.Cases {
	cases := d.HttpDomain.QueryRequestDailyData(ctx, country)
	return cases
}
