run-client:
	go run cmd/main.go

run-worker-service:
	go run services/worker/cmd/main.go

run-summary-service:
	go run services/summary/cmd/main.go

docker-compose:
	@echo Starting local docker compose...
	docker-compose -f docker-compose.yaml up -d --build

protoc:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
      --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
      proto/*.proto