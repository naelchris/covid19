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

	//ROUTE
	getR := router.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/greet", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello Backend"))
	})
	getR.Use(auth.MiddlewareValidateUserToken)

	//ROUTE
	//router.HandleFunc("/class", class.AddClassHandler).Methods("POST")
	router.HandleFunc("/covid", covid.AddCovidCases).Methods("POST")
	router.HandleFunc("/user", user.AddUser).Methods("POST")
	router.HandleFunc("/covid/days", covid.GetCasesByDay).Methods("GET")
	router.HandleFunc("/covid/increment", covid.GetCaseIncrement).Methods("GET")
	router.HandleFunc("/covid/months", covid.MonthlyCasesQueryHTTP).Methods("GET")
	router.HandleFunc("/user", user.GetUser).Methods("GET")

	authR := router.Methods(http.MethodPost).Subrouter()
	authR.HandleFunc("/login", auth.LoginHandler)

	postR := router.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/covid", covid.AddCovidCases)

	return router
}
