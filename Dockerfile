FROM golang:1.23.0-alpine AS base

RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd
COPY internal internal

FROM base AS builder

WORKDIR cmd/user-service
RUN go build -o /build/user-service .

FROM base AS debug
WORKDIR cmd/user-service
RUN go build -gcflags="all=-N -l" -o /build/user-service .

ENV DEBUG_PORT=2345
EXPOSE $DEBUG_PORT
ENTRYPOINT /go/bin/dlv --listen=:$DEBUG_PORT --headless=true --api-version=2 --accept-multiclient exec /build/user-service

FROM alpine AS run
COPY --from=builder /build/user-service /
CMD ["/user-service"]

