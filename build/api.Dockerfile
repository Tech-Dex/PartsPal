FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY cmd/api/main.go .
COPY pkg/ ./pkg/
COPY assets/ ./assets/
COPY go.mod .
COPY go.sum .

RUN CGO_ENABLED=0 go build -o partspal-server main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/partspal-server .

EXPOSE 3000

CMD ["./partspal-server"]
