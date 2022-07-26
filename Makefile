
build:
	docker-compose build todoapp-golang

run:
	docker-compose up todoapp-golang

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go
