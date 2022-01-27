build:
	docker-compose up --build -d
build-db:
	docker-compose up -d --build db
stop:
	docker stop faaaar-server
	docker stop faaaar-db
run:
	docker-compose exec server go mod download
	docker-compose exec server go run main.go
srv-e:
	docker exec -it faaaar-server sh
db-e:
	docker exec -it faaaar-db bash
db-s:
	docker stop faaaar-db