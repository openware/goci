FROM golang:1.13-alpine AS builder

WORKDIR /build
ENV CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build --ldflags "-X main.Version=$(cat .tags)" -o /build/goci ./cmd/goci

FROM alpine:3.9

RUN apk add git
WORKDIR app
COPY --from=builder /build/goci ./

CMD ["./goci"]
