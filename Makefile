BINARY=engine

engine :
	go build -o ${BINARY} app/*.go

gorun_http :
	go run main.go