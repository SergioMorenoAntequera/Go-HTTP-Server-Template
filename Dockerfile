FROM golang:latest

WORKDIR /app 

COPY . /app

RUN go build

CMD ["go", "run", "."]

EXPOSE 3000

# ENV ENV_VAR=ohye

