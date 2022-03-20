# start with base image
FROM postgres:latest

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./database/migration_postgre.sql /docker-entrypoint-initdb.d/