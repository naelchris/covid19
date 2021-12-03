package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/naelchris/covid19/Internal/repository/user"
	"github.com/naelchris/covid19/server"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	var (
		ctx  = r.Context()
		data user.User
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("[classHandler][AddUserHandler][jsonNewDocoder] decode error err, ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	if data.Name == "" || data.Email == "" || data.Password == "" || data.DateOfBirth.Year() == 1 {
		log.Println("[classHandler][AddUserHandler] invalid body")
		server.RenderError(w, http.StatusInternalServerError, errors.New("invalid body"), timeStart)
		return
	}

	res, err := server.UserUsecase.AddUser(ctx, data)
	if err != nil {
		log.Println("[classHandler][AddUserHandler] unable to read body err :", err)
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, res, timeStart)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	ctx := r.Context()

	//get data from params
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		server.RenderError(w, http.StatusBadRequest, errors.New("invalid params"), timeStart)
		return
	}

	res, err := server.UserUsecase.GetUser(ctx, email, password)
	if err != nil {
		log.Println("[classHandler][GetUserHandler] get case by day err :", err)
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, res, timeStart)
}
