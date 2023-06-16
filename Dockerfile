FROM golang:1.20.5-bookworm AS builder

ENV GOINSECURE="proxy.golang.org/*,github.com,github.com/*"
ENV GOPROXY=direct

WORKDIR /app
RUN git config --global http.sslverify false

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./dag

FROM debian:12.0

WORKDIR /app
COPY --from=builder /app/dag ./dag

ENTRYPOINT ["./dag"]
