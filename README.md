# Isekai Shop API by Go Echo

Package

* Go Viper - to connect from config.yaml to struct config.go
* Validator package - to alert and validate the field.

## Create docker container of postgresDB

docker run --name isekaishopdb -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres:latest

docker exec -it b5b40248aec2 bash

psql -U postgres 

CREATE DATABASE isekaishopdb;
