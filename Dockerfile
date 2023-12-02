FROM golang:1.21

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

RUN apt-get update && apt-get install -y wget

RUN wget -O /usr/local/bin/migrate https://github.com/golang-migrate/migrate/releases/download/v4.16.0/migrate.linux-amd64.tar.gz && \
    tar -xvzf /usr/local/bin/migrate -C /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate

COPY . .

RUN go build -o main ./cmd/api/

CMD ["sh", "-c", "migrate -path ./migrations -database postgres://pinata_user:12345678@db_container/pinata_db?sslmode=disable up && ./main"]