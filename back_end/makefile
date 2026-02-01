#SHELL=/bin/bash
ENV_FILE?=.env.dev

# start up local services
local:
	sudo docker compose --env-file .env.dev -f docker-compose.yaml up --build

# exec unit tests for providers and services
unit-test:
	go test ./providers/... ./services/...

# clean any existing containers related to application
clean:
	source ./$(ENV_FILE); sudo docker compose --env-file $(ENV_FILE) -f docker-compose.yaml down -v
