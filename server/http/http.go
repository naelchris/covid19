package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naelchris/covid19/server/http/covid"
)

func ConfigureMuxRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello Backend"))
	}).Methods("GET")

	//ROUTE
	//router.HandleFunc("/class", class.AddClassHandler).Methods("POST")
	router.HandleFunc("/covid", covid.AddCovidCases).Methods("POST")
	router.HandleFunc("/covid/days", covid.GetCasesByDay).Methods("GET")
	router.HandleFunc("/covid/increment", covid.GetCaseIncrement).Methods("GET")
	router.HandleFunc("/covid/months", covid.MonthlyCasesQueryHTTP).Methods("GET")

	return router
}
