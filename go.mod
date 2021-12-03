module github.com/naelchris/covid19

// +heroku goVersion go1.16
go 1.16

require (
	firebase.google.com/go/v4 v4.6.1
	github.com/gorilla/mux v1.8.0
	github.com/huandu/go-sqlbuilder v1.13.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.3
	github.com/robfig/cron/v3 v3.0.1
	github.com/stretchr/testify v1.7.0 // indirect
	google.golang.org/api v0.60.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
