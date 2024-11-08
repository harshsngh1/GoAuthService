FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goauthservice ./cmd/main.go

FROM alpine:latest  

COPY --from=builder /app/goauthservice .

EXPOSE 8080

CMD ["./goauthservice"]
