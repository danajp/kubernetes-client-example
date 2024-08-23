FROM golang:1.22 AS builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd/ cmd/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/ ./...

FROM ubuntu:latest

COPY --from=builder /workspace/bin/* /usr/local/bin/
