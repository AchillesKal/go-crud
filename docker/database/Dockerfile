FROM postgres

ENV POSTGRES_USER=pqgouser \
  POSTGRES_DB=gocrud_db

COPY ./docker/database/sql/init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432