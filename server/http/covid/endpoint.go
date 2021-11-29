package covid

import (
	"encoding/json"
	"log"
	"net/http"
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
		log.Println("[classHandler][AddClassHandler][jsonNewDocoder] decode error err, ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	//TODO : Usecase Covid
	res, err := server.CovidUsecase.AddCases(ctx, data)
	if err != nil {
		log.Println("[classHandler][AddClassHandler] unable to read body err :", err)
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, res, timeStart)

}
