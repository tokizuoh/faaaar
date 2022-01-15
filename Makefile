dc-build:
	docker-compose up --build -d
dep:
	docker exec -it postgres-syani bash
dsp:
	docker stop postgres-syani