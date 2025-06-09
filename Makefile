.PHONY: start gen

start:
	go run cmd/main.go

gen:
	go generate ./internal/wire