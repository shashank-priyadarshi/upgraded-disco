# TODO: Test whether all APIs are functioning in current code: dev environment check

.PHONY: default
default:
	@echo 'No default target.'

.PHONY: build
build:
	docker buildx build --build-arg CONFIG_SOURCE=$(config_source) --build-arg CONFIG_PATH=$(config_path) -f $(base_dir)/build/$(mode)/upgraded-disco.Dockerfile -t $(image):$(version) .

.PHONY: run
run:
	docker-compose -f $(base_dir)/build/$(mode)/docker-compose.yml up -d

.PHONEY: down
down:
	docker-compose -f $(base_dir)/build/$(mode)/docker-compose.yml down

.PHONEY: deps
run-deps:

.PHONY: push-image
push:
	docker push

.PHONY: lint
lint:
	golangci-lint run

.PHONY: unit-test
unit-test:
	go test -race -coverprofile=coverage.out '$(base_dir)/...'
	go tool cover -html=coverage.out

.PHONY: runtime-checkst
