package main

import (
	"log"
	"os"
	"time"

	"github.com/naelchris/covid19/server"
	http "github.com/naelchris/covid19/server/http"
)

func main() {
	// util.InsertIntoCloudStorage()
	server.InitServer()

	router := http.ConfigureMuxRouter()

	//router := mux.NewRouter()
	//
	//router.HandleFunc("/class", homeRoute).Methods("GET")

	serverConfig := server.ServerConfig{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         GetPort(),
	}

	c := server.InitCron()

	server.Server(serverConfig, router, c)
}

//func homeRoute(res http.ResponseWriter, req *http.Request) {
//	fmt.Fprint(res, "<b>Hello world</b>")
//}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[GETPORT] connection port :", port)
	}

	return ":" + port
}
