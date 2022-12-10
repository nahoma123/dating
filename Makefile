swagger-init:
	- swag init -g initiator/initiator.go

run:
	- go run cmd/main.go

# note: docker will create the database file
prepare-db:
	- sudo docker compose up -d