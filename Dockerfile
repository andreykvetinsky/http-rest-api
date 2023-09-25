FROM golang:1.20-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY cmd ./cmd
COPY configs ./configs
COPY internal ./internal
COPY migrations ./migrations

RUN go build -o ./bin/docker-apiserver cmd/apiserver/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/docker-apiserver /
COPY configs/apiserver.toml /configs/apiserver.toml

EXPOSE 8080

CMD ["/docker-apiserver"]