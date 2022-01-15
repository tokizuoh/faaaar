build:
	docker-compose up --build -d
# TODO: Makeでは `$` 以降が読まれないので対応する
# stop:
# docker stop $(docker ps -q)
run:
	docker-compose exec server go run main.go
db-e:
	docker exec -it postgres-syani bash
db-s:
	docker stop postgres-syani