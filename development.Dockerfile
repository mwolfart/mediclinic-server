FROM golang:1.21.6

WORKDIR /app

COPY . . 

RUN go mod download \
    && go install github.com/pressly/goose/v3/cmd/goose@latest \
    && goose postgres "host=$DB_HOSTNAME port=$DB_PORT user=$DB_USERNAME password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" up

CMD ["bash", "-c", "go run ."]
