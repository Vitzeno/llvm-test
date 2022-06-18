parser-gen:
	@echo "Generating parser..."
	@goyacc -o cmd/parser.go -v cmd/parser.output cmd/grammar.y
	@echo "Done."

run:
	@go run cmd/*.go
