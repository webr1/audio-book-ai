FROM golang:1.26 AS builder
WORKDIR /build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY ./src ./src
COPY ./cmd ./cmd

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/main.go

FROM alpine:3.23
ENV CONTAINER_MODE=1
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /out/app ./app

ENTRYPOINT ["./app"]
