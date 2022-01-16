build:
	docker-compose up --build -d
stop:
	docker stop faaaar-server
	docker stop postgres-syani
run:
	docker-compose exec server go run main.go
db-e:
	docker exec -it postgres-syani bash
db-s:
	docker stop postgres-syani