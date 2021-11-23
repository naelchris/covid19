package covid

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type CasesResponse struct {
	ID          int64  `json:"case_ID,omitempty" db:"id"`
	Country     string `json:"country, omitempty" db:"country"`
	CountryCode string `json:"country_code, omitempty" db:"countr_code"`
	Province    string `json:"province, omitempty" db:"province"`
	City        string `json:"city, omitempty" db:"city"`
	CityCode    string `json:"city_code, omitempty" db:"city_code"`
	Lat         string `json:"lat, omitempty" db:"lat"`
	Lon         string `json:"lon, omitempty" db:"lon"`
	Confirmed   int64  `json:"confirmed, omitempty" db:"confirmed"`
	Deaths      int64  `json:"death, omitempty" db:"death"`
	Recovered   int64  `json:"recovered, omitempty" db:"recovered"`
	Active      int64  `json:"active, omitempty" db:"active"`
}

type InsertCasesRequest struct {
	Country     string `json:"country, omitempty" db:"country"`
	CountryCode string `json:"country_code, omitempty" db:"countr_code"`
	Province    string `json:"province, omitempty" db:"province"`
	Lat         string `json:"lat, omitempty" db:"lat"`
	Lon         string `json:"lon, omitempty" db:"lon"`
	Confirmed   int64  `json:"confirmed, omitempty" db:"confirmed"`
	Deaths      int64  `json:"death, omitempty" db:"death"`
	Recovered   int64  `json:"recovered, omitempty" db:"recovered"`
	Active      int64  `json:"active, omitempty" db:"active"`
}

type UpdateCasesRequest struct {
	Country   string `json:"country, omitempty" db:"country"`
	Province  string `json:"province, omitempty" db:"province"`
	Confirmed int64  `json:"confirmed, omitempty" db:"confirmed"`
	Deaths    int64  `json:"death, omitempty" db:"death"`
	Recovered int64  `json:"recovered, omitempty" db:"recovered"`
	Active    int64  `json:"active, omitempty" db:"active"`
}

func (u UpdateCasesRequest) BuildQuery(id int64) (string, []interface{}) {
	var fieldQuery string
	fieldValues := make([]interface{}, 0)

	var i = 0
	if u.Country != "" {
		fieldQuery += fmt.Sprintf("Country=%d", i)
		fieldValues = append(fieldValues, u.Country)
		i++
	}
	if u.Province != "" {
		fieldQuery += fmt.Sprintf("province=%d", i)
		fieldValues = append(fieldValues, u.Province)
		i++
	}
	if u.Confirmed != 0 {
		fieldQuery += fmt.Sprintf("confirmed=%d", i)
		fieldValues = append(fieldValues, u.Confirmed)
		i++
	}
	if u.Deaths != 0 {
		fieldQuery += fmt.Sprintf("death=%d", i)
		fieldValues = append(fieldValues, u.Deaths)
		i++
	}
	if u.Recovered != 0 {
		fieldQuery += fmt.Sprintf("recovered=%d", i)
		fieldValues = append(fieldValues, u.Recovered)
		i++
	}

	finalQuery := fmt.Sprintf(updateCaseQuery, fieldQuery[:len(fieldQuery)-1], id)
	return finalQuery, fieldValues

}

type Metadata struct {
	BatchType int `json:"batch_type"`
}

type storage struct {
	CasesDB *sql.DB
}

//json encoded representation of struct
func (md Metadata) Value() (driver.Value, error) {
	return json.Marshal(md)
}

//json decoded value into the struct
func (md *Metadata) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &md)
}
