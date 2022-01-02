build:
	docker-compose build
up:
	docker-compose up -d
down:
	docker-compose down
exec:
	docker-compose exec go bash
fmt:
	docker-compose run --rm go go fmt ./...
vet:
	docker-compose run --rm go go vet ./...
gotest:
	docker-compose run --rm go go test -v ./...
