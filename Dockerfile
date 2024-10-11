FROM golang:1.21.5-alpine AS builder 
WORKDIR /app
COPY . .
RUN go build -o main main.go


FROM alpine:3.20.3
WORKDIR /app

COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]