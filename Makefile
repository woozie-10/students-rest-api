build:
	docker-compose build app

run:
	docker-compose up app

test:
	go test -v ./tests