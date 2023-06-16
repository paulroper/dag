FROM golang:1.20.5-bookworm AS builder

WORKDIR /app
COPY go.mod ./

RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./dag

FROM debian:12.0

WORKDIR /app
COPY --from=builder /app/dag ./dag

CMD ["./dag"]



