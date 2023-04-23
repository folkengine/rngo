
.PHONY: run
run:
	 go run cmd/rngo/main.go

.PHONY: tidy
tidy:
	go mod tidy
