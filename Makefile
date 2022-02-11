build:
	docker-compose up --build -d
rbuild:
	docker stop faaaar-front
	docker stop faaaar-server
	docker stop faaaar-db
	docker-compose up --build -d
build-db:
	docker-compose up -d --build db
stop:
	docker stop faaaar-front
	docker stop faaaar-server
	docker stop faaaar-db
run-srv:
	docker-compose exec server go mod download
	docker-compose exec server go run main.go
run-frt:
	docker-compose exec front go mod download
	docker-compose exec front go run main.go
srv-e:
	docker exec -it faaaar-server sh
db-e:
	docker exec -it faaaar-db bash
db-s:
	docker stop faaaar-db