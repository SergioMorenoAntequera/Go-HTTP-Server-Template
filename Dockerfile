FROM go-base:latest

WORKDIR /app 

COPY . /app

RUN go build

CMD ["go", "run", "seran.dev"]

EXPOSE 3000

# ENV ENV_VAR=ohye

