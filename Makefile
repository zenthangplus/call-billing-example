up:
	docker-compose up

down:
	docker-compose down

tidy:
	go mod tidy

swagger:
	cd src/api && swag init --parseDependency --parseDepth=3

test:
	go test ./...
