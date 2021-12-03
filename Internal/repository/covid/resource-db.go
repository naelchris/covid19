package covid

import (
	"context"
	"database/sql"
	"log"
	"time"

	sqlbuilder "github.com/huandu/go-sqlbuilder"
	sqlx "github.com/jmoiron/sqlx"
)

func (s storage) AddCases(ctx context.Context, data Cases) (Cases, error) {
	var resp Cases
	//	var id int64

	log.Println("[ClassRepository][ResourceDB][addClass] Data Class,", data)

	//prepare
	qr, err := s.CasesDB.Prepare(addCaseQuery)
	if err != nil {
		log.Println("[ClassRepository][ResourceDB][addClass] prepare failed err, ", err.Error())
		return resp, err
	}

	err = qr.QueryRow(
		data.Country,
		data.CountryCode,
		data.Province,
		data.City,
		data.CityCode,
		data.Lat,
		data.Lon,
		data.Confirmed,
		data.Deaths,
		data.Recovered,
		data.Active,
		data.Date,
	).Scan(&resp.ID, &resp.Country, &resp.CountryCode, &resp.Province, &resp.City, &resp.CityCode,
		&resp.Lat, &resp.Lon, &resp.Confirmed, &resp.Deaths, &resp.Recovered, &resp.Active, &resp.Date)
	if err != nil {
		log.Println("[ClassRepository][ResourceDB][addClass] problem query to db err", err.Error())
		return resp, err
	}

	return resp, nil
}

func (s storage) AddCasesBulk(ctx context.Context, data []Cases) (err error) {
	//build add cases bulk query
	query, args := s.buildAddCasesBulkQuery(data)

	query += ` ON CONFLICT (date, country) DO UPDATE SET confirmed = EXCLUDED.confirmed, deaths = EXCLUDED.deaths, recovered = EXCLUDED.recovered,  active = EXCLUDED.active`

	//bind type postgres -> 2
	_, err = s.CasesDB.Exec(sqlx.Rebind(2, query), args...)
	if err != nil {
		log.Println("[ClassRepository][ResourceDB][addClassBulk] problem query to db err", err.Error())
		return err
	}

	return
}

func (s storage) buildAddCasesBulkQuery(data []Cases) (string, []interface{}) {
	insertBuilder := sqlbuilder.NewInsertBuilder()

	insertBuilder.InsertInto("covid19_data")
	insertBuilder.Cols(addCaseColumn)

	for _, rows := range data {
		insertBuilder.Values(
			rows.Country,
			rows.CountryCode,
			rows.Province,
			rows.City,
			rows.CityCode,
			rows.Lat,
			rows.Lon,
			rows.Confirmed,
			rows.Deaths,
			rows.Recovered,
			rows.Active,
			rows.Date,
		)
	}

	return insertBuilder.Build()
}

func (s storage) GetCasesByDay(ctx context.Context, country string, date time.Time, dateRange int) (resp []Cases, err error) {
	log.Println("[ClassRepository][ResourceDB][GetCasesByDay] Country:", country, "Date:", date, ", DateRange:", dateRange)

	fromDate := date.Format("2006-01-02")
	toDate := date.AddDate(0, 0, dateRange).Format("2006-01-02")

	//get cases row
	rows, err := s.CasesDB.Query(getCasesByDay, country, fromDate, toDate)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[ClassRepository][ResourceDB][GetCasesByDay] problem query to db err", err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var respRow Cases

		err = rows.Scan(&respRow.Confirmed, &respRow.Deaths, &respRow.Recovered, &respRow.Active, &respRow.Date)
		if err != nil {
			return
		}

		resp = append(resp, respRow)
	}

	return
}

func (s storage) GetCasesByMonth(ctx context.Context, country string, startDate string, endDate string) (CasesSummary, error) {
	//create query for month

	var resp CasesSummary

	//prepare query
	qr, err := s.CasesDB.Prepare(filterMonthCasesQuery)
	if err != nil {
		log.Println("[Prepare][ResourceDB][GetCasesByMonth] prepare failed err, ", err.Error())
		return CasesSummary{}, err
	}

	err = qr.QueryRow(country, startDate, endDate).Scan(&resp.Confirmed, &resp.Deaths, &resp.Recovered, &resp.Active)
	if err != nil {
		log.Println("[QueryRow][ResourceDB][GetCasesByMonth] Query row failed err, ", err.Error())
		return CasesSummary{}, err
	}

	return resp, nil

}
