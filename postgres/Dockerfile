FROM postgres:alpine
ADD schema.sql /docker-entrypoint-initdb.d/schema.sql
ENV PGDATA=/var/lib/postgresql/pgdata/
RUN adduser -D docker
USER docker
#VOLUME "pg_data"

USER postgres
ENV POSTGRES_USER postgres
ENV POSTGRES_DB switter
ENV POSTGRES_PASSWORD=password
EXPOSE 5432