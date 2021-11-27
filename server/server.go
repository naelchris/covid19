package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/naelchris/covid19/Internal/config"
	coviddomain "github.com/naelchris/covid19/Internal/repository/covid"
	covidusecase "github.com/naelchris/covid19/Internal/usecase/covid"
)

type ServerConfig struct {
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Port         string
}

var (
	cfg config.Config
)

var (
	//domain
	CovidDomain coviddomain.Domain

	//usecase
	CovidUsecase *covidusecase.CovidUsecase
)

func InitServer() {
	//init the config
	cfg.InitConfig()

	//init the database connection
	connection, err := pq.ParseURL(cfg.Server.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal("Fail Connecting to database")
		return
	}

	if err := db.Ping(); err != nil {
		log.Fatal("fail to Ping database")
		return
	}

	log.Println("[InitServer][database] succesful connection")

	//init domain
	CovidDomain = coviddomain.InitDomain(db)

	//init usecase
	CovidUsecase = covidusecase.NewCovidUsecase(CovidDomain)
}

func Server(cfg ServerConfig, router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.Port,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	go func() {
		if err := http.ListenAndServe(cfg.Port, router); err != nil {
			log.Fatal("[Server] unable to listen and serve : ", err)
		}

	}()

	log.Println("[Server] HTTP server is running at port ", cfg.Port)

	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println("[Server] error on shutting down HTTP Server, err: ", err.Error())
	}
}
