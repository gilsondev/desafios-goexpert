.PHONY: setup
setup:
	@go mod tidy

.PHONY: export-diagrams
export-diagrams:
	@docker-compose run --rm plantuml /usr/docs/assets/diagrams/*.puml

.PHONY: run-server
run-server:
	@go run server/main.go

.PHONY: run-client
run-client:
	@go run client/main.go
