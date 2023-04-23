
.PHONY: run
run:
	 go run cmd/rngo/main.go

.PHONY: tidy
tidy:
	go mod tidy

test: test-ginkgo

.PHONE: test-ginkgo
test-ginkgo:
	ginkgo pkg/rngo