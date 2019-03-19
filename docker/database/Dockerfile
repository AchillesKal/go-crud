FROM postgres

ENV POSTGRES_DB gocrud_db

COPY ./sql/init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432