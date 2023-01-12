.PHONY: help test test-coverage-html install-golangci-lint lint

help: ## Show this help
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Execute unit tests
	go test -cover -covermode=atomic -coverprofile=coverage.out ./...

test-coverage-html: test ## Generate HTML coverage report
	go tool cover -html=coverage.out -o cover.html

install-golangci-lint: ## Install golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

lint: install-golangci-lint ## Lint the code
	$(shell go env GOPATH)/bin/golangci-lint run ./...
