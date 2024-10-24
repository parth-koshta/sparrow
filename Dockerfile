FROM golang:1.21.5-alpine AS builder 
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-w -s" -o main main.go

RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz \
    && rm -rf /var/cache/apk/*

FROM alpine:3.20.3
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY db/migration ./migration
COPY start.sh .

EXPOSE 8080

RUN chmod +x /app/start.sh
ENTRYPOINT ["/app/start.sh"]
