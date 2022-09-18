test:
	go test -timeout 30s -v -run ./^.go github.com/kunalrsagar/go-dsa/problems

# https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
benchmark:
	go test -v -bench=. -count 5 -run=./^.go github.com/kunalrsagar/go-dsa/problems

