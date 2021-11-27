package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/naelchris/covid19/server"
	http "github.com/naelchris/covid19/server/http"
	cron "github.com/robfig/cron/v3"
)

func main() {
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

	server.Server(serverConfig, router)
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

func initCron() {
	//jakarta time
	jakartaTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("err get time")
	}

	//init Cron
	c := cron.New(cron.WithLocation(jakartaTime))
	defer c.Stop()

	c.AddFunc("* * * * *", func() { fmt.Println("testing") })

	go c.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
}
