default: build

build: 
	go build main.go

run: 
	go run main.go

install:	
	go install main.go

db:
	docker-compose -f project-infra/docker-postgres-db/compose.yml up

db-clean:
	docker-compose -f project-infra/docker-postgres-db/compose.yml down -v --remove-orphans
