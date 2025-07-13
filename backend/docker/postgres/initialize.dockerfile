FROM postgres:17.5-alpine
COPY ./sql/* /docker-entrypoint-initdb.d/
