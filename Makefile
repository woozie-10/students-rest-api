build:
	docker-compose build app

run:
	docker-compose up app

test:
	go test -v ./tests

swag:
	swag init -g cmd/main.go -o docs
