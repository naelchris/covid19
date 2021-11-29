package covid

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/naelchris/covid19/Internal/repository/covid"
	"github.com/naelchris/covid19/server"
)

func AddCovidCases(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	var (
		ctx  = r.Context()
		data covid.Cases
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("[cassesHandler][AddCovidCases][jsonNewDocoder] decode error err, ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	//TODO : Usecase Covid
	res, err := server.CovidUsecase.AddCases(ctx, data)
	if err != nil {
		log.Println("[cassesHandler][AddCovidCases] unable to read body err :", err)
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, res, timeStart)
}

func MonthlyCasesQueryHTTP(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	var (
		ctx = r.Context()
	)

	//get data from params
	country := r.FormValue("country")
	yearStr := r.FormValue("year")
	monthStr := r.FormValue("start_month")
	monthRangeStr := r.FormValue("month_range")

	//convert string to int
	year, errYear := strconv.Atoi(yearStr)
	month, errMonth := strconv.Atoi(monthStr)
	monthRange, errMonthRange := strconv.Atoi(monthRangeStr)

	if errYear != nil || errMonth != nil || errMonthRange != nil {
		server.RenderError(w, http.StatusBadRequest, errors.New("invalid params"), timeStart)
		return
	}

	res, err := server.CovidUsecase.MonthlyCasesQuery(ctx, country, year, month, monthRange)
	if err != nil {
		log.Println("[casesHandler][MonthlyCasesQuery] err,", err)
		return
	}

	server.RenderResponse(w, http.StatusOK, res, timeStart)
}

func GetCasesByDay(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	ctx := r.Context()

	//get data from params
	country := r.FormValue("country")
	yearStr := r.FormValue("year")
	monthStr := r.FormValue("month")
	startDateStr := r.FormValue("start_date")
	dateRangeStr := r.FormValue("date_range")

	//convert string to int
	year, errYear := strconv.Atoi(yearStr)
	month, errMonth := strconv.Atoi(monthStr)
	startdate, errStartDate := strconv.Atoi(startDateStr)
	dateRange, errDateRange := strconv.Atoi(dateRangeStr)

	if errYear != nil || errMonth != nil || errStartDate != nil || errDateRange != nil {
		server.RenderError(w, http.StatusBadRequest, errors.New("invalid params"), timeStart)
		return
	}

	res, err := server.CovidUsecase.GetCasesByDay(ctx, country, year, month, startdate, dateRange)
	if err != nil {
		log.Println("[classHandler][GetClassDayHandler] get case by day err :", err)
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, res, timeStart)
}

func GetCaseIncrement(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	ctx := r.Context()

	//get data from params
	country := r.FormValue("country")
	yearStr := r.FormValue("year")
	monthStr := r.FormValue("month")
	dateStr := r.FormValue("date")

	//convert string to int
	year, errYear := strconv.Atoi(yearStr)
	month, errMonth := strconv.Atoi(monthStr)
	date, errDate := strconv.Atoi(dateStr)

	if errYear != nil || errMonth != nil || errDate != nil {
		server.RenderError(w, http.StatusBadRequest, errors.New("invalid params"), timeStart)
		return
	}

	res, err := server.CovidUsecase.GetCaseIncrement(ctx, country, year, month, date)
	if err != nil {
		log.Println("[classHandler][GetClassDayHandler] get case by day err :", err)
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, res, timeStart)
}
