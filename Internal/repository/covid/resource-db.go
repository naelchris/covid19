package covid

import (
	"context"
	"log"
)

func (s storage) AddCases(ctx context.Context, data Cases) (Cases, error) {
	var resp Cases
	//	var id int64

	log.Println("[ClassRepository][ResourceDB][addClass] Data Class,", data)

	//prepare
	qr, err := s.CasesDB.Prepare(addCaseQuery)
	if err != nil {
		log.Fatalln("[ClassRepository][ResourceDB][addClass] prepare failed err, ", err.Error())
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
		log.Fatalln("[ClassRepository][ResourceDB][addClass] problem query to db err", err.Error())
		return resp, err
	}

	return resp, nil
}
