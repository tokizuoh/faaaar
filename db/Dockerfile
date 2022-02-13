FROM mysql:8.0.28
WORKDIR ./
COPY init /docker-entrypoint-initdb.d/
COPY other /other/
ADD ./my.cnf /etc/mysql/my.cnf