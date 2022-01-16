build:
	docker-compose up --build -d
stop:
	docker stop faaaar-server
	docker stop postgres-syani
run:
	docker-compose exec server go mod download
	docker-compose exec server go run main.go
srv-e:
	docker exec -it faaaar-server sh
db-e:
	docker exec -it postgres-syani bash
db-s:
	docker stop postgres-syani