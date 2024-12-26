
run:
	go run ./cmd/app

docker:
	docker build -t go-http-server-template .
	docker run -it --rm -p 3000:3000 --name go-http-server-template go-http-server-template 

all: run
