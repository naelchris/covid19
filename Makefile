BINARY=engine

engine :
	go build -o bin/BackendEkskul main.go

gorun_http :
	go run main.go