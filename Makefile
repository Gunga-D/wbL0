d-ord-serv:
	go run ./cmd/main.go

t-ord-serv:
	go run ./test/test_nats_streaming.go
	
ord-serv:
	go build -o order-service ./cmd/main.go

serv: 
	docker compose up -d --build

p-serv:
	docker-compose down

clear-serv:
	docker-compose stop
	docker system prune