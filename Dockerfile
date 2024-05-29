FROM golang:1.22.3-bullseye as builder

WORKDIR /app

COPY . .

RUN go build -o /app/main ./cmd/api

EXPOSE 8080

CMD ["/app/main"]