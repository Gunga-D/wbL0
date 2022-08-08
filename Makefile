test:
	go test -v -cover -covermode=atomic ./...

order-service:
	go run ./cmd/main.go

builded-order-service:
	go build -o ./bin/order-service ./cmd/main.go

builded-server:
	docker-compose build 

launched-server:
	docker-compose up -d

stopped-server:
	docker-compose stop
	docker system prune

paused-server:
	docker-compose down

launched-builded-server: 
	docker compose up -d --build