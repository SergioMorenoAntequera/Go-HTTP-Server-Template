FROM golang:latest

WORKDIR /app 

COPY . /app

CMD ["go", "run", "./cmd/app"]

EXPOSE 3000

# ENV ENV_VAR=ohye

