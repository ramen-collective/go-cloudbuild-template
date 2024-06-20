CACHE               := ../.cache
GOLANG_BIN          := go
GOLANG_111MODULE    := on
CGO_ENABLED         := 0
GOLANG_OS           :=
GOLANG_ARCH         :=
GOLANG_BUILD_OPTS   := GO111MODULE=$(GOLANG_111MODULE)
GOLANG_BUILD_OPTS   += GOOS=$(GOLANG_OS)
GOLANG_BUILD_OPTS   += GOARCH=$(GOLANG_ARCH)
GOLANG_BUILD_OPTS   += CGO_ENABLED=$(CGO_ENABLED)
GOLANG_LDFFLAGS     := -ldflags="-w -s"
GOLANG_BUILD_ARGS   := -mod vendor
GOLANG_LINT         := $(BIN)/golangci-lint

$(GOLANG_LINT): $(BIN)
	GOBIN=$$(pwd)/$(BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0

golang-clean:
	-rm $(GOLANG_BUILD_OUT)
	-rm .coverage

golang-get:
	$(GOLANG_BUILD_OPTS) $(GOLANG_BIN) get ./...

golang-generate:
	$(GOLANG_BUILD_OPTS) $(GOLANG_BIN) generate ./...

golang-build: $(BIN) golang-get golang-generate
	$(GOLANG_BUILD_OPTS) $(GOLANG_BIN) build $(GOLANG_LDFFLAGS) -o $(BIN)/api_template ./cmd/main.go
	chmod +x $(BIN)/api_template

golang-build-with-vendor: $(BIN)
	$(GOLANG_BUILD_OPTS) $(GOLANG_BIN) build $(GOLANG_LDFFLAGS) $(GOLANG_BUILD_ARGS) -o $(BIN)/api_template ./cmd/main
	chmod +x $(BIN)/api-template

$(BIN)/api_template: golang-build

golang-test: golang-get golang-generate
	$(GOLANG_BIN) test ./... -coverprofile=.coverage

golang-checks: $(GOLANG_LINT)
	go vet ./...
	gofmt -d -l .
	GOLANGCI_LINT_CACHE=$$(pwd)/$(CACHE) $(GOLANG_LINT) run ./...

golang-run: $(BIN)/api_template
	./$(BIN)/api_template
