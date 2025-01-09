# go_http_server_template:latest
FROM golang:latest

RUN apt-get update && apt-get install -y postgresql-client

WORKDIR /app 
COPY . .
EXPOSE 3000

RUN go mod tidy

CMD ["go", "run", "./cmd/app"]
