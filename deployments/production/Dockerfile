FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o tensor-graphql ./cmd/webservice/main.go

FROM scratch

WORKDIR /app
COPY --from=builder /app/tensor-graphql ./
COPY --from=builder /app/migration ./migration
COPY --from=builder /app/docs ./docs

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/tensor-graphql"]