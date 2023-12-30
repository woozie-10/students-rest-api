build:
	docker-compose build app
run:
	docker-compose up app
down:
	sudo docker-compose down
test:
	go test -v ./tests
swag:
	swag init -g cmd/main.go -o docs
