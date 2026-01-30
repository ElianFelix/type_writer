include .env.dev

# start up local services
local:
	sudo docker compose --env-file .env.dev up --build
# exec unit tests for providers and services
unit-test:
	go test ./providers/... ./services/...
