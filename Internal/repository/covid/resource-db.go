package covid

import (
	"context"
	"log"
)

func (s storage) AddCases(ctx context.Context, data InsertCasesRequest) (CasesResponse, error) {
	var resp CasesResponse
	var id int64

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
		data.Lat,
		data.Lon,
		data.Confirmed,
		data.Deaths,
		data.Recovered,
		data.Active,
	).Scan(&id)
	if err != nil {
		log.Fatalln("[ClassRepository][ResourceDB][addClass] problem query to db err", err.Error())
		return resp, err
	}

	resp = CasesResponse{
		ID: id,
	}

	return resp, nil
}
