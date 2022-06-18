parser-gen:
	@echo "Generating parser..."
	@goyacc -o parser.go -v parser.output grammar.y
	@echo "Done."

run:
	@go run *.go