#SHELL=/bin/bash
ENV_FILE?=.env.dev

# start up local services
local:
	sudo docker compose --env-file $(ENV_FILE) -f docker-compose.yaml up --build

# exec unit tests for providers and services
unit-test:
	go test ./providers/... ./services/...

# exec integration tests defined in tests.yaml
integration-test: ENV_FILE=.env.integration
integration-test: clean
	sudo docker compose --env-file $(ENV_FILE) -f docker-compose.yaml up --build -d
	venom run ./testing/integration.yaml
	sleep 2s
	make clean

test: unit-test integration-test
	@echo testing complete

# clean any existing containers related to application
clean:
	rm -f ./testing/test_results_._testing_integration.xml ./testing/venom.log
	sudo docker compose --env-file $(ENV_FILE) -f docker-compose.yaml down -v

# install dependencies to run tests
deps-install:
	@# install venom integration test runner
	sudo curl https://github.com/ovh/venom/releases/download/v1.2.0/venom.linux-amd64 -L \
		-o /usr/local/bin/venom && sudo chmod +x /usr/local/bin/venom
	venom version
	@# TODO install mockgen

