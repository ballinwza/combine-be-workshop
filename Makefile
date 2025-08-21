dev:
	go run main.go

docker-build:
	docker compose build 

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-setup:
	@make docker-build
	@make docker-up