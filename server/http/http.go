package http

import (
	"github.com/naelchris/covid19/server/http/auth"
	"net/http"

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
	covidGetRouter.HandleFunc("/user", user.GetUser)

	return router
}
