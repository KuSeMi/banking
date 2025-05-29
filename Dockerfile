FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /banking-service

FROM alpine:latest

COPY --from=builder /banking-service /banking-service

EXPOSE 8080

CMD ["/banking-service"]