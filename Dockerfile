FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-s -w" \
    -o repository-service \
    ./cmd/api


FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=builder /app/repository-service /repository-service

USER nonroot:nonroot

EXPOSE 50052

ENTRYPOINT ["/bank-repository-service"]
