.PHONY: build
build:
	docker-compose build
.PHONY: up
up:
	docker-compose up -d
.PHONY: down
down:
	docker-compose down
.PHONY: exec
exec:
	docker-compose exec go bash

.PHONY: fmt
fmt:
	docker-compose run --rm go go fmt ./...
.PHONY: vet
vet:
	docker-compose run --rm go go vet ./...
.PHONY: gotest
gotest:
	docker-compose run --rm go go test -v ./...
