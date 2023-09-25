# .PHONY:
# build:
#		go build -v ./cmd/apiserver

# # .PHONY: test
all: test build build_image containers_up
test:
		go test -v -race -timeout 30s ./internal/app/apiserver
		go test -v -race -timeout 30s ./internal/app/model
		go test -v -race -timeout 30s ./internal/app/store/sqlstore
		go test -v -race -timeout 30s ./internal/app/store/teststore

#.PHONY:
build:
		go build -v ./cmd/apiserver

build_image:
		docker build --tag andreikvetinskii/docker-apiserver:latest .
containers_up:
		docker compose up -d
		docker compose -f docker-compose.yaml --profile tools run --rm migrate up
#		docker compose -f docker-compose.yaml --profile tools run --rm migrate down
		

