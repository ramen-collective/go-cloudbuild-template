DOCKER_COMPOSE := docker-compose

compose-down:
	$(DOCKER_COMPOSE) -f deployments/compose/docker-compose.yaml down

compose-up:
	$(DOCKER_COMPOSE) -f deployments/compose/docker-compose.yaml up
