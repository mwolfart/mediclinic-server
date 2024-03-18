FROM golang:1.21.6

WORKDIR /app

COPY . . 

RUN go mod download \
    && go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["bash", "-c", "(cd migrations && chmod a+x run_migrations.sh && ./run_migrations.sh) && sqlboiler psql && go run ."]
