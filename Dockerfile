FROM golang:1.20-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download

# build
COPY app ./
RUN go build -o ./bin/docker-apiserver cmd/apiserver/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/docker-apiserver /
COPY app/configs/apiserver.toml /configs/apiserver.toml

COPY configs/configs.yaml /configs.yaml

CMD ["/docker-apiserver"]