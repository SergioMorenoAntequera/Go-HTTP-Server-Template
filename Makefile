
run:
	go run ./cmd/app

docker:
	docker-compose up --build

all: run
