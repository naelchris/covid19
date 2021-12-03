package user

import (
	"encoding/json"
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/naelchris/covid19/Internal/repository/user"
	"github.com/naelchris/covid19/server"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	timeStart := time.Now()

	ctx := r.Context()

	vars := mux.Vars(r)

	email := vars["email"]

	log.Println(email)

	userInfo := ctx.Value("data")

	log.Println(userInfo)

	// //get data from params
	// email := r.FormValue("email")
	// password := r.FormValue("password")

	// if email == "" || password == "" {
	// 	server.RenderError(w, http.StatusBadRequest, errors.New("invalid params"), timeStart)
	// 	return
	// }

	// res, err := server.UserUsecase.GetUser(ctx, email, password)
	// if err != nil {
	// 	log.Println("[classHandler][GetUserHandler] get case by day err :", err)
	// 	server.RenderError(w, http.StatusBadRequest, err, timeStart)
	// 	return
	// }

	server.RenderResponse(w, http.StatusCreated, "test", timeStart)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	timeStart := time.Now()

	ctx := r.Context()

	var files = make(map[string]multipart.File)

	vars := mux.Vars(r)

	email := vars["email"]

	for _, c := range []string{"certificate_1", "certificate_2"} {
		file, _, err := r.FormFile(c)
		r.ParseMultipartForm(10 << 20)
		if err != nil {
			continue
		}
		files[c] = file
	}

	dateOfBirth, err := time.Parse(time.RFC3339, r.FormValue("date_of_birth"))
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, errors.New("invalid params"), timeStart)
		return
	}

	log.Println(len(files))
	req := user.UserInfo{
		Email:        email,
		Name:         r.FormValue("name"),
		DateOfBirth:  dateOfBirth,
		HealthStatus: r.FormValue("health_status"),
	}

	if req.Name == "" || dateOfBirth.Year() == 1 {
		err = errors.New("invalid params")
		log.Println("[UpdateUser endpoint][UpdateUser] invalid,", err)
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	log.Println(req)

	resp, err := server.UserUsecase.UpdateUser(ctx, req, files)
	if err != nil {
		log.Println("[UpdateUser endpoint][UpdateUser] err,", err)
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	for _, f := range files {
		f.Close()
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
}
