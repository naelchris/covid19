package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/naelchris/covid19/Internal/config"
	coviddomain "github.com/naelchris/covid19/Internal/repository/covid"
	fetcherdomain "github.com/naelchris/covid19/Internal/repository/fetcher"
	"github.com/naelchris/covid19/Internal/repository/firestore"
	userdomain "github.com/naelchris/covid19/Internal/repository/user"
	authusecase "github.com/naelchris/covid19/Internal/usecase/auth"
	covidusecase "github.com/naelchris/covid19/Internal/usecase/covid"
	userusecase "github.com/naelchris/covid19/Internal/usecase/user"
	"github.com/robfig/cron/v3"
	"google.golang.org/api/option"
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
	CovidDomain     coviddomain.Domain
	UserDomain      userdomain.Domain
	FetcherDomain   fetcherdomain.Domain
	FirestoreDomain firestore.Domain

	//usecase
	CovidUsecase *covidusecase.CovidUsecase
	UserUsecase  *userusecase.UserUsecase
	AuthUsecase  *authusecase.AuthUsecase
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

	//init firestore gcs
	bucket := InitFirestoreGoogleCloudStorage()

	log.Println("[InitServer][database] succesful connection")

	//init domain
	CovidDomain = coviddomain.InitDomain(db)
	FetcherDomain = fetcherdomain.InitDomain()
	UserDomain = userdomain.InitDomain(db)
	FirestoreDomain = firestore.InitDomain("storage-1bb41.appspot.com", bucket)

	//init usecase
	CovidUsecase = covidusecase.NewCovidUsecase(CovidDomain, FetcherDomain)
	UserUsecase = userusecase.NewUserUsecase(UserDomain, FirestoreDomain, cfg)
	AuthUsecase = authusecase.NewAuthUsecase(UserDomain)

	//firestore.GenerateV4PutObjectSignedURL("storage-1bb41.appspot.com", "naelchris@gmail", "files/storage-1bb41-firebase-adminsdk-esal9-7d4133437b.json")

}

func InitCron() *cron.Cron {
	jakartaTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("err get time")
	}

	//init Cron
	c := cron.New(cron.WithLocation(jakartaTime))

	//Cron Scheduler
	c.AddFunc("0 0 * * *", UpsertCasesDataCron)

	return c
}

func Server(cfg ServerConfig, router *mux.Router, cron *cron.Cron) {
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

	go cron.Start()
	defer func() {
		cron.Stop()
		log.Println("[Server] Cron Finish ====")
	}()
	log.Println("[Server] Cron initialize ====")

	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println("[Server] error on shutting down HTTP Server, err: ", err.Error())
	}
}

func UpsertCasesDataCron() {
	fmt.Println("Upsert Cases Data === Start")
	for _, country := range coviddomain.ListCountry {
		CovidUsecase.UpsertDailyCasesData(context.Background(), country)
	}
	fmt.Println("Upsert Cases Data === Finish")
}

func InitFirestoreGoogleCloudStorage() *storage.BucketHandle {

	config := &firebase.Config{
		StorageBucket: "storage-1bb41.appspot.com",
	}

	opt := option.WithCredentialsFile("files/storage-1bb41-firebase-adminsdk-esal9-7d4133437b.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Created bucket handle: %v\n", bucket)
	return bucket
}
