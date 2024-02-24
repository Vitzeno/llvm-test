parser-gen:
	@# -p YY changes yacc prefix to upper case, this exports it for golang packages to access
	@echo "Generating parser..."
	@goyacc -p YY -o parser/parser.go -v parser/parser.output parser/grammar.y
	@echo "Done."

run:
	@go run main.go -in=asm/mock

clang:
	@clang asm/mock.ll -o asm/mock.out

run-bin:
	@./asm/mock.out

test:
	@go test -v ./...
