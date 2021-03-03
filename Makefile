BUILD_FLAGS := GOOS=linux CGO_ENABLED=0

build: bin/api

.PHONY: clean
clean:
	@rm -f ./bin/*

bin/api: clean
	@$(BUILD_FLAGS) go build -v -o ./bin/api ./cmd/api

build-Api: bin/api
	@cp ./bin/api $(ARTIFACTS_DIR)/api

.PHONY: up
up:
	@sam build
	@sam local start-api --warm-containers LAZY
