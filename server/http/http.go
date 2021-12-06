package http

import (
	"net/http"

	"github.com/naelchris/covid19/server/http/auth"

	"github.com/gorilla/mux"
	"github.com/naelchris/covid19/server/http/covid"
	"github.com/naelchris/covid19/server/http/user"
)

func ConfigureMuxRouter() *mux.Router {
	router := mux.NewRouter()

	//Greeting but needed to login first
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/greet", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello Backend"))
	})
	getRouter.Use(auth.MiddlewareValidateUserToken)

	//Auth Route
	authRouter := router.Methods(http.MethodPost).Subrouter()
	authRouter.HandleFunc("/login", auth.LoginHandler)
	authRouter.HandleFunc("/register", auth.RegisterHandler)
	authRouter.Use(auth.MiddlewareValidateSession)

	//Auth Route Patch user
	authPatchRouter := router.Methods(http.MethodPatch).Subrouter()
	authPatchRouter.HandleFunc("/user", user.UpdateUser)
	authPatchRouter.Use(auth.MiddlewareValidateUserToken)

	//Auth Route GET user Info
	authGetUserInfo := router.Methods(http.MethodGet).Subrouter()
	authGetUserInfo.HandleFunc("/user", user.GetUser)
	authGetUserInfo.Use(auth.MiddlewareValidateUserToken)

	//Covid Route
	covidPostRouter := router.Methods(http.MethodPost).Subrouter()
	covidPostRouter.HandleFunc("/covid", covid.AddCovidCases)
	covidPostRouter.HandleFunc("/covid/init", covid.InitCovidCases)
	covidPostRouter.HandleFunc("/covid/daily", covid.UpsertDailyCasesData)
	covidPostRouter.HandleFunc("/user", user.AddUser)

	covidGetRouter := router.Methods(http.MethodGet).Subrouter()
	covidGetRouter.HandleFunc("/covid/days", covid.GetCasesByDay)
	covidGetRouter.HandleFunc("/covid/increment", covid.GetCaseIncrement)
	covidGetRouter.HandleFunc("/covid/months", covid.MonthlyCasesQueryHTTP)

	//router.HandleFunc("/user/{email}", user.UpdateUser).Methods("PATCH")

	return router
}
