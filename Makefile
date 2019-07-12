default: build

build: 
	go build go-rest-template.go

run: 
	go run go-rest-template.go

install:	
	go install go-rest-template.go
