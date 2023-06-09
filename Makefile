
.PHONY: init
init:
	git submodule init
	git submodule update

.PHONY: run
run:
	 go run cmd/rngo/main.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go run github.com/onsi/ginkgo/v2/ginkgo -r -v --randomize-all --randomize-suites --fail-on-pending --keep-going --cover --race --trace

.PHONE: test-ginkgo
test-ginkgo:
	ginkgo pkg/rngo