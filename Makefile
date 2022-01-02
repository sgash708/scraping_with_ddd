build:
	docker-compose build
up:
	docker-compose up -d
down:
	docker-compose down
exec:
	docker-compose exec golang bash
fmt:
	docker-compose run --rm golang go fmt ./...
vet:
	docker-compose run --rm golang go vet ./...
gotest:
	docker-compose run --rm golang go test -v ./...
