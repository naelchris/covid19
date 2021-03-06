module github.com/naelchris/covid19

// +heroku goVersion go1.16
go 1.16

require (
	cloud.google.com/go/storage v1.10.0
	firebase.google.com/go/v4 v4.6.1
	github.com/GerardSoleCa/wordpress-hash-go v0.0.0-20161116172340-2bdd71ec2eb6 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/huandu/go-sqlbuilder v1.13.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.3
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/cors v1.8.0
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1
	google.golang.org/api v0.60.0
	gopkg.in/yaml.v2 v2.4.0
)
