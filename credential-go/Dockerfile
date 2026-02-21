FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY . .


RUN go mod tidy
RUN go build -o credential-service ./cmd

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/credential-service .

EXPOSE 8090

CMD ["./credential-service"]
