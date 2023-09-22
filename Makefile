.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build ./cmd/main.go
.PHONY:build

unittest:
	go test ./pkg/usecases

integration-test:
	go test ./cmd ./pkg/drivers

test:
	go test -cover ./cmd ./pkg/usecases

plan:
	terraform plan -out=output.tfplan
	terraform show -no-color -json output.tfplan > output.json
