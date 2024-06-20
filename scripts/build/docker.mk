SRC             := .
VERSION         := VERSION
MIGRATIONS      := scripts/sql/migrations

CURRENT_COMMIT		:=	$(shell git rev-parse HEAD)

GOLANG_BIN          := go
GOLANG_MIGRATE      := $(BIN)/migrate

DOCKER_BIN          := docker
DOCKER_BUILD_OPTS   := -f scripts/build/api-template.dockerfile
DOCKER_RUN_OPTS     := --rm -p 80:8000
DOCKER_REGISTRY     := 
DOCKER_TAG          := latest
DOCKER_IMAGE 		:= $(DOCKER_REGISTRY)/api-template:${DOCKER_TAG}

export DOCKER_IMAGE
export DOCKER_TAG
export DOCKER_REGISTRY
export CURRENT_COMMIT

LOCAL_DOCKER_MYSQL_IMAGE := mysql:latest
LOCAL_DOCKER_MYSQL_NAME  := root
LOCAL_DOCKER_MYSQL_DB    := api-template
LOCAL_DOCKER_MYSQL_PWD   := root
LOCAL_DOCKER_MYSQL_OPTS  := --name $(LOCAL_DOCKER_MYSQL_NAME)
LOCAL_DOCKER_MYSQL_OPTS  += -e MYSQL_DATABASE=$(LOCAL_DOCKER_MYSQL_DB)
LOCAL_DOCKER_MYSQL_OPTS  += -e MYSQL_ROOT_PASSWORD=$(LOCAL_DOCKER_MYSQL_PWD)
LOCAL_DOCKER_MYSQL_OPTS  += -d $(LOCAL_DOCKER_MYSQL_IMAGE)
LOCAL_DOCKER_MYSQL_OPTS  += --ngram-token-size=2 --innodb_ft_enable_stopword="OFF"

LOCAL_DOCKER_REDIS_IMAGE := redis:6-alpine
LOCAL_DOCKER_REDIS_NAME  := api-template-redis
LOCAL_DOCKER_REDIS_OPTS  := --name $(LOCAL_DOCKER_REDIS_NAME)
LOCAL_DOCKER_REDIS_OPTS  += -d $(LOCAL_DOCKER_REDIS_IMAGE)

SEMVER          := $(BIN)/git-semver
SEMVER_OPTS     := --author loutre -d -f raw -o $(VERSION)

UNAME_S = $(shell uname -s)

ifeq ($(UNAME_S),Darwin)
DOCKER_TERRAFORM_OPTS += --platform linux/amd64
endif

$(GOLANG_MIGRATE): $(BIN)
	GOBIN=$$(pwd)/$(BIN) go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

$(SEMVER): $(BIN)
	GOBIN=$$(pwd)/$(BIN) go install github.com/meinto/git-semver@latest

docker-clean:
	-$(DOCKER_BIN) rmi $(DOCKER_IMAGE)

docker-no-cache:
	$(eval DOCKER_BUILD_OPTS += --no-cache)

setup-local-config:
	cp scripts/api-template.yaml.local scripts/api-template.yaml

setup-local-db:
	$(DOCKER_BIN) run $(LOCAL_DOCKER_MYSQL_OPTS)

setup-local-cache:
	$(DOCKER_BIN) run $(LOCAL_DOCKER_REDIS_OPTS)

setup-local-env: setup-local-config setup-local-db setup-local-cache

start-local-env:
	$(DOCKER_BIN) start $(LOCAL_DOCKER_MYSQL_NAME)
	$(DOCKER_BIN) start $(LOCAL_DOCKER_REDIS_NAME)

stop-local-env:
	$(DOCKER_BIN) stop $(LOCAL_DOCKER_MYSQL_NAME)
	$(DOCKER_BIN) stop $(LOCAL_DOCKER_REDIS_NAME)

migrate-up: $(GOLANG_MIGRATE) start-local-env
	$(GOLANG_MIGRATE) -source file://$(MIGRATIONS) -database "mysql://root:$(LOCAL_DOCKER_MYSQL_PWD)@tcp(172.17.0.2:3306)/$(LOCAL_DOCKER_MYSQL_DB)" up

migrate-down: $(GOLANG_MIGRATE) start-local-env
	$(GOLANG_MIGRATE) -source file://$(MIGRATIONS) -database "mysql://root:$(LOCAL_DOCKER_MYSQL_PWD)@tcp(172.17.0.2:3306)/$(LOCAL_DOCKER_MYSQL_DB)" down 1

docker-build:
	$(DOCKER_BIN) build $(DOCKER_BUILD_OPTS) -t $(DOCKER_IMAGE):$(DOCKER_TAG) -t $(DOCKER_IMAGE):$(CURRENT_COMMIT) .

docker-push:
	./scripts/docker/push-docker-image.sh

docker-run: docker-build
	$(DOCKER_BIN) run $(DOCKER_RUN_OPTS) $(DOCKER_IMAGE)