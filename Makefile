run:
	go run cmd/main.go

docker-compose:
	@echo Starting local docker compose...
	docker-compose -f docker-compose.yaml up -d --build

