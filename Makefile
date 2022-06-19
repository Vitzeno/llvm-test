parser-gen:
	@echo "Generating parser..."
	@goyacc -p YY -o internal/parser.go -v internal/parser.output internal/grammar.y
	@echo "Done."

run:
	@go run cmd/main.go

test:
	@go test -v ./...
