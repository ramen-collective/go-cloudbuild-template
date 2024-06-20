BIN := bin

$(BIN):
	mkdir -p $(BIN)

include scripts/build/compose.mk
include scripts/build/docker.mk
include scripts/build/golang.mk

clean: golang-clean docker-clean

# Golang targets.
graph: golang-generate
dependencies: golang-get
build: golang-build
test: golang-test
checks: golang-checks

# Docker targets.
image: docker-build
image-nc: docker-no-cache docker-build
push: docker-push

# Compose targets.
run: compose-up
stop: compose-down

# Local
run-local: golang-run