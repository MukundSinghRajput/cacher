test: 
	@go test ./tests -v

benchmark:
	@go test -bench=. ./tests -v