FROM postgres:alpine
ADD 01-schema.sql  /docker-entrypoint-initdb.d/01-schema.sql
ADD 02-init.sql /docker-entrypoint-initdb.d/02-init.sql
ENV POSTGRES_USER postgres
ENV POSTGRES_DB switterdb
ENV POSTGRES_PASSWORD=password
EXPOSE 5432